# projects_sync Agent Guide

## Purpose

`projects_sync` is a Go service that synchronizes GitHub project data.

This backend is also an educational project for its owner.

Therefore Codex should assist the developer rather than replace the
developer when implementing production Go code.

## Architecture

Respect the existing dependency direction.

Domain must not depend on infrastructure.

Application coordinates use cases.

Infrastructure implements external concerns such as PostgreSQL,
GitHub and Supabase.

Presentation contains entrypoints and scheduling concerns.

Do not introduce new architectural abstractions without an explicit request.

## Production code

Unless explicitly instructed otherwise:

- do not implement business logic;
- do not redesign interfaces;
- do not introduce frameworks;
- do not move packages;
- do not add dependencies.

Instead, explain the change and allow the developer to implement it.

## Tests

Codex may write tests when explicitly requested.

Prefer Go's standard testing facilities and the testing conventions
already present in the repository.

Existing project tools such as Testify and generated Mockery mocks
should be reused rather than replaced.

Tests must exercise behavior, not implementation details.

Never change production behavior merely to make a test pass.

## Required validation

Before declaring Go work complete:

go fmt ./...
go vet ./...
go test ./...

When coverage is relevant:

go test ./... -cover

When race-sensitive/concurrent code is involved:

go test -race ./...