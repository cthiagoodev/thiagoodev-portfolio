---
name: testing
description: Test and validate the Go backend, especially services/projects_sync. Use when creating, reviewing, debugging, or improving Go tests.
---

# Testing Skill

Use this skill when working with tests in the Go backend.

Before testing backend behavior, read:

- `ARCHITECTURE.md`
- `docs/backend.md`
- `docs/testing.md`, if available
- the production code being tested
- nearby existing tests

Do not infer behavior from test names alone.

---

## Purpose

Tests are part of the repository feedback loop.

The goal is not to maximize the number of tests or coverage percentage.

The goal is to verify meaningful behavior and detect regressions while keeping
tests readable, deterministic, and maintainable.

---

## Scope

The primary Go service is:

```text
services/projects_sync
```

The current testing stack includes:

- Go `testing`
- Testify
- Mockery-generated Testify mocks
- `httptest`
- PostgreSQL integration tests where appropriate

Reuse the existing testing stack.

Do not introduce another testing or mocking framework unless explicitly
requested.

---

## Core Rules

Always understand the behavior before writing the test.

Test observable behavior rather than internal implementation details.

Do not change production code merely to satisfy an incorrect test.

Do not delete, weaken, skip, or ignore a valid failing test just to make the
suite pass.

Do not make real network requests from unit tests.

Prefer deterministic tests.

Avoid arbitrary sleeps.

Do not mock pure functions unnecessarily.

Do not mock PostgreSQL behavior extensively when a real database integration
test would provide more confidence.

---

# Workflow

When asked to create or modify tests, follow this process.

## 1. Inspect the target

Read the production code being tested.

Identify:

- inputs;
- outputs;
- errors;
- side effects;
- dependencies;
- important branches;
- cancellation behavior;
- transaction boundaries, when applicable.

Do not start generating tests before understanding these behaviors.

---

## 2. Inspect existing tests

Look for nearby tests and follow the conventions already used by the project.

Reuse:

- naming conventions;
- Testify usage;
- generated mocks;
- test helpers;
- fixture patterns;
- package organization.

Do not create a parallel testing style without a concrete reason.

---

## 3. Identify the test boundary

Choose the appropriate type of test.

### Pure logic

Test directly.

Examples:

- mappers;
- transformations;
- validation;
- deterministic helper functions.

Do not mock pure functions unless there is a concrete architectural reason.

### Application use case

Use unit tests when external boundaries can be replaced by existing mocks.

Examples:

- repository;
- GitHub service;
- Supabase service.

Test orchestration and observable behavior.

Do not assert private implementation details.

### HTTP infrastructure

Use `httptest`.

Do not call the real GitHub API.

### PostgreSQL repository

Prefer integration tests against a real PostgreSQL instance.

Test actual SQL and transaction behavior rather than reproducing pgx behavior
through mocks.

### Scheduler

Test scheduling-specific behavior separately from synchronization behavior.

The scheduler should delegate synchronization to the use case.

Do not duplicate use-case tests inside scheduler tests.

---

# Go Test Structure

Prefer idiomatic Go tests.

A typical test follows:

```go
func TestSomething(t *testing.T) {
    // Arrange

    // Act

    // Assert
}
```

Comments such as Arrange, Act, and Assert are optional.

Use them only when they improve readability.

For multiple related scenarios, table-driven tests are encouraged when they
make the test easier to understand.

Example:

```go
func TestMapper(t *testing.T) {
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
}
```

Do not force table-driven tests when individual tests would be clearer.

---

# Assertions

Use Testify consistently with the existing codebase.

Use `require` when continuing the test after failure would make the remaining
assertions meaningless.

Example:

```go
require.NoError(t, err)
require.NotNil(t, result)
```

Use `assert` when multiple independent assertions can still provide useful
information.

Example:

```go
assert.Equal(t, expected.Name, result.Name)
assert.Equal(t, expected.URL, result.URL)
```

Do not create large assertion chains that obscure which behavior failed.

---

# Mocking

The project uses Mockery-generated Testify mocks.

Generated mocks live under the project's test mock structure.

Do not manually modify generated Mockery files.

Do not create handwritten mocks when an appropriate generated mock already
exists.

When an interface changes, regenerate its mock using the project's Mockery
configuration.

Mocks should represent external boundaries, not every internal function.

Good candidates include:

```text
ProjectsRepository
GithubService
SupabaseService
```

Pure transformations generally should not be mocked.

---

## Mock Expectations

Configure only expectations relevant to the behavior being tested.

Example:

```go
githubService.
    EXPECT().
    FetchRepositories(mock.Anything).
    Return(repositories, nil).
    Once()
```

