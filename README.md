# ⚡ Thiago Sousa — Portfolio & Engineering Platform

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![Astro Version](https://img.shields.io/badge/Astro-5.x-BC52EE?style=flat&logo=astro&logoColor=white)](https://astro.build/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178C6?style=flat&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18.6-4169E1?style=flat&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Este repositório é um **monorepo de engenharia de software** que abriga meu portfólio profissional e uma plataforma de serviços backend completa. Projetado do zero sem depender de soluções de BaaS (Backend-as-a-Service) fechadas, o projeto atua como laboratório prático para padrões arquiteturais robustos, **Test-Driven Development (TDD)** estrito, concorrência nativa em Go e infraestrutura própria.

---

## 🏛️ Visão Geral da Arquitetura

O ecossistema é organizado em um monorepo limpo e modular:

```text
portfolio/
├── apps/
│   └── web/                   # Frontend Astro 5 + React + TypeScript + Tailwind
├── services/
│   └── app-api/               # Backend modular em Go (API REST, GoAdmin e Sync Worker)
│       ├── cmd/
│       │   ├── api/           # Entrypoint HTTP REST Server & GoAdmin
│       │   └── sync/          # Worker autônomo de sincronização com GitHub
│       ├── internal/
│       │   ├── domain/        # Entidades puras, regras de negócio e contratos (interfaces)
│       │   ├── application/   # Casos de uso e orquestração de domínio
│       │   ├── infrastructure/# Adaptadores SQL (pgx/v5), API do GitHub e GoAdmin
│       │   └── presentation/  # Handlers HTTP, middlewares e roteador Chi v5
│       └── migrations/        # Migrações versionadas para PostgreSQL (golang-migrate)
└── docs/                      # Guias de skills, arquitetura e styleguide do repositório
```

---

## 🛠️ Stack Tecnológica

### Frontend (`apps/web`)
- **[Astro 5](https://astro.build/)**: Framework web focado em conteúdo e performance com arquitetura de ilhas (*Islands Architecture*).
- **[React](https://react.dev/) & [TypeScript](https://www.typescriptlang.org/)**: Componentes interativos de alta fidelidade e tipagem estrita de ponta a ponta.
- **[Tailwind CSS](https://tailwindcss.com/) & [Magic UI](https://magicui.design/)**: Design refinado com animações suaves (`BlurFade`, `FlickeringGrid`, `Dock`).

### Backend (`services/app-api`)
- **[Go (Golang)](https://go.dev/)**: Linguagem compilada de alto desempenho, tipagem forte e concorrência nativa.
- **[Chi Router (v5)](https://github.com/go-chi/chi)**: Roteador HTTP leve, idiomático e compatível com `net/http`.
- **[GoAdmin](https://github.com/GoAdminGroup/go-admin)**: Painel administrativo self-hosted embutido no binário para gestão operacional direta.
- **[Testify](https://github.com/stretchr/testify) & [Mockery](https://github.com/vektra/mockery)**: Testes unitários com assertions expressivas e geração automática de mocks tipados baseados nos contratos de domínio.

### Persistência & Dados
- **[PostgreSQL 18.6](https://www.postgresql.org/)**: Banco de dados relacional com suporte a tipos compostos, arrays nativos e constraints de integridade referencial.
- **[golang-migrate](https://github.com/golang-migrate/migrate)**: Controle de versão do schema de banco via arquivos SQL incrementais (`.up.sql` e `.down.sql`).

---

## 💎 Diferenciais e Padrões de Engenharia

### 1. Clean Architecture com Dependências Inward
As dependências apontam estritamente para dentro:
$$\text{Presentation} \longrightarrow \text{Application} \longrightarrow \text{Domain} \longleftarrow \text{Infrastructure}$$
O Domínio (`internal/domain`) é o núcleo puro: não possui nenhuma dependência externa, imports de banco de dados ou tags de serialização. Todos os contratos de persistência vivem nele (`internal/domain/repositories`).

### 2. Sincronização Inteligente com GitHub (`cmd/sync`)
- Um worker independente consome a API REST do GitHub.
- Executa **upserts não-destrutivos** baseados no `ExternalID` (o identificador único do repositório no GitHub).
- Preserva associações manuais de tecnologias e competências (`skills`) vinculadas via painel operacional sem sofrer sobrescrita durante sincronizações em lote.

### 3. Test-Driven Development (TDD) Estrito
Todo o desenvolvimento de regras de negócio, entidades e casos de uso segue o ciclo clássico **Red-Green-Refactor**:
1. **Red**: Criação prévia do teste unitário cobrindo cenários de sucesso e falha com o mock gerado via `mockery`.
2. **Green**: Implementação mínima e suficiente do caso de uso ou entidade para atingir 100% de aprovação.
3. **Refactor**: Otimização contínua de legibilidade e performance mantendo os testes verdes.
- **Regra de Zero Comentários em Testes**: Os arquivos de teste (`*_test.go`) não contêm comentários. A estrutura AAA (Arrange, Act, Assert) é delimitada estritamente por blocos isolados com linhas em branco.

### 4. Styleguide Go Contemporâneo (Effective Go + Go 1.21+)
- **Compile-Time Interface Checks**: Garantia estática de conformidade de tipos sem custo de execução:
  ```go
  var _ domainusecases.GetAboutUseCase = (*GetAboutUseCaseImpl)(nil)
  ```
- **Prevenção de Memory Loitering**: Manipulação de coleções via pacote nativo `slices` (`slices.ContainsFunc`, `slices.IndexFunc`, `slices.Delete`).
- **Nomes Completos e Sem Abreviações**: Proibição de nomes crípticos no código (`aboutRepository` em vez de `aboutRepo`, `workExperience` em vez de `workExp`).
- **Sentinel Errors & Wrapping**: Tratamento explícito de erros com sentinelas (`ErrDomainEntityValidation`, `ErrNotFound`) e inspeção semântica com `%w` e `errors.Is`.

---

## 🚀 Como Executar o Projeto Localmente

### Pré-requisitos
- [Go 1.24+](https://go.dev/dl/)
- [Node.js 20+](https://nodejs.org/) & [pnpm](https://pnpm.io/)
- [Docker & Docker Compose](https://www.docker.com/)

### 1. Clonar o Repositório
```bash
git clone git@github.com:cthiagoodev/thiagoodev-portfolio.git
cd thiagoodev-portfolio
```

### 2. Subir o Banco de Dados (PostgreSQL 18 & Migrations)
```bash
docker compose up -d
```
> O Docker Compose inicializa a instância do PostgreSQL 18.6 e executa automaticamente todas as migrations via container do `golang-migrate`.

### 3. Executar o Backend (`services/app-api`)
```bash
cd services/app-api
cp .env.example .env
go run cmd/api/main.go
```
Para rodar a suíte de testes do backend:
```bash
go test -v ./...
go vet ./...
```

### 4. Executar o Frontend (`apps/web`)
```bash
cd apps/web
pnpm install
pnpm dev
```
Acesse a aplicação em `http://localhost:4321`.

---

## 📦 Padrão de Commits

O repositório segue estritamente mensagens de commit atômicas e de linha única:
```text
[escopo] Ação executada
```
- `[portfolio-web]`: Modificações na interface e componentes Astro/React (`apps/web`).
- `[app-api]`: Modificações no serviço backend em Go (`services/app-api`).
- `[repo]`: Modificações de orquestração global, documentação ou configurações raiz.

---

## 👨‍💻 Autor

Desenvolvido por **Thiago Sousa**  
- **LinkedIn**: [linkedin.com/in/thiagoodev](https://www.linkedin.com/in/thiagoodev/)  
- **GitHub**: [github.com/cthiagoodev](https://github.com/cthiagoodev)