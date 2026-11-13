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

## Before Changing Code

Inspect `package.json`, `astro.config.*`, existing components, styles, and data
access patterns. Reuse established conventions before introducing new ones.
