# Backend Testing

Use Go's `testing` package, Testify, Mockery-generated mocks, `httptest`, and
real database integration tests where the repository already uses them. Do not
add a test or mocking framework without a concrete need.

## Test Boundary

Test observable behavior and important failure paths. Choose the smallest test
boundary that provides meaningful confidence:

- Test pure functions directly.
- Test use cases with mocks at external boundaries.
- Test HTTP adapters with `httptest`; never call external APIs.
- Test database repositories against an isolated real database when database
  behavior, SQL mapping, or transactions matter.
- Test schedulers only for scheduling behavior, separate from the use case.

Do not mock pure transformations. Do not simulate database behavior with large
mock hierarchies when an integration test provides the relevant confidence.

## Scenario Structure

Divide test coverage into focused scenarios: successful behavior, meaningful
failure paths, and cancellation or transaction behavior when they are part of
the contract. Use descriptive test and subtest names that state the behavior.

Test code must contain no comments. Do not add `Arrange`, `Act`, or `Assert`
comments. Each scenario must use distinct blank-line blocks in this order:

1. context, mocks, and test data;
2. mock expectations;
3. construction and execution;
4. assertions.

Use table-driven tests only when several scenarios genuinely share the same
structure. Keep each case readable and avoid dense setup or assertion chains.

## Mockery Workflow

Before creating or modifying Go tests, open the affected Go service and run:

```bash
mockery
```

Reuse the generated mocks. Do not edit generated Mockery files manually.

## Assertions and Determinism

Use `require` when a failed assertion makes later checks meaningless. Use
`assert` when independent checks can still provide useful failure information.

Keep tests deterministic. Do not depend on real external services, arbitrary
sleeps, execution order, shared mutable state, or production data.

## Required Validation

After creating or modifying Go tests, run these commands from the affected Go
service:

```bash
go test ./...
go vet ./...
```

Run `go fmt ./...` whenever Go source formatting changes. Run `go test -race
./...` for concurrency-sensitive work and `go test -cover ./...` when coverage
analysis is needed.

Inspect the final diff. Do not weaken or delete a valid test merely to make the
suite pass. Classify failures before fixing them: production bug, test bug,
outdated expectation, environment issue, mock configuration, or dependency
problem.
