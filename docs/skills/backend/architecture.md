# Backend Architecture

The backend lives under `services/app-api` and follows a modular multi-binary Clean Architecture in Go, designed to protect business rules from external technology details without overengineering.

## Overview

`services/app-api` produces two distinct binaries from the same codebase:
- `cmd/api`: The main HTTP REST server powered by `go-chi/chi/v5` serving public endpoints under `/api/*` (consumed by the Astro frontend) and hosting the `go-admin` panel under `/admin/*`.
- `cmd/sync`: The GitHub repository synchronization worker, designed to run as an independent scheduled process/container.

Both binaries share business logic, domain models, and infrastructure implementations located inside `internal/`.

## Directory Structure

```text
services/app-api/
├── cmd/
│   ├── api/             # Entrypoint: compõe HTTP server & GoAdmin (main.go)
│   └── sync/            # Entrypoint: compõe Worker de sincronização (main.go)
├── internal/
│   ├── domain/          # Entidades puras, regras de negócio e erros de domínio
│   │   ├── entities/    # Structs de domínio (Project, Education, etc.) com métodos de validação
│   │   ├── repositories/# Contratos de persistência e serviços (sem camada ports separada)
│   │   └── errors/      # Erros sentinela de domínio
│   │       └── validation.go # ErrDomainEntityValidation para validações gerais de entidades
│   ├── application/     # Casos de uso e DTOs de orquestração
│   │   ├── usecases/    # Implementações de casos de uso (New...UseCase)
│   │   └── dto/         # DTOs de entrada/saída e funções de mapeamento (toDomain, toOutput)
│   ├── infrastructure/  # Implementações concretas de tecnologia
│   │   ├── database/    # Pool pgx/v5, ping e helper de teste de integração (setup_test_pool.go)
│   │   ├── repositories/# Repositórios SQL concretos (AboutDatabaseRepository, etc.)
│   │   ├── github/      # Cliente HTTP da API do GitHub
│   │   └── admin/       # Setup e tabelas do GoAdmin
│   ├── presentation/    # Camada de transporte HTTP (handlers independentes de rotas)
│   │   ├── handlers/    # Handlers HTTP que chamam apenas Casos de Uso
│   │   └── response/    # Helpers de serialização JSON e mapeamento de erros HTTP
│   ├── modules/         # Camada de controle/composição: constrói os módulos da aplicação
│   │   ├── module.go    # Struct agregadora Module (NewModule) que instancia todos os submódulos
│   │   └── about/       # Subpacote do módulo About (AboutModule: Repo -> UseCase -> Handler)
│   └── router/          # Camada de controle de roteamento: recebe *modules.Module e registra rotas com chi/v5
├── pkg/                 # Código utilitário compartilhado sem regra de negócio
└── migrations/          # SQL migrations gerenciadas pelo golang-migrate
```

## Dependency Direction

Dependencies point strictly inward and linearly from top-level orchestration down to the core domain:

```text
cmd/api → Router → Modules → [Presentation, Application, Infrastructure]
Presentation → Application → Domain
Infrastructure → Domain contracts
```

### 1. Camada de Domínio (`internal/domain`)
- O coração do sistema: regras de negócio puras, **sem nenhuma dependência externa** (sem SQL, sem HTTP, sem tags de framework).
- **Entidades com Invariantes e Construtores Completos:**
  - O construtor `New<Entity>(...)` deve receber **todos os campos** da entidade.
  - **Inicialização Defensiva:**
    - Slices nunca permanecem `nil`:
      ```go
      if slice == nil {
          slice = []T{}
      }
      ```
    - Timestamps verificam `.IsZero()` e atribuem `time.Now()` se vazios:
      ```go
      if createdAt.IsZero() {
          createdAt = time.Now()
      }
      if updatedAt.IsZero() {
          updatedAt = time.Now()
      }
      ```
    - Strings higienizadas com `strings.TrimSpace(...)`.
  - Construtores invocam obrigatoriamente o método `Validate() error` antes de retornar a instância.
