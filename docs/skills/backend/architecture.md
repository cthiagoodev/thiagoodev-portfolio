# Backend Architecture

The backend lives under `services/app-api` and follows a modular multi-binary architecture in Go.

## Overview

`services/app-api` produces two distinct binaries from the same codebase:
- `cmd/api`: The main HTTP REST server powered by `go-chi/chi/v5` serving the public endpoints (consumed by the Astro frontend under `/api/*`) and hosting the `go-admin` panel (under `/admin/*`).
- `cmd/sync`: The GitHub repository synchronization worker, designed to run as an independent scheduled process/container.

Both binaries share business logic, domain models, and infrastructure implementations located inside `internal/`.

## Directory Structure

```text
services/app-api/
├── cmd/
│   ├── api/             # Entrypoint for HTTP server & GoAdmin
│   └── sync/            # Entrypoint for GitHub sync batch job
├── internal/
│   ├── domain/          # Core entities and behavioral contracts
│   ├── application/     # Use cases orchestrating domain logic
│   ├── infrastructure/  # Database (pgx/v5), GitHub client, GoAdmin setup
│   └── presentation/    # HTTP routes, handlers (Chi v5), middlewares
└── migrations/          # SQL migrations managed by golang-migrate
```

## Dependency Direction

Dependencies point strictly inward:

```text
Presentation → Application → Domain
Infrastructure → Application/Domain contracts
```

- **Domain**: Contains pure entities and interfaces. Must not depend on infrastructure technologies, database drivers, third-party libraries, HTTP frameworks, or scheduler types.
- **Application**: Contains use cases coordinating behavior between domain and repositories. Free from HTTP details, SQL strings, or framework bindings.
- **Infrastructure**: Implements persistence (PostgreSQL with `pgx`), external APIs (GitHub HTTP client), and administrative engines (GoAdmin).
- **Presentation**: Sets up Chi routers, routes, request/response serialization, and delegates execution to application use cases.

## Design Rules

- Use explicit constructor injection for all dependencies.
- Define small, focused interfaces where behavioral abstraction is necessary.
- Do not introduce unnecessary frameworks or external dependencies without explicit requirement.
- Propagate `context.Context` across all I/O boundaries. Do not store contexts in struct fields.
- Keep database transactions contained inside the infrastructure layer.
- Return errors to the caller with meaningful context.

## PostgreSQL 18+ Volume & Storage Standards

The project uses PostgreSQL 18 (`postgres:18.6`).

- **Volume Mount Path**: In `docker-compose.yaml`, the volume MUST mount to `/var/lib/postgresql`, never `/var/lib/postgresql/data`:
  ```yaml
  volumes:
    - postgres_data:/var/lib/postgresql
  ```
- **Technical Rationale & Mechanics**:
  - Starting in PostgreSQL 18, the official Docker images switched to a layout compatible with `pg_ctlcluster` (storing cluster data in major-version subdirectories rather than directly in `/data`).
  - This architecture exists to support zero-copy major version upgrades via `pg_upgrade --link`.
  - In Linux, *hard links* cannot be created across mount point (filesystem) boundaries. By mounting the volume at the parent `/var/lib/postgresql`, both current and future version directories reside within the same filesystem mount, allowing seamless `pg_upgrade --link`.
  - Attempting to mount `/var/lib/postgresql/data` in PostgreSQL 18+ triggers an intentional startup check failure.
