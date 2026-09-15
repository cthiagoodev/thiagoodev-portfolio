# Testing Strategy

This document defines the testing strategy for the repository.

It describes what should be tested, which test boundaries should be used, and
what constitutes meaningful validation.

For backend conventions, read [`backend.md`](backend.md).

For system architecture, read [`../ARCHITECTURE.md`](../ARCHITECTURE.md).

---

## 1. Purpose

Tests exist to provide confidence in system behavior and detect regressions.

The objective is not:

```text
maximum number of tests
maximum coverage percentage
testing every implementation detail
```

The objective is:

```text
meaningful behavior
        +
important failure paths
        +
architectural boundaries
        +
regression protection
```

Coverage is a diagnostic metric, not the goal of the test suite.

---

## 2. Test Pyramid

Prefer many inexpensive tests and fewer expensive integration tests.

Conceptually:

```text
             ┌─────────────┐
             │ End-to-End  │
             └──────┬──────┘
                    │
          ┌─────────┴─────────┐
          │ Integration Tests │
          └─────────┬─────────┘
                    │
       ┌────────────┴────────────┐
       │       Unit Tests        │
       └─────────────────────────┘
```

The exact ratio is not a target.

Choose the smallest test boundary capable of providing meaningful confidence.

---

## 3. Go Testing Stack

The Go backend currently uses:

- Go `testing`;
- Testify;
- Mockery;
- `httptest`;
- real PostgreSQL integration tests where appropriate.

Do not introduce additional testing frameworks without a concrete reason.

---

## 4. Unit Tests

Unit tests should verify behavior in isolation when isolation provides value.

Good candidates include:

- application use cases;
- mappers;
- deterministic transformations;
- validation logic;
- error propagation.

Pure functions should generally be tested directly.

Do not mock a pure mapper merely because it is possible.

---

## 5. Application Tests

Application use cases coordinate external boundaries.

Their unit tests should focus on orchestration.

For `SyncProjectsUseCase`, the important conceptual flow is:

```text
GitHub
   ↓
Mapper
   ↓
Repository
   ↓
Supabase
```

Tests should verify meaningful outcomes and error propagation.

They should not reproduce the implementation tests belonging to each
infrastructure adapter.

---

## 6. Mocking

Mocks should represent meaningful external boundaries.

Examples include:

```text
ProjectsRepository
GithubService
SupabaseService
```

The project uses Mockery-generated Testify mocks.

Generated mocks should be reused.

Do not manually edit generated mocks.

Avoid mocking every internal object.

Excessive mocking couples tests to implementation structure and makes
refactoring unnecessarily difficult.

---

## 7. HTTP Integration Tests

HTTP adapters should use Go's `httptest` facilities.

Tests for GitHub infrastructure must not call the real GitHub API.

The test owns the HTTP server:

```text
Test
 │
 ├── httptest.Server
 │        ↑
 │        │ HTTP
 │        │
 └── GithubService
```

This allows tests to control:

- response status;
- response body;
- headers;
- malformed responses;
- delays;
- request validation.

External network availability must not determine whether the test suite passes.

---

## 8. PostgreSQL Repository Tests

Concrete PostgreSQL repositories should primarily be tested using integration
tests against PostgreSQL.

Avoid building large pgx mock hierarchies to simulate database behavior.

A repository test should verify actual behavior such as:

- SQL execution;
- row mapping;
- PostgreSQL types;
- constraints;
- bulk insertion;
- transaction semantics;
- rollback behavior.

The database used for testing must not be a production database.

---

## 9. Transaction Tests

Transaction tests should verify atomicity rather than merely checking returned
errors.

For project replacement:

```text
existing data
     │
     ▼
BEGIN
     │
     ├── DELETE
     │
     └── COPY
           │
           └── failure
     │
ROLLBACK
```

The important invariant is:

> If replacement fails, the previously committed project state remains intact.

A meaningful test therefore verifies the database state after failure.

---

## 10. Context Tests

Context cancellation and deadlines should be tested when they are meaningful
to the contract.

Examples include:

- HTTP request cancellation;
- database operation cancellation;
- scheduler shutdown;
- long-running operations.

