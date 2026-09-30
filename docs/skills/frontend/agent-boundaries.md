# Frontend Agent Boundaries

Agents may implement and refactor frontend code under `apps/` when requested.

Keep changes within frontend responsibilities:

- preserve existing visual identity and prefer incremental changes;
- do not implement backend logic in the frontend;
- do not change backend contracts without explicit approval;
- do not query external backend sources directly or duplicate backend
  synchronization;
- follow [`architecture.md`](architecture.md) and
  [`testing.md`](testing.md).

The repository root [`AGENTS.md`](../../../AGENTS.md) is the entry point for
agent instructions.

## Code Style

Do not add unnecessary comments to implementations. Write clean, self-documenting code. Use comments strictly for documentation purposes and only when explicitly requested by the owner. Do not clutter the code with obvious explanations.
