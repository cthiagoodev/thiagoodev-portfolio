# Architecture

This document describes the high-level architecture of the portfolio repository.

It serves as the architectural reference for developers and coding agents working
in this repository.

Implementation-specific conventions are documented separately under `docs/`.

---

## 1. Overview

This repository is a monorepo containing the portfolio web application, backend
services, and infrastructure required to run the system.

At a high level, the portfolio follows this data flow:

```text
GitHub
   │
   │ repositories
   ▼
projects_sync
   │
   ├── PostgreSQL
   │
   └── Supabase
          │
          ▼
      Astro Web

The frontend does not use GitHub as its runtime data source.

External project information is collected, normalized, and persisted by the
backend before being consumed by the web application.

This keeps external integrations and synchronization behavior independent from
the frontend implementation.

2. Repository Structure

The repository is organized as:

portfolio/
├── apps/
│   └── web/
│
├── services/
│   └── projects_sync/
│
├── infra/
│   └── supabase/
│
├── docs/
│
└── ARCHITECTURE.md

Each top-level directory represents a different architectural responsibility.

apps

Contains user-facing applications.

Currently:

apps/web

web is the portfolio website implemented with Astro.

Applications consume portfolio data but must not contain backend synchronization
or integration logic.

services

Contains independently executable backend services.

Currently:

services/projects_sync

projects_sync is a Go service responsible for synchronizing project information
from external sources into the portfolio data stores.

Services own application behavior and integrations required to produce or
maintain portfolio data.

infra

Contains infrastructure configuration required to run the system.

Currently this includes the self-hosted Supabase environment and its supporting
configuration.

Infrastructure configuration must remain separate from application and domain
logic.

docs

Contains implementation guidelines and engineering documentation.

The high-level architecture belongs in ARCHITECTURE.md.

Technology-specific conventions belong under docs/.

3. Architectural Principles
3.1 Backend Owns External Synchronization

External synchronization belongs to backend services.

The frontend must not call GitHub to discover or synchronize portfolio projects.

The expected flow is:

GitHub
   ↓
projects_sync
   ↓
Portfolio Data
   ↓
Web

This prevents the frontend implementation from defining how portfolio data is
collected.

3.2 Frontend Is Replaceable

The frontend is considered a replaceable consumer of portfolio data.

Changing Astro, the UI architecture, or the visual implementation must not
require redesigning backend services or synchronization behavior.

Conceptually:

                 ┌───────────────┐
                 │ Portfolio Data│
                 └───────┬───────┘
                         │
                         ▼
                    Astro Web

Astro may eventually be replaced by another frontend technology without changing
how backend synchronization works.

3.3 Services Are Independently Executable

Backend services should be capable of being:

built independently;
tested independently;
configured independently;
executed independently;
deployed independently.

A service must not depend on another service's internal packages.

Shared behavior should only be extracted when a concrete shared requirement
exists.

3.4 Infrastructure Is an Implementation Detail

Technologies used to communicate with external systems are implementation
details.

Examples include:

PostgreSQL;
pgx;
Supabase;
HTTP clients;
GitHub APIs;
cron schedulers.

These technologies must not define the core domain model.

Infrastructure-specific types should remain inside infrastructure or entrypoint
boundaries whenever possible.

3.5 Dependencies Point Inward

Backend code follows an inward dependency direction.

Conceptually:

Presentation ───────┐
                    │
                    ▼
               Application
                    │
                    ▼
                  Domain
                    ▲
                    │
Infrastructure ─────┘

The important rule is:

Domain does not depend on Infrastructure.

The responsibilities are:

Presentation
    │
    │ initiates operations
    ▼
Application
    │
    │ coordinates behavior
    ▼
Domain

while:

Infrastructure
    │
    │ implements external boundaries
    ▼
Domain/Application contracts

The architecture is based on dependency direction and responsibility rather than
strict adherence to a framework or architecture template.

4. projects_sync

projects_sync is currently the primary backend service.

Its responsibility is to synchronize portfolio project information.

The service is implemented in Go.

Its internal organization follows:

services/projects_sync/
└── internal/
    ├── application/
    │   └── usecases/
    │
    ├── domain/
    │   ├── entities/
    │   ├── repositories/
    │   └── usecases/
    │
    ├── infrastructure/
    │   ├── github/
    │   ├── repositories/
    │   └── supabase/
    │
    ├── presentation/
    │   └── scheduler/
    │
    └── test/
        └── mocks/

Each package has a distinct responsibility.

Domain

Contains core entities and contracts.

The domain must not know how GitHub, PostgreSQL, Supabase, HTTP, or scheduling
are implemented.

Application

Contains use cases that coordinate application behavior.

Application code determines what the system does, but should avoid knowing
low-level details about how external systems perform those operations.

Infrastructure

Contains implementations for communication with external systems.

Current examples include:

GitHub API
PostgreSQL
Supabase
Presentation

Contains entrypoints that initiate application behavior.

The current scheduler belongs here because it determines when synchronization
is requested.

Presentation should delegate actual application behavior to use cases.

5. Project Synchronization

The current synchronization flow is:

Scheduler
    │
    ▼
SyncProjectsUseCase
    │
    ├───────────────┐
    │               │
    ▼               │
GitHubService       │
    │               │
    ▼               │
 GitHub             │
                    │
                    ▼
            ProjectsMapper
                    │
                    ▼
          ProjectsRepository
                    │
                    ▼
               PostgreSQL
                    │
                    ▼
             Get Projects
                    │
                    ▼
            SupabaseService
                    │
                    ▼
                Supabase

The responsibilities are intentionally separated.

Scheduler

Answers:

When should synchronization execute?

It does not define synchronization business behavior.

SyncProjectsUseCase

Answers:

What should happen when project synchronization is requested?

It coordinates the synchronization process.

GitHubService

Answers:

How are repositories retrieved from GitHub?

GitHub-specific HTTP behavior belongs to infrastructure.

ProjectsMapper

Transforms external GitHub representations into the internal project
representation used by the application.

ProjectsRepository

Defines persistence operations required by the application.

The concrete PostgreSQL implementation belongs to infrastructure.

SupabaseService

Handles synchronization with the Supabase-backed portfolio data source.

Supabase-specific behavior remains an infrastructure concern.

6. Persistence and Transactions

Replacing persisted project data is considered one logical operation.

Removing the previous projects and inserting the new project collection must
therefore be atomic.

Conceptually:

BEGIN
   │
   ├── DELETE existing projects
   │
   ├── COPY new projects
   │
COMMIT

If any operation fails:

BEGIN
   │
   ├── DELETE existing projects
   │
   ├── COPY new projects
   │       │
   │       └── ERROR
   │
ROLLBACK

The system must not leave the project store empty or partially replaced because
an insertion failed after deletion.

Transaction mechanics belong to persistence infrastructure.

PostgreSQL-specific transaction types must not leak into the domain or
application contracts unless there is a concrete architectural reason.

7. Context Propagation

Operations involving I/O or potentially long-running work use
context.Context.

The context originates at an application entrypoint and is propagated through
the operation.

Conceptually:

Entrypoint
    │
    │ context.Context
    ▼
Use Case
    │
    ├───────────────┐
    ▼               ▼
GitHub           Repository
    │               │
    ▼               ▼
HTTP            PostgreSQL

This allows cancellation and deadlines to propagate to external operations.

Infrastructure implementations must not replace an incoming context with
context.Background() simply to avoid propagating the caller's context.

Context represents the lifetime of an operation.

It should not be stored as state inside long-lived application or infrastructure
objects.

8. Scheduling

Project synchronization is scheduled work.

Scheduling answers:

When should the use case execute?

Synchronization answers:

What should happen when it executes?

These concerns remain separate.

The current architecture is:

Cron Scheduler
      │
      ▼
SyncProjectsUseCase

The scheduler is an application entrypoint and therefore belongs to the
presentation boundary.

The scheduler must not:

communicate directly with GitHub;
execute SQL;
communicate directly with Supabase;
perform project mapping;
contain synchronization business logic.

Its responsibility is to trigger the appropriate application use case.

The scheduling mechanism must remain replaceable.

For example, the current scheduler could eventually be replaced by:

Linux cron
Kubernetes CronJob
Cloud Scheduler
Container Job
CI/CD scheduled job

without changing the synchronization use case.

9. Data Ownership

Backend services are responsible for producing and maintaining portfolio data.

The frontend consumes that data.

This creates the boundary:

External Sources
       │
       ▼
Backend Services
       │
       ▼
Portfolio Data
       │
       ▼
Frontend

The frontend should not reproduce backend synchronization logic.

Likewise, backend services should not depend on frontend implementation details.

10. External Integrations

External systems must be accessed through explicit boundaries.

Current external integrations include:

GitHub
PostgreSQL
Supabase

The architecture should make it possible to replace an implementation without
changing core application behavior when practical.

For example:

GitHubService
      │
      ▼
GitHub HTTP implementation

or:

ProjectsRepository
       │
       ▼
PostgreSQL implementation

Interfaces should represent meaningful behavioral boundaries.

Interfaces should not be created solely to abstract every concrete type or to
make every implementation mockable.

11. Observability

Operational behavior should eventually be observable through:

Logs
Metrics
Traces

Observability must remain a cross-cutting concern rather than becoming business
logic.

Instrumentation should not alter domain behavior.

The project may introduce dedicated observability tooling as the deployment
architecture evolves.

12. Security

Secrets and privileged credentials must remain outside source code.

Frontend/browser code must never receive backend credentials.

Infrastructure configuration and environment variables are responsible for
providing runtime secrets.

Security-specific infrastructure will evolve separately from application
business logic.

13. Future Evolution

The architecture should evolve from concrete requirements rather than
hypothetical future complexity.

Potential future additions include:

Messaging
Additional backend services
Administrative interfaces
Authentication
Observability infrastructure
Additional deployment targets

These should only be introduced when they solve an actual requirement.

In particular, current services must not be designed around a future messaging
system before an independent asynchronous communication requirement exists.

Prefer:

current requirement
       ↓
simplest appropriate architecture
       ↓
measure / learn
       ↓
evolve architecture

over:

possible future requirement
       ↓
premature infrastructure
       ↓
unnecessary complexity
14. Architectural Changes

Changes that modify system boundaries should be treated differently from normal
implementation changes.

Examples include:

introducing a new service;
changing service responsibilities;
introducing messaging between components;
changing persistence ownership;
changing dependency direction;
introducing a new runtime entrypoint;
moving business behavior between frontend and backend.

These changes should be documented before or alongside their implementation.

Small implementation details do not require architectural documentation.

15. Related Documentation

Implementation-specific guidance is maintained separately:

docs/backend.md
docs/frontend.md
docs/testing.md

ARCHITECTURE.md describes what the system architecture is and why its major
boundaries exist.

The documents under docs/ describe how work should be performed inside those
boundaries.

When documentation and implementation disagree, the discrepancy should be
investigated rather than silently assuming either one is correct.