Do not create context tests solely to increase coverage.

---

## 11. Scheduler Tests

Scheduler tests should focus on scheduling responsibilities.

They should not duplicate synchronization tests.

The boundary is:

```text
Scheduler
    ↓
SyncProjectsUseCase
```

The scheduler is responsible for triggering the use case.

The use case is responsible for synchronization behavior.

---

## 12. Error Paths

Important failure paths deserve explicit tests.

Examples include:

```text
GitHub unavailable
GitHub invalid response
database operation failure
transaction rollback
Supabase synchronization failure
context cancellation
```

Not every theoretical failure needs its own test.

Prioritize failures that affect correctness, data consistency, or important
system behavior.

---

## 13. Table-Driven Tests

Use table-driven tests when several scenarios share the same test structure.

Example:

```go
tests := []struct {
    name string
    // input
    // expected
}{
    // cases
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // test
    })
}
```

Do not force table-driven tests when separate tests would communicate intent
more clearly.

---

## 14. Assertions

Use `require` when failure means the rest of the test cannot provide meaningful
information.

Example:

```go
require.NoError(t, err)
```

Use `assert` for independent assertions where continuing the test remains
useful.

Example:

```go
assert.Equal(t, expected.Name, actual.Name)
assert.Equal(t, expected.URL, actual.URL)
```

Assertions should make failures understandable.

---

## 15. Test Naming

Test names should communicate behavior.

Prefer names such as:

```text
TestSyncProjectsReturnsErrorWhenGithubFails

TestResetAndCreateAllRollsBackWhenCopyFails

TestGithubServiceReturnsErrorForUnexpectedStatus
```

over vague names such as:

```text
TestExecute

TestRepository

TestError
```

The name should help identify what behavior broke without reading the entire
test.

---

## 16. Determinism

Tests should produce the same result when executed repeatedly under the same
conditions.

Avoid:

- real external APIs;
- arbitrary sleeps;
- dependence on execution order;
- shared mutable global state;
- dependence on production data.

Control time and external behavior when they materially affect the test.

---

## 17. Race Detection

Concurrency-sensitive Go code should be validated with the race detector.

```bash
go test -race ./...
```

This is particularly relevant when introducing:

- goroutines;
- shared mutable state;
- asynchronous jobs;
- concurrent scheduling behavior.

It does not need to run after every trivial change.

---

## 18. Coverage

Coverage can be inspected with:

```bash
go test -cover ./...
```

Coverage should help identify important behavior that lacks tests.

Do not create meaningless tests solely to increase the percentage.

A line being executed does not prove that its behavior was meaningfully
verified.

---

## 19. Standard Validation

The baseline Go validation loop is:

```bash
go fmt ./...
go vet ./...
go test ./...
```

For diagnosis:

```bash
go test -v ./...
```

For concurrency-sensitive work:

```bash
go test -race ./...
```

For coverage analysis:

```bash
go test -cover ./...
```

Commands should be executed from the relevant Go module.

---

## 20. Frontend Testing

Frontend tests and validation should follow the tooling actually configured
inside `apps/web`.

Before running frontend validation, inspect:

```text
apps/web/package.json
```

Do not assume that a script such as:

```text
npm test
npm run lint
npm run check
```

exists without verifying it.

Frontend changes should at minimum be validated using the static checks, tests,
and production build commands available in the project.

---

## 21. Generated Tests

Tests produced with coding-agent assistance are held to the same standard as
manually written tests.

An agent-generated test must:

- test meaningful behavior;
- follow existing conventions;
- avoid unnecessary mocks;
- compile;
- be executed;
- produce a verified result.

Generated test code is not considered complete until it has actually been run.

---

## 22. Failing Tests

A failing test does not automatically mean production code should change.

First determine whether the failure represents:

```text
production bug
test bug
outdated expectation
environment failure
missing dependency
incorrect mock configuration
database setup problem
race condition
```

Fix the layer responsible for the failure.

Never weaken a valid test simply to make the suite green.

---

## 23. Testing Principle

The preferred testing question is not:

> How can we test this function?

It is:

> What behavior would hurt us if it silently stopped working?

Tests should protect those behaviors first.