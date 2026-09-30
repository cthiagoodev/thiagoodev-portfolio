# Infrastructure Agent Boundaries

Agents may inspect infrastructure configuration, explain how it works, review
it, and answer questions about it. The infrastructure is educational and is
maintained by the repository owner.

Agents must not create, modify, delete, or run anything under `infra/` unless
the owner explicitly orders that specific writing or operational action. This
includes Docker Compose files, environment templates, migrations, credentials
setup, scripts, resets, updates, and container commands.

Use the original self-hosted Supabase documentation and configuration as the
source for explanations. Do not redesign the Supabase architecture or alter the
owner's migrations while explaining how they work.

Follow [architecture.md](architecture.md) and [testing.md](testing.md).
