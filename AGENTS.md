# Repository Guide

This is Thiago Sousa's portfolio monorepo.

## Repository structure

- `apps/web`: Astro frontend.
- `services/projects_sync`: Go service responsible for synchronizing project data.
- `infra`: infrastructure and local environment configuration.
- `docs`: repository architecture, conventions, and engineering documentation.

Read the nearest `AGENTS.md` before modifying code.

## Architecture

Read `ARCHITECTURE.md` before making architectural changes.

## Boundaries

### Frontend

Codex may autonomously implement and refactor code under:

`apps/web`

Follow:

`apps/web/AGENTS.md`

### Go backend

The Go backend is primarily educational and is implemented by the repository owner.

Do NOT implement production backend behavior unless explicitly requested.

Codex may:

- inspect backend code;
- explain problems;
- review code;
- write or improve tests when explicitly requested;
- identify bugs;
- suggest refactorings.

Codex must NOT autonomously redesign backend architecture.

Follow:

`services/projects_sync/AGENTS.md`

## Validation

Never claim a task is complete without running the validation commands
defined by the nearest AGENTS.md.

Do not weaken or delete tests merely to make validation pass.