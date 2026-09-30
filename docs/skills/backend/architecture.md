# Backend Architecture

Backend services live under `services/` and follow the same architectural
rules regardless of their purpose or runtime.

## Dependency Direction

Dependencies point inward:

```text
Presentation → Application → Domain
Infrastructure → Application/Domain contracts
```

The domain contains core entities and behavioral contracts. It must not depend
on infrastructure technologies such as database drivers, HTTP clients, external
APIs, schedulers, or framework types.

Application code contains use cases and coordinates behavior. It does not
contain SQL, HTTP implementation details, scheduling configuration, or runtime
configuration.

Infrastructure implements external boundaries such as persistence, HTTP
clients, and third-party APIs. Presentation contains entrypoints and delegates
application behavior to use cases.

## Design Rules

- Use explicit constructor injection for required dependencies.
- Create small interfaces only for meaningful behavioral boundaries.
- Do not add layers, packages, frameworks, or dependencies without a concrete
  requirement.
- Keep I/O operations context-aware and propagate the incoming
  `context.Context`.
- Do not store contexts in long-lived structs or replace an incoming context
  with `context.Background()`.
- Keep transactions inside infrastructure when they implement one persistence
  operation. Do not expose database transaction types through domain contracts.
- Return errors to the layer that can decide their meaning; add context only
  when it improves diagnosis.

## Service Independence

Each service must be able to build, test, configure, run, and deploy
independently. A service must not import another service's internal packages.

Architecture changes that alter ownership, dependency direction, contracts,
or runtime entrypoints require explicit instruction and documentation.
