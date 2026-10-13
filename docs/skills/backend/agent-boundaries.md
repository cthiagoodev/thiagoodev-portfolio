# Backend Agent Boundaries

The backend is educational and primarily implemented by the repository owner. Agents assist the owner while strictly preserving code ownership.

## Writing Boundaries

- Agents must NEVER write, modify, or delete files under `services/app-api` autonomously.
- Agents may only write to `services/app-api` when explicitly instructed by the repository owner.
- Every modification or proposed code must be presented and reviewed before altering any file.
- Agents may inspect and explain code, review implementations, identify bugs, suggest refactorings, and write tests when requested.

## Migration & Database Standards

- When reviewing or proposing PostgreSQL migrations with array columns (e.g., `VARCHAR(255)[]`, `TEXT[]`), always check and recommend `DEFAULT '{}'`.
- **Rationale for Go & Frontend integration:**
  - Prevents `NULL` values in the database, ensuring database drivers (`pgx`) scan directly into initialized empty slices (`[]string{}`) instead of `nil`.
  - Guarantees predictable JSON serialization (`"field": []` instead of `"field": null`), protecting consumer frontends (such as Astro/TypeScript) from runtime exceptions during array operations (e.g. `.map()`).
  - Eliminates the need for defensive `COALESCE(field, '{}')` in SQL array operators and queries.

## Code Style

- Do not add unnecessary comments to implementations. Write clean, self-documenting code.
- Test code must contain no comments. Use distinct blank-line blocks and descriptive test scenario names.
