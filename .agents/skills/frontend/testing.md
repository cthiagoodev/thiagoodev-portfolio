# Frontend Testing and Validation

Follow the tooling configured in `apps/web`. Inspect `apps/web/package.json`
before choosing commands; do not assume scripts such as `test`, `lint`, or
`check` exist.

Before declaring frontend work complete:

1. install dependencies only when necessary;
2. run the configured formatter and linter;
3. run available tests;
4. run the production build;
5. inspect the final git diff.

Use the available static checks, tests, and production build commands for the
project. Never suppress an error merely to make validation pass.

## Test Readability

Do not add comments to tests. Explain the behavior through descriptive test and
subtest names, focused cases, and clear code structure.

Separate test preparation, execution, and assertions into visually distinct
blocks with blank lines. Avoid dense setup and keep each test focused on one
observable behavior.