Avoid over-specifying unrelated implementation details.

Use `.Once()` when the number of calls is part of the expected behavior.

Do not require an exact call count when the count is irrelevant to the
contract.

---

# Context

Production operations commonly receive:

```go
context.Context
```

Tests should pass an explicit context.

For normal tests:

```go
ctx := context.Background()
```

When testing cancellation:

```go
ctx, cancel := context.WithCancel(context.Background())
cancel()
```

When testing deadlines:

```go
ctx, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()
```

Only test cancellation or deadlines when they are relevant to the behavior.

Do not add context tests merely to increase coverage.

---

# HTTP Tests

HTTP infrastructure must be tested with Go's `httptest` facilities.

Typical structure:

```go
server := httptest.NewServer(
    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // validate request

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

        _, _ = w.Write([]byte(`...`))
    }),
)
defer server.Close()
```

Point the GitHub service at the test server.

Never allow these tests to fall back to:

```text
api.github.com
```

Tests should control the response completely.

Relevant scenarios may include:

- successful response;
- non-success HTTP status;
- invalid response body;
- request cancellation.

Only implement scenarios relevant to the current contract.

---

# PostgreSQL Integration Tests

Concrete repository behavior should preferably be tested against PostgreSQL.

Do not attempt to recreate PostgreSQL semantics with mocks.

Repository integration tests should focus on behavior such as:

- inserting data;
- reading data;
- mapping database rows;
- replacing data;
- constraints;
- transaction behavior;
- rollback on failure.

Tests should use isolated test data.

They must not depend on production databases.

---

## Transaction Testing

Atomic repository operations require failure-path testing.

For an operation conceptually equivalent to:

```text
BEGIN

DELETE FROM projects

COPY new projects

COMMIT
```

the important invariant is:

```text
Either the entire replacement succeeds,
or the previous state remains intact.
```

A useful integration test therefore follows:

```text
seed existing projects
        ↓
execute replacement that fails
        ↓
query database
        ↓
verify original projects still exist
```

This verifies transaction behavior rather than merely verifying that an error
was returned.

---

# Testing SyncProjectsUseCase

The synchronization use case coordinates several boundaries.

Its tests should focus on application behavior.

Relevant scenarios may include:

```text
GitHub returns error
        ↓
Execute returns error
```

```text
GitHub returns no repositories
        ↓
Execute returns the expected error
```

```text
GitHub succeeds
        ↓
projects are mapped
        ↓
repository replacement succeeds
        ↓
projects are retrieved
        ↓
Supabase replacement succeeds
        ↓
Execute succeeds
```

And failure propagation from meaningful boundaries.

Do not reproduce tests already owned by the GitHub adapter, PostgreSQL
repository, mapper, or Supabase adapter.

---

# Failure Diagnosis

When a test fails, do not immediately modify code.

First classify the failure.

Possible categories:

```text
Production bug
Test bug
Outdated test
Environment problem
Infrastructure dependency unavailable
Incorrect mock expectation
Race/concurrency problem
```

Inspect the failure before deciding what should change.

If the test exposes a likely production bug, report it before changing
production behavior unless explicitly asked to implement the fix.

---

# Validation

After modifying Go tests, run formatting:

```bash
go fmt ./...
```

Run static analysis:

```bash
go vet ./...
```

Run the test suite:

```bash
go test ./...
```

For more detailed output when diagnosing failures:

```bash
go test -v ./...
```

For coverage when requested or useful:

```bash
go test -cover ./...
```

For concurrency-sensitive code:

```bash
go test -race ./...
```

Do not claim that tests pass unless the relevant command was actually executed
successfully.

---

# Working Directory

Before running Go commands, identify the Go module containing the target code.

For `projects_sync`, commands may need to run from:

```text
services/projects_sync
```

Inspect `go.mod` rather than assuming the repository root is a Go module.

---

# Generated Mocks

If mocks need regeneration, use the Mockery configuration already present in
the project.

Before running Mockery, verify that it is available:

```bash
mockery --version
```

Do not manually recreate generated mock files if Mockery is unavailable.

Report the missing tool instead.

After generation, inspect the diff to ensure only expected mocks changed.

---

# Final Review

Before declaring testing work complete:

1. inspect the final diff;
2. ensure production behavior was not changed accidentally;
3. ensure generated files were not manually edited;
4. ensure tests do not access real external services;
5. run the relevant validation commands;
6. report what was tested;
7. report any test that could not be executed.

A test implementation is not complete merely because test code was generated.

It must be executed and its result verified.