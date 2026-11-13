# Repository Agent Guide

This is Thiago Sousa's portfolio monorepo. This root file is the single
repository-wide agent instruction entry point.

## Before Working

1. Read this file before changing repository content.
2. Inspect the target code, its nearby tests, and current Git changes.
3. Read the relevant area skill under [`.agents/skills/`](.agents/skills/)
   before changing component responsibilities, contracts, system boundaries,
   tests, or operational configuration.

Keep repository-wide agent rules here. Do not add nested `AGENTS.md` files
without an explicit request to restore scoped instructions.

## Repository Structure

- `apps`: frontend applications (`apps/web` with Astro & TypeScript).
- `services`: independently executable backend services (`services/api` with Java 25 & Spring Boot 4).
- `.agents/skills`: architecture, agent skills, and engineering documentation.

## Work Boundaries

### Frontend: `apps`

Agents have full autonomy to implement, refactor, and polish frontend code under `apps/web` when requested. Preserve the existing visual identity and prefer incremental changes.

- Do not implement backend logic in the frontend.
- Do not change backend contracts without explicit approval.
- Consume portfolio data from the backend REST API (`services/api`); do not query external sources directly.
- Do not duplicate backend synchronization behavior.
- Before changing code, inspect `package.json`, `astro.config.*`, nearby components, styles, and existing data access patterns.

Follow [`.agents/skills/frontend/`](.agents/skills/frontend/SKILL.md).

### Backend: `services`

The active backend is `services/api`, a layered Java 25 and Spring Boot 4.1+ service providing the HTTP REST API (`spring-boot-starter-webmvc`), JPA persistence (`spring-boot-starter-data-jpa`), and Flyway database migrations (`spring-boot-starter-flyway`), backed by PostgreSQL 18.

The Java/Spring Boot backend is strictly educational and is implemented by the repository owner. The agent must NEVER autonomously implement production business logic, redesign interfaces or architecture, introduce dependencies, or move packages. The backend is always pair-programmed alongside the owner.

- When a production change is needed but was not requested, explain the issue and let the owner implement it.
- When explicitly requested to write or refactor backend code, agents must strictly adhere to the Java & Spring Boot styleguide ([`.agents/skills/backend/styleguide.md`](.agents/skills/backend/styleguide.md)) and accompany code changes with didactic explanations of Java 25 idioms, JVM/Spring runtime mechanics, and architectural design decisions.
- Follow Test-Driven Development (TDD) strictly: write unit tests first (Red), implement minimum code (Green), and refactor (Refactor).
- Backend tests must verify behavior using JUnit 5, Mockito, and Spring test slices (`@WebMvcTest`), and never change production behavior merely to make a test pass. After creating or changing backend code or tests, run `./gradlew test` from `services/api`.
- Test code (`*Test.java`) must contain **no comments**. Divide behavior into focused scenarios, and use descriptive names and distinct blank-line blocks for setup, execution, and assertions. Do not add `// Arrange`, `// Act`, or `// Assert` comments.

Follow [`.agents/skills/backend/`](.agents/skills/backend/SKILL.md).

### Git Operations

Follow [`.agents/skills/git/`](.agents/skills/git/SKILL.md).

- Commit messages must follow the single-line format: `[scope] Action performed` (e.g., `[portfolio-web] Cria arquivo Dockerfile` or `[api] Adiciona entidade Project`).
- **Frontend (`apps/web`)**: Agents have autonomy to commit and push changes directly to `main`.
- **Services (`services/`)**: Agents must NEVER commit, push, or run mutating Git commands autonomously. Only run Git commands in this area upon explicit user request and after user review.

### Educational & Didactic Mode

Follow [`.agents/skills/educator/`](.agents/skills/educator/SKILL.md).

When the repository owner requests didactic explanations or asks "what happens" and "why it happens", provide clear mental models and analogies, line-by-line dissections, and separate runtime mechanics from architectural design decisions across all repository technologies (Java 25, Spring Boot 4, JPA/Hibernate, Flyway, Docker, TypeScript/Astro, PostgreSQL, Linux, networking, etc.).

## Architecture Changes

Do not change system boundaries, persistence ownership, dependency direction, service responsibilities, or frontend/backend contracts without explicit instruction. Document an approved boundary change in the relevant `.agents/skills/<area>/architecture.md` document before or alongside implementation.

## Validation and Reporting

Before declaring implementation work complete, run the validation commands specified in the relevant skill and project configuration. For backend changes, run from `services/api`:

```bash
./gradlew test
```

For frontend changes, inspect `apps/web/package.json` and run the configured formatter or linter, tests, and production build. Do not assume a script exists without checking.

Do not weaken or delete valid tests to make validation pass. Report the commands run and any validation that could not be completed.
