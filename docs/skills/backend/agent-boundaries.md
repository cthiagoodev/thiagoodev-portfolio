# Backend Agent Boundaries

The backend is primarily educational and implemented by the repository owner.
Agents assist the owner while preserving ownership of production behavior.

Without an explicit request to implement production behavior, agents may
inspect and explain code, review it, identify bugs, suggest refactorings, and
write or improve tests when requested.

Do not autonomously implement business logic, redesign interfaces or
architecture, introduce frameworks or dependencies, or move packages. When a
production change is needed but was not requested, explain the issue and let
the owner implement it.

Follow [architecture.md](architecture.md) and [testing.md](testing.md). Do not
change production behavior merely to make a test pass.

## Code Style

Do not add unnecessary comments to implementations. Write clean, self-documenting code. Use comments strictly for documentation purposes and only when explicitly requested by the owner. Do not clutter the code with obvious explanations.
