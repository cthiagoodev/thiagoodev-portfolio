# Repository Agent Guide

This is Thiago Sousa's portfolio monorepo. This root file is the single
repository-wide agent instruction entry point.

## Before Working

1. Read this file before changing repository content.
2. Inspect the target code, its nearby tests, and current Git changes.
3. Read the relevant area guide under [`docs/skills/`](docs/skills/README.md)
   before changing component responsibilities, contracts, system boundaries,
   tests, or operational configuration.

Keep repository-wide agent rules here. Do not add nested `AGENTS.md` files
without an explicit request to restore scoped instructions.

## Repository Structure

- `apps`: frontend applications (`apps/web` with Astro & TypeScript).
- `services`: independently executable backend services (`services/app-api`).
- `docs`: architecture, agent skills, and engineering documentation.

## Work Boundaries

### Frontend: `apps`

Agents have full autonomy to implement, refactor, and polish frontend code under `apps/web` when requested. Preserve the existing visual identity and prefer incremental changes.

- Do not implement backend logic in the frontend.
- Do not change backend contracts without explicit approval.
- Consume portfolio data from the backend REST API (`app-api`); do not query external sources directly.
- Do not duplicate backend synchronization behavior.
- Before changing code, inspect `package.json`, `astro.config.*`, nearby components, styles, and existing data access patterns.

Follow [`docs/skills/frontend/`](docs/skills/frontend/).

### Backend: `services`

The active backend is `services/app-api`, a modular Go service providing the HTTP REST API and GoAdmin panel (`cmd/api`) and the background synchronization worker (`cmd/sync`), backed directly by PostgreSQL 18.

The Go backend is strictly educational and is implemented by the repository owner. The agent must NEVER autonomously implement production business logic, redesign interfaces or architecture, introduce frameworks or dependencies, or move packages. The backend is always pair-programmed alongside the owner.

- When a production change is needed but was not requested, explain the issue and let the owner implement it.
- When explicitly requested to write or refactor backend code, agents must strictly adhere to the Go styleguide ([`docs/skills/backend/styleguide.md`](docs/skills/backend/styleguide.md)) and accompany code changes with didactic explanations of idioms, Go runtime mechanics, and architectural design decisions.
- Follow Test-Driven Development (TDD) strictly: write tests first (Red), implement minimum code (Green), and refactor (Refactor).
- Backend tests must verify behavior, use the existing Go testing stack and generated Mockery mocks, and never change production behavior merely to make a test pass. Before creating or changing Go tests, run `mockery` from the Go service being changed. Afterward, run `go test ./...` and `go vet ./...` from that same service.
- Test code must contain no comments. Divide behavior into focused scenarios, and use descriptive names and distinct blank-line blocks for setup, expectations, execution, and assertions. Do not add `Arrange`, `Act`, or `Assert` comments.

Follow [`docs/skills/backend/`](docs/skills/backend/).

### Git Operations

Follow [`docs/skills/git/`](docs/skills/git/SKILL.md).

- Commit messages must follow the single-line format: `[scope] Action performed` (e.g., `[portfolio-web] Cria arquivo Dockerfile` ou `[app-api] Adiciona entidade Project`).
- **Frontend (`apps/web`)**: Agents have autonomy to commit and push changes directly to `main`.
- **Services (`services/`)**: Agents must NEVER commit, push, or run mutating Git commands autonomously. Only run Git commands in this area upon explicit user request and after user review.

### Educational & Didactic Mode

Follow [`docs/skills/educator/`](docs/skills/educator/SKILL.md).

When the repository owner requests didactic explanations or asks "what happens" and "why it happens", provide clear mental models and analogies, line-by-line dissections, and separate runtime mechanics from architectural design decisions across all repository technologies (Docker, Go, TypeScript/Astro, PostgreSQL, Linux, networking, etc.).

## Architecture Changes

Do not change system boundaries, persistence ownership, dependency direction, service responsibilities, or frontend/backend contracts without explicit instruction. Document an approved boundary change in the relevant `docs/skills/<area>/architecture.md` document before or alongside implementation.

## Validation and Reporting

Before declaring implementation work complete, run the validation commands specified in the relevant skill and project configuration. For Go test changes, run these from the affected Go service:

```bash
go test ./...
go vet ./...
```

For frontend changes, inspect `apps/web/package.json` and run the configured formatter or linter, tests, and production build. Do not assume a script exists without checking.

Do not weaken or delete valid tests to make validation pass. Report the commands run and any validation that could not be completed.
