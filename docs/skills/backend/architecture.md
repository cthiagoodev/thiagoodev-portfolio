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
│   │   ├── database/    # Pool pgx/v5 e repositórios SQL
│   │   ├── github/      # Cliente HTTP da API do GitHub
│   │   └── admin/       # Setup e tabelas do GoAdmin
│   └── presentation/    # Camada de transporte HTTP
│       ├── handlers/    # Handlers HTTP que chamam apenas Casos de Uso
│       ├── router/      # Definição e agrupamento de rotas Chi v5
│       └── response/    # Helpers de serialização JSON e mapeamento de erros HTTP
├── pkg/                 # Código utilitário compartilhado sem regra de negócio
└── migrations/          # SQL migrations gerenciadas pelo golang-migrate
```

## Dependency Direction

Dependencies point strictly inward:

```text
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
- **Testes de Domínio:** Devem rodar de forma pura e imediata, **sem necessidade de mocks**.

### 2. Camada de Aplicação (`internal/application`)
- Orquestra o fluxo de dados entre as regras de domínio e os repositórios/serviços declarados no domínio.
- **Dependency Injection:** Casos de uso recebem interfaces no construtor (`New...UseCaseImpl(repo Repository)`).
- **DTOs & Mapeamento:** Desacopla o payload externo da entidade interna com structs explícitas (`CreateProjectInput`, `ProjectOutput`) e funções puras de conversão.
- **Testes de Casos de Uso:** Testados isoladamente utilizando mocks das interfaces (via `testify/mock`).

### 3. Camada de Infraestrutura (`internal/infrastructure`)
- Implementações concretas dos contratos: adaptadores de banco de dados (`pgx/v5`), clientes HTTP externos (GitHub API) e painel GoAdmin.
- **Garantia em Tempo de Compilação:** Todo repositório deve declarar a verificação de cumprimento de contrato:
  ```go
  var _ repositories.ProjectRepository = (*ProjectPostgresRepository)(nil)
  ```
- **Tradução de Erros:** Converte erros específicos de infraestrutura (como `pgx.ErrNoRows`) para erros sentinela de domínio (`domain.ErrNotFound`).

### 4. Camada de Apresentação / Transporte (`internal/presentation`)
- Handlers HTTP implementados com `go-chi/chi/v5`.
- **Regra de Ouro:** Handlers HTTP conhecem **apenas Casos de Uso**. Nunca acessam repositórios, pools de conexão ou clients externos diretamente.
- Responsabilidades do Handler:
  1. Decodificar e validar o JSON da requisição no DTO de entrada.
  2. Invocar `usecase.Execute(ctx, input)`.
  3. Mapear erros de domínio para códigos HTTP semânticos (ex: `ErrNotFound` &rarr; 404, validação &rarr; 400, outros &rarr; 500).
  4. Serializar o DTO de saída em JSON para a resposta.

### 6. Composição e Injeção de Dependências (`cmd/*/main.go`)
- Composição manual e explícita no `main.go` ("de fora para dentro"):
  1. Inicializa infraestrutura (banco de dados, clients HTTP).
  2. Instancia repositórios concretos.
  3. Instancia casos de uso passando os repositórios como dependência.
  4. Instancia handlers HTTP passando os casos de uso.
  5. Sobe o servidor HTTP ou dispara a rotina do worker.
- Não utiliza frameworks mágicos de Dependency Injection (como Fx ou Wire) enquanto a árvore for direta e compreensível.

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