- **Contratos no Domínio (`internal/domain/repositories`):** Não usamos uma camada `ports/` separada; todos os contratos e interfaces de repositórios e serviços vivem no domínio, definindo as operações mínimas que a aplicação consome.
- **Erros Sentinela de Validação (`internal/domain/errors/validation.go`):** Define `ErrDomainEntityValidation` para representar falhas gerais de validação de entidades. As entidades embrulham suas mensagens específicas com `%w`:
  ```go
  return fmt.Errorf("%w: project name is required", errors.ErrDomainEntityValidation)
  ```
  Isso permite que a camada de apresentação identifique erros de validação via `errors.Is` e responda automaticamente com HTTP 400 Bad Request.
- **Tabelas Relacionais Auxiliares / Pivot N:N Não São Entidades:**
  - Tabelas de junção relacional (como `projects_skills`) existem exclusivamente para atender restrições de bancos relacionais (SQL). Elas **não** devem ser representadas como entidades em `internal/domain/entities`.
  - O modelo de domínio agrega a associação diretamente na entidade principal (ex: `Project` contém `Skills []Skill`), expondo métodos de comportamento como `AddSkill(skill Skill)` e `RemoveSkill(skillUuid string)`.
  - A leitura e gravação na tabela associativa física é responsabilidade da camada de infraestrutura dentro dos repositórios SQL.
- **Testes de Domínio:** Devem rodar de forma pura e imediata, **sem necessidade de mocks**.

### 2. Camada de Aplicação (`internal/application`)
- Orquestra o fluxo de dados entre as regras de domínio e os repositórios/serviços declarados no domínio.
- **Dependency Injection:** Casos de uso recebem interfaces no construtor (`New...UseCase(repository Repository)`).
- **Sem Abreviações:** Variáveis, campos e dependências injetadas usam nomes completos (`aboutRepository` em vez de `aboutRepo`, `workExperience` em vez de `workExp`).
- **DTOs & Mapeamento:** Desacopla o payload externo da entidade interna com structs explícitas (`CreateProjectInput`, `ProjectOutput`) e funções puras de conversão, ou retorna a entidade diretamente para operações simples de leitura.
- **Testes de Casos de Uso:** Testados isoladamente utilizando mocks das interfaces (via `testify/mock`).

### 3. Camada de Infraestrutura (`internal/infrastructure`)
- Implementações concretas dos contratos: adaptadores de banco de dados (`pgx/v5`), clientes HTTP externos (GitHub API) e painel GoAdmin.
- **Divisão Interna de Responsabilidades:**
  - `infrastructure/database/`: Gerencia o ciclo de vida da conexão com PostgreSQL (`pool.go`). Contém também o helper centralizado de integração `setup_test_pool.go` (`database.SetupTestDatabase(t *testing.T)`).
    > **Mecânica Go de Testes Cross-Package:** O helper `setup_test_pool.go` usa a tag `//go:build integration` e não o sufixo `_test.go`. Arquivos com terminação `_test.go` nunca são exportados para outros pacotes no Go. A combinação de nome normal + build constraint `integration` permite que os repositórios (`infrastructure/repositories_test`) compartilhem o setup sem que o helper seja compilado nos binários de produção.
  - `infrastructure/repositories/`: Repositórios SQL concretos (`AboutDatabaseRepository`, etc.) que implementam as interfaces de `internal/domain/repositories/`. Consomem `COALESCE` para mapear colunas nulas diretamente em campos de structs de domínio sem ponteiros intermediários.
- **Garantia em Tempo de Compilação:** Todo repositório deve declarar a verificação de cumprimento de contrato:
  ```go
  var _ repositories.ProjectRepository = (*ProjectPostgresRepository)(nil)
  ```
- **Tradução de Erros:** Converte erros específicos de infraestrutura (como `pgx.ErrNoRows`) para erros sentinela de domínio (`domain.ErrNotFound`).

