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

- `apps`: frontend applications.
- `services`: independently executable backend services.
- `infra`: infrastructure configuration.
- `docs`: architecture, agent skills, and engineering documentation.

## Work Boundaries

### Frontend: `apps`

Agents may implement and refactor frontend code when requested. Preserve the
existing visual identity and prefer incremental changes.

- Do not implement backend logic in the frontend.
- Do not change backend contracts without explicit approval.
- Consume portfolio data from Supabase; do not query GitHub directly.
- Do not duplicate backend synchronization behavior.
- Before changing code, inspect `package.json`, `astro.config.*`, nearby
  components, styles, and existing data access patterns.

Follow [`docs/skills/frontend/`](docs/skills/frontend/).

### Backend: `services`

The Go backend is primarily educational and is implemented by the repository
owner. Unless production behavior is explicitly requested, agents may inspect
and explain code, review it, identify bugs, suggest refactorings, and write or
improve tests when explicitly requested.

Do not autonomously implement production business logic, redesign interfaces
or architecture, introduce frameworks or dependencies, or move packages.
When a production change is needed but was not requested, explain the issue and
let the owner implement it.

Backend tests must verify behavior, use the existing Go testing stack and
generated Mockery mocks, and never change production behavior merely to make a
test pass. Before creating or changing Go tests, run `mockery` from the Go
service being changed. Afterward, run `go test ./...` and `go vet ./...` from
that same service.

Test code must contain no comments. Divide behavior into focused scenarios,
and use descriptive names and distinct blank-line blocks for setup,
expectations, execution, and assertions. Do not add `Arrange`, `Act`, or
`Assert` comments.

Follow [`docs/skills/backend/`](docs/skills/backend/).

### Infrastructure: `infra`

Infrastructure is read-only for agents. Agents may inspect it, explain how it
works, review configuration, and answer questions. They must not create,
modify, delete, execute, reset, update, or migrate anything under `infra`
unless the repository owner gives an explicit writing or operational command.

The infrastructure is a self-hosted Supabase deployment running with Docker.
It is educational, like the Go backend: explain its original Supabase
self-hosted architecture and existing migrations so the owner can learn how it
works internally.

Follow [`docs/skills/infrastructure/`](docs/skills/infrastructure/), including
the documentation of the component being changed.

## Architecture Changes

Do not change system boundaries, persistence ownership, dependency direction,
service responsibilities, or frontend/backend contracts without explicit
instruction. Document an approved boundary change in the relevant
`docs/skills/<area>/architecture.md` document before or alongside
implementation.

## Validation and Reporting

Before declaring implementation work complete, run the validation commands
specified in the relevant skill and project configuration. For Go test changes,
run these from the affected Go service:

```bash
go test ./...
go vet ./...
```

For frontend changes, inspect `apps/web/package.json` and run the configured
formatter or linter, tests, and production build. Do not assume a script exists
without checking. Do not run infrastructure tests or operational commands
unless the owner explicitly requests them.

Do not weaken or delete valid tests to make validation pass. Report the
commands run and any validation that could not be completed.
