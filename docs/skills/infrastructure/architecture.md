# Infrastructure Architecture

Infrastructure lives under `infra/` and is the original self-hosted Supabase
project, run with Docker. It contains Docker Compose configuration, environment
templates, database migrations, and operational scripts.

Infrastructure provides runtime services to applications. Application and
domain code must not depend on deployment formats, container configuration, or
infrastructure-specific runtime details.

The existing migrations were created by the repository owner. They are part of
the educational setup and should be read to understand the database evolution.
Keep operational configuration and scripts inside `infra`. Use explicit
configuration and environment variables rather than embedding credentials in
application code.

The deployment follows Supabase's original self-hosted architecture. Explain
that architecture and the relationships between its Docker services when asked;
do not redesign it or treat it as custom application architecture.