### 4. Camada de Apresentação / Transporte (`internal/presentation`)
- Handlers HTTP implementados com a biblioteca padrão `net/http`, completamente independentes e desacoplados do roteador da API.
- **Regra de Ouro:** Handlers HTTP conhecem **apenas Casos de Uso** (`internal/domain/usecases`). Nunca acessam repositórios, pools de conexão, roteadores ou clients externos diretamente.
- Responsabilidades do Handler:
  1. Decodificar e validar o JSON da requisição no DTO de entrada (quando aplicável).
  2. Invocar `usecase.Execute(ctx, ...)`.
  3. Mapear erros de domínio para códigos HTTP semânticos (ex: `ErrNotFound` &rarr; 404, validação &rarr; 400, outros &rarr; 500).
  4. Serializar a entidade ou DTO de saída em JSON para a resposta.

### 5. Camada de Módulos (`internal/modules`)
- Camada de controle e injeção de dependência situada acima de `presentation`, `application` e `infrastructure`.
- **Organização Modular:**
  - **Subpastas por Feature (`internal/modules/<feature>/`)**: Cada funcionalidade possui sua própria pasta e subpacote (ex: `internal/modules/about/about_module.go` com `AboutModule` e `NewAboutModule(pool *pgxpool.Pool) *AboutModule`), responsável por instanciar `Repository -> UseCase -> Handler`.
  - **Agregador Raiz (`internal/modules/module.go`)**: Contém a struct `Module` e o construtor `NewModule(pool *pgxpool.Pool) *Module`, que centraliza a construção de todos os submódulos da aplicação em um único ponto.
- Isso garante que os `handlers` permaneçam 100% independentes e que o `router` receba apenas um único argumento (`*modules.Module`), evitando poluição de parâmetros conforme novos módulos forem adicionados.

### 6. Camada de Roteamento (`internal/router`)
- Camada de controle superior que recebe a struct agregadora `*modules.Module` (`NewRouter(applicationModule *modules.Module) *chi.Mux`) e registra todas as rotas HTTP da aplicação utilizando `go-chi/chi/v5` (`chi.NewRouter()`).
- Como está posicionada acima de `modules`, o fluxo de dependências permanece estritamente linear (`router -> modules -> presentation/handlers`), evitando qualquer caminho inverso entre camadas.
- Define os endpoints REST (ex: `GET /about` apontando para `applicationModule.About.Handler.Get`).

### 7. Entrypoints e Composição (`cmd/*/main.go`)
- Ponto de entrada enxuto que orquestra a inicialização ("de fora para dentro"):
  1. Carrega variáveis de ambiente do `.env` (`DATABASE_URL`, `PORT`).
  2. Inicializa a infraestrutura base (pool de conexões PostgreSQL via `database.NewPool`).
  3. Instancia o agregador de módulos (`applicationModule := modules.NewModule(pool)`) e o roteador (`router.NewRouter(applicationModule)`).
  4. Sobe o servidor HTTP na porta configurada (`PORT`) ou dispara a rotina do worker (`cmd/sync`).
- Não utiliza frameworks mágicos de Dependency Injection (como Fx ou Wire), mantendo a árvore de dependências explícita, legível e rastreável em tempo de compilação.

## PostgreSQL 18+ Volume & Storage Standards

O projeto utiliza PostgreSQL 18 (`postgres:18.6`).

- **Volume Mount Path**: No `docker-compose.yaml`, o volume DEVE ser montado em `/var/lib/postgresql`, nunca em `/var/lib/postgresql/data`:
  ```yaml
  volumes:
    - postgres_data:/var/lib/postgresql
  ```
- **Mecânica Técnica**:
  - A partir do PostgreSQL 18, as imagens oficiais adotam o layout do `pg_ctlcluster` (subpastas por versão maior).
  - Isso viabiliza upgrades zero-copy com `pg_upgrade --link`. Como *hard links* do Linux não cruzam pontos de montagem de volumes, a pasta pai `/var/lib/postgresql` garante que versões antigas e novas coexistam no mesmo volume.
  - Montar em `/var/lib/postgresql/data` no PostgreSQL 18+ causa falha intencional de inicialização no container.
