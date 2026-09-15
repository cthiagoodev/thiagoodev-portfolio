# Frontend Agent Guide

Codex owns implementation work in this directory when requested.

## Goals

Maintain the existing visual identity.

Prefer incremental changes over rewrites.

Do not change backend contracts without explicit approval.

Do not introduce backend logic into the frontend.

## Data

The frontend consumes portfolio data from Supabase.

Do not query GitHub directly.

Do not duplicate backend synchronization logic.

## Before changing code

Inspect:

- package.json
- astro.config.*
- existing components
- existing styles
- existing data access patterns

Reuse existing conventions before introducing new ones.

## Validation

Before declaring work complete:

1. install dependencies only when necessary;
2. run formatter/linter configured by the project;
3. run tests;
4. run the production build;
5. inspect the final git diff.

Never suppress an error merely to make validation pass.