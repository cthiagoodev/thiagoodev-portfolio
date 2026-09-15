# Backend Engineering Guide

This document defines the engineering conventions for backend development in
this repository.

For system-wide architectural boundaries, read [`../ARCHITECTURE.md`](../ARCHITECTURE.md).

---

## 1. Scope

Backend services live under:

```text
services/
```

The current backend service is:

```text
services/projects_sync
```

`projects_sync` is responsible for synchronizing portfolio project information
from external sources into the portfolio data stores.

---

## 2. Language

Backend services are primarily implemented in Go.

Prefer idiomatic Go and the standard library when they provide an appropriate
solution.

Do not reproduce patterns from Java, Spring, or other object-oriented
frameworks when Go provides a simpler model.

Abstractions should solve concrete problems.

Avoid introducing abstractions only because they are common in other
ecosystems.

---

## 3. Service Structure

The current `projects_sync` service follows:

```text
services/projects_sync/
├── internal/
│   ├── application/
│   │   └── usecases/
│   │
│   ├── domain/
│   │   ├── entities/
│   │   ├── repositories/
│   │   └── usecases/
│   │
│   ├── infrastructure/
│   │   ├── github/
│   │   ├── repositories/
│   │   └── supabase/
│   │
│   ├── presentation/
│   │   └── scheduler/
│   │
│   └── test/
│       └── mocks/
│
├── migrations/
├── go.mod
└── go.sum
```

The structure exists to establish clear dependency and responsibility
boundaries.

Do not create additional layers or packages without a concrete reason.

---

## 4. Domain

The domain contains the core types and contracts used by the service.

It must remain independent from infrastructure implementations.

Domain code must not depend directly on technologies such as:

- pgx;
- PostgreSQL drivers;
- Supabase implementations;
- HTTP clients;
- GitHub API implementations;
- cron implementations.

Domain interfaces should describe behavior required by the application.

Example:

```go
type ProjectsRepository interface {
    GetAll(ctx context.Context) ([]entities.Project, error)
    ResetAndCreateAll(ctx context.Context, projects []entities.Project) error
}
```

The contract describes what persistence capability is required.

It does not describe how PostgreSQL performs that operation.

---

## 5. Application

Application packages contain use cases.

Use cases coordinate application behavior.

For example, project synchronization conceptually performs:

```text
Fetch repositories
        ↓
Map repositories
        ↓
Persist projects
        ↓
Read persisted projects
        ↓
Synchronize portfolio data
```

A use case determines the sequence and application rules.

It should not contain:

- SQL;
- PostgreSQL transaction implementation;
- HTTP request implementation;
- cron expressions;
- infrastructure configuration.

---

## 6. Infrastructure

Infrastructure contains concrete implementations for communication with
external systems.

Current infrastructure includes:

```text
GitHub API
PostgreSQL
Supabase
```

Examples include:

```text
GithubService implementation
ProjectsRepository PostgreSQL implementation
SupabaseService implementation
```

Infrastructure packages may depend on external libraries required to implement
these integrations.

Those dependencies should not unnecessarily propagate into the domain.

---

## 7. Presentation

Presentation contains entrypoints into application behavior.

The current scheduler is an example.

Its responsibility is:

```text
schedule
    ↓
invoke use case
    ↓
handle execution result
```

It must not implement synchronization behavior itself.

---

## 8. Dependency Injection

Dependencies should be explicit.

Prefer constructor injection:

```go
func NewService(
    repository Repository,
    client Client,
) *Service {
    return &Service{
        repository: repository,
        client:     client,
    }
}
```

Avoid hidden global dependencies.

Constructors should make required dependencies visible.

---

## 9. Interfaces

Interfaces represent behavioral boundaries.

Create an interface when there is a meaningful reason for consumers to depend
on behavior rather than an implementation.

Do not create interfaces simply because a concrete type exists.

Prefer small interfaces.

The consumer of an abstraction should generally determine the behavior it
requires.

