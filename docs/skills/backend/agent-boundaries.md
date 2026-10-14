# Backend Agent Boundaries

The backend is educational and primarily implemented by the repository owner. Agents assist the owner while strictly preserving code ownership.

## Writing Boundaries

- Agents must NEVER write, modify, or delete files under `services/app-api` autonomously.
- Agents may only write to `services/app-api` when explicitly instructed by the repository owner.
- When explicitly requested to write or refactor backend code, agents must strictly adhere to [styleguide.md](styleguide.md) and provide didactic explanations of idioms, Go runtime mechanics, and architecture decisions.
- Every modification or proposed code must be presented and reviewed before altering any file.
- Agents may inspect and explain code, review implementations, identify bugs, suggest refactorings, and write tests when requested.

## Domain Entities & Agent Workflow

When creating or modifying domain entities under `internal/domain/entities`:

- **Follow Architecture Specifications**: All entity code must strictly adhere to the technical patterns, constructors, defensive initialization, and validation standards documented in [architecture.md](architecture.md).
- **No Autonomous Validation Assumptions**: The agent must **NEVER** autonomously decide or assume which validations belong to an entity.
- **Proactive Suggestion from Migrations**: The agent must cross-check database migrations (`NOT NULL`, types, lengths, constraints) and proactively **suggest** recommended validations.
- **Explicit Alignment**: The agent must explicitly **ask and confirm with the repository owner** which fields and rules should be validated before implementing or altering any entity.
- **TDD First**: Development must strictly follow Test-Driven Development (TDD) as defined in [testing.md](testing.md). The tests defining the expected behaviors must be written before final production code.

## Migration & Database Standards

- When reviewing or proposing PostgreSQL migrations with array columns (e.g., `VARCHAR(255)[]`, `TEXT[]`), always check and recommend `DEFAULT '{}'`.
- **Rationale for Go & Frontend integration:**
  - Prevents `NULL` values in the database, ensuring database drivers (`pgx`) scan directly into initialized empty slices (`[]string{}`) instead of `nil`.
  - Guarantees predictable JSON serialization (`"field": []` instead of `"field": null`), protecting consumer frontends (such as Astro/TypeScript) from runtime exceptions during array operations (e.g. `.map()`).
  - Eliminates the need for defensive `COALESCE(field, '{}')` in SQL array operators and queries.

## Code Style

- Adhere strictly to the Go guidelines, formatting, naming, and error handling patterns defined in [styleguide.md](styleguide.md).
- **No Abbreviations**: Always use full, descriptive names for variables, struct fields, parameters, and injected dependencies (e.g., `aboutRepository` instead of `aboutRepo`, `workExperience` instead of `workExp`).
- **Named Imports**: In production code, avoid renaming imported packages (aliases) unless strictly necessary to resolve package name collisions. Named imports are allowed in test files.
- **Compile-Time Checks**: When implementing interfaces, declare compile-time checks (`var _ Interface = (*Struct)(nil)`) and provide a brief didactic review of the mechanic to the owner.
- Do not add unnecessary comments to implementations. Write clean, self-documenting code.
- Test code must contain no comments. Use distinct blank-line blocks and descriptive test scenario names as defined in [testing.md](testing.md).
