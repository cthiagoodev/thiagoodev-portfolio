---
name: backend
description: >-
  Architecture, Java 25 and Spring Boot 4 styleguide, strict TDD testing rules,
  and agent boundaries for the backend service under services/api. Activate this
  skill whenever reading, planning, testing, or modifying backend code.
---

# Skill: Backend Engineering (Java 25 & Spring Boot 4)

This skill defines the architecture, coding standards, testing rules, and agent boundaries for the backend service (`services/api`).

## Reference Guides

Read the relevant reference files in this directory before acting on backend tasks:

- **[architecture.md](./architecture.md)**: Pragmatic layered Clean Architecture (`domain`, `application`, `infrastructure`, `presentation`), JPA/Hibernate mapping, Flyway migrations, and PostgreSQL 18 standards.
- **[styleguide.md](./styleguide.md)**: Java 25 & Spring Boot 4 conventions, constructor injection, Lombok restrictions on JPA entities, `Sanitizer.sanitize`, and exception handling.
- **[testing.md](./testing.md)**: Strict TDD (Red-Green-Refactor), AAA blank-line blocks with zero comments, JUnit 5, Mockito (`@ExtendWith(MockitoExtension.class)`), and `@WebMvcTest`.
- **[agent-boundaries.md](./agent-boundaries.md)**: Pair-programming rules, entity validation alignment, and code ownership constraints.
