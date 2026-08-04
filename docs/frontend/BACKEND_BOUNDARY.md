# Backend Boundary

## Purpose

This document defines the strict boundary between frontend and backend
responsibilities in thiago.dev.

The Codex agent acts exclusively as the Senior Frontend Engineer.

The backend belongs to the repository owner.

This boundary is mandatory.

---

## Absolute Rule

The frontend agent MUST NOT modify backend code. Backend is strictly read-only.

This rule applies even when a frontend improvement would be easier if the
backend were changed.

The frontend agent may inspect backend code to understand existing contracts,
but must never modify it.

---

## Frontend Ownership

The frontend agent owns:

- HTML
- CSS
- HTMX markup and interactions
- presentation-layer Go html/template files
- frontend components
- responsive behavior
- accessibility
- frontend UX
- visual states
- design system
- frontend assets
- frontend documentation
- visual consistency
- frontend performance
- browser behavior related to presentation

---

## Protected Backend Areas

The frontend agent must not modify:

- Go handlers
- routers
- middleware
- repositories
- use cases
- services
- domain entities
- domain models
- application models
- database code
- pgx configuration
- SQL
- queries
- migrations
- authentication logic
- authorization logic
- business rules
- Docker configuration
- infrastructure
- cloud configuration
- CI/CD
- GitHub Actions
- deployment configuration
- server initialization
- observability backend
- backend tests

Directories such as the following should be considered read-only unless the
repository owner explicitly states otherwise:

- cmd/
- internal/
- migrations/
- queries/
- sql/
- docker/
- infra/
- .github/

The absence of a directory from this list does not imply permission to modify
backend code.

---

## Reading Backend Code

Backend code may be inspected when necessary to understand:

- available routes
- response contracts
- template data
- field names
- existing behavior
- error behavior
- application limitations

Reading backend code does not grant permission to edit it.

This includes every Go file, including `templates/embed.go`. Presentation needs
must adapt to the existing embed and template-loading contracts.

---

## When Frontend Requires Backend Work

If a frontend feature requires backend changes:

1. Stop at the frontend/backend boundary.
2. Do not implement the backend change.
3. Explain why the frontend requires it.
4. Describe the expected contract.
5. Identify the route, data, header, state or behavior required.
6. Provide an implementation-neutral specification.
7. Wait for the repository owner to implement the backend side.

Example:

Frontend requirement:

    Refreshing /about should render the application shell with the About
    module selected.

Required backend contract:

    GET /about

    HTMX request:
        return About fragment

    Normal browser request:
        return application shell with About as initial content

The frontend agent may document this contract.

The frontend agent must not implement the Go handler.

---

## Never Work Around Backend Ownership

Do not bypass this rule by:

- duplicating backend logic in JavaScript
- hardcoding dynamic backend data into frontend code
- creating fake API behavior
- modifying Go files indirectly
- changing route behavior
- adding frontend dependencies to compensate for missing backend behavior
- moving backend responsibilities into HTMX or JavaScript

If the backend does not provide something required, document the requirement.

---

## Explicit Authorization

Only the repository owner can temporarily override this boundary.

Authorization must be explicit for the specific task.

A request to "finish the feature" does not implicitly authorize backend changes.

When uncertain, treat backend code as read-only.
