# Frontend Architecture

Frontend applications live under `apps/` and follow the same architecture
rules regardless of their framework.

## Responsibilities

- Maintain the existing visual identity.
- Prefer incremental changes over rewrites.
- Keep backend logic and synchronization behavior out of the frontend.
- Do not change backend contracts without explicit approval.
- Consume application data through its defined contracts.
- Do not duplicate backend synchronization or integration logic.

## Data Access (`services/api`)

`apps/web` consumes portfolio data exclusively from the Spring Boot REST API (`services/api`), without Supabase or direct database access:

- **Configuration**: `getApiUrl()` in `src/config/env.ts` reads `API_URL` on the server (`import.meta.env.SSR`).
  - Development: `http://localhost:8080/api`
  - Production: internal Docker network URL to `services/api`
- **HTTP Client**: `fetchApi<T>(path)` in `src/api/client.ts` performs server-side `fetch` requests against `API_URL`.
- **Endpoints consumed**: `/about`, `/social-contacts`, `/work-experiences`, `/educations`, `/skills`, `/projects`.

## Before Changing Code

Inspect `package.json`, `astro.config.*`, existing components, styles, and data
access patterns. Reuse established conventions before introducing new ones.