---

## 10. Context

Operations involving I/O or operation lifetime should accept
`context.Context`.

Convention:

```go
func Operation(ctx context.Context, ...) error
```

Context should normally be the first parameter.

Propagate the incoming context through the call chain.

Example:

```text
Scheduler
    ↓ ctx
Use Case
    ↓ ctx
Repository / Service
    ↓ ctx
PostgreSQL / HTTP
```

Do not replace an incoming context with `context.Background()`.

Do not store context inside long-lived structs.

Context represents the lifetime of an operation, not application state.

---

## 11. PostgreSQL

PostgreSQL access currently uses pgx.

Use context-aware pgx operations.

Examples:

```go
pool.Query(ctx, ...)
pool.Exec(ctx, ...)
pool.CopyFrom(ctx, ...)
```

Prefer APIs appropriate to the operation being performed.

Use `Exec` for commands that do not return rows.

Use query APIs when rows are expected.

Use PostgreSQL bulk mechanisms such as COPY when appropriate for bulk
persistence.

---

## 12. Transactions

Transactions should represent logical atomic operations.

For example, replacing all projects consists of:

```text
DELETE existing projects
COPY new projects
```

These operations must execute within the same transaction.

Conceptually:

```text
BEGIN

DELETE
COPY

COMMIT
```

On failure:

```text
BEGIN

DELETE
COPY → ERROR

ROLLBACK
```

Transaction ownership belongs to the infrastructure layer when the transaction
exists purely to implement a persistence operation.

Do not expose `pgx.Tx` through domain contracts solely to coordinate SQL.

---

## 13. Errors

Errors should be returned to the layer capable of deciding what they mean.

Never silently discard an error.

Avoid logging the same error at every layer.

Prefer returning errors upward until reaching an appropriate application
boundary.

Wrap errors when additional context makes diagnosis materially easier.

Do not expose driver-specific errors through domain contracts unless the
application genuinely needs that distinction.

---

## 14. Collections

Slices should normally be passed directly:

```go
func Create(projects []Project)
```

rather than through pointers:

```go
func Create(projects *[]Project)
```

A slice already contains a reference to its backing storage together with
length and capacity metadata.

Return a new slice when an operation needs to change the caller-visible slice
descriptor.

---

## 15. External HTTP Services

External HTTP integrations belong to infrastructure.

Use explicit `http.Client` dependencies rather than hidden global clients when
testability or configuration matters.

Requests participating in an application operation should use the operation's
context.

Example:

```go
req, err := http.NewRequestWithContext(
    ctx,
    http.MethodGet,
    url,
    nil,
)
```

Responses must have their bodies closed after a successful request.

---

## 16. Dependencies

Add external dependencies only when they solve a concrete problem better than
the standard library or existing project dependencies.

Before adding a dependency, consider:

- whether Go already provides the required functionality;
- whether the project already has an appropriate dependency;
- maintenance status;
- API stability;
- complexity introduced.

Do not introduce frameworks solely to impose structure already expressed by
the repository architecture.

---

## 17. Generated Code

Generated code must not be manually edited unless explicitly documented
otherwise.

This currently includes Mockery-generated mocks.

When the source interface changes, regenerate the corresponding generated
artifacts.

---

## 18. Formatting and Static Analysis

Go code should remain compatible with standard Go tooling.

Baseline validation:

```bash
go fmt ./...
go vet ./...
go test ./...
```

Run commands from the Go module containing the code being validated.

For `projects_sync`, this is normally:

```text
services/projects_sync
```

---

## 19. Engineering Principle

Prefer simple Go code with explicit dependencies and clear responsibilities.

Do not optimize the architecture for hypothetical future requirements.

The preferred progression is:

```text
understand requirement
        ↓
implement simplest appropriate solution
        ↓
test
        ↓
observe
        ↓
refactor when evidence justifies it
```

Architecture should evolve because the system requires it, not because a
pattern could theoretically be introduced.