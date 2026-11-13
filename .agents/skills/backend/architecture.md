# Backend Architecture

The backend lives under `services/api` and follows a pragmatic, layered Clean Architecture in **Java 25** and **Spring Boot 4.1+**, designed to protect business rules and enforce high unit testability while fully leveraging Spring Boot's native capabilities (Dependency Injection, WebMVC, Data JPA, Flyway, and Global Exception Handling).

## Overview

`services/api` is a modular Spring Boot 4 application built with Gradle (Kotlin DSL):
- **REST API (`presentation/controllers`)**: Powered by `spring-boot-starter-webmvc`, serving public and administrative endpoints consumed by the Astro frontend.
- **Persistence & Migrations (`domain/repositories` & `db/migration`)**: Powered by `spring-boot-starter-data-jpa` (Hibernate/JPA), HikariCP connection pooling, and `spring-boot-starter-flyway` with PostgreSQL 18.
- **Background Synchronization (`infrastructure/github`)**: Scheduled synchronization jobs (`@Scheduled`) and HTTP integrations using Spring's modern `RestClient`.

## Directory & Package Structure

```text
services/api/
├── build.gradle.kts
├── settings.gradle.kts
└── src/
    ├── main/
    │   ├── java/br/com/thiagoodev/api/
    │   │   ├── ApiApplication.java          # Entrypoint (@SpringBootApplication)
    │   │   ├── domain/                      # Core business entities, contracts, utilities, and domain exceptions
    │   │   │   ├── entities/                # Domain/JPA Entities (About, SocialContact, Project, Skill, Education, WorkExperience)
    │   │   │   ├── repositories/            # Repository interfaces (extending JpaRepository)
    │   │   │   ├── usecases/                # UseCase contracts (interfaces consumed by controllers)
    │   │   │   ├── exceptions/              # Domain exceptions (DomainValidationException, ResourceNotFoundException)
    │   │   │   └── utils/                   # Pure domain utilities (Sanitizer.sanitize for string trimming/null safety)
    │   │   ├── application/                 # Business orchestration and use case implementations
    │   │   │   ├── usecases/                # UseCase implementations (@Service: GetAboutUseCaseImpl, etc.)
    │   │   │   └── dto/                     # Immutable Java Records for Input/Output mapping when needed
    │   │   ├── infrastructure/              # External integrations and framework configurations
    │   │   │   ├── github/                  # GitHub REST API client (Spring RestClient)
    │   │   │   └── config/                  # Spring @Configuration classes (CORS, RestClient, Scheduler)
    │   │   └── presentation/                # HTTP transport layer
    │   │       ├── controllers/             # @RestController classes (AboutController, SocialContactController, etc.)
    │   │       └── exception/               # @RestControllerAdvice centralized error mapping (GlobalExceptionHandler)
    │   └── resources/
    │       ├── application.properties       # Datasource, JPA, Flyway, and Server configuration
    │       └── db/migration/                # Versioned SQL migrations managed by Flyway (V1__*.sql)
    └── test/
        └── java/br/com/thiagoodev/api/      # Mirror package structure for unit and slice tests
```

## What Spring Boot 4 Handles Natively (Do Not Reinvent)

Unlike manual setups in minimalist languages, the architecture relies on Spring Boot's native engines rather than custom boilerplate:
- **No manual `modules/` or wiring layer**: Spring's IoC container automatically wires `@Repository` → `@Service` (UseCases) → `@RestController` via constructor injection.
- **No manual `router/` layer**: Routes are declared declaratively on controllers via `@RestController`, `@RequestMapping`, and `@GetMapping` / `@PostMapping`.
- **No manual database pool management**: HikariCP manages connection lifecycle, pooling, and health checks automatically via `spring.datasource.*`.
- **No manual JSON response helpers**: Spring WebMVC + Jackson serialize entities, records, and `ResponseEntity<T>` automatically.
- **No manual per-endpoint error switches**: `@RestControllerAdvice` intercepts exceptions globally and converts them into standardized HTTP responses (`ProblemDetail` RFC 9457 or structured JSON).

## Dependency Direction

Dependencies point strictly inward toward the `domain` package:

```text
Presentation (Controllers) → Domain UseCase Contracts ← Application (UseCase Implementations)
Application (UseCase Implementations) → Domain (Entities, Repository Contracts, Exceptions)
Infrastructure → Domain Contracts
```

### 1. Camada de Domínio (`domain`)
- O coração do sistema: concentra as entidades de negócio, invariantes de validação, contratos de repositórios, contratos de casos de uso e exceções semânticas.
- **Abordagem Pragmática Spring (Domain + JPA Entities):**
  - As classes em `domain.entities` representam o modelo de domínio rico e incluem o mapeamento Jakarta Persistence (`@Entity`, `@Table`, `@Column`, etc.), evitando duplicação desnecessária de classes (`DomainEntity` vs `JpaEntity`) e mapeadores repetitivos.
  - **Construtores e Invariantes de Domínio:**
    - Além do construtor protegido sem argumentos exigido pelo JPA (`protected Entity() {}`), toda entidade deve possuir um construtor de domínio completo (ou factory method) que higienize entradas via `Sanitizer.sanitize(...)` (`domain.utils.Sanitizer`), inicialize coleções defensivamente (`new ArrayList<>()`, nunca `null`), defina timestamps padrão (`LocalDateTime.now()`) quando ausentes e invoque `validate()`.
    - O método `validate()` verifica todas as invariantes de negócio e lança `DomainValidationException` em caso de violação.
- **Relacionamentos N:N (Tabelas Associativas):**
  - Tabelas relacionais auxiliares como `projects_skills` não são classes de entidade separadas em `domain.entities`.
  - O relacionamento é modelado diretamente na entidade agregadora (`Project` possui `List<Skill> skills` mapeado via `@ManyToMany` e `@JoinTable(name = "projects_skills", ...)`), expondo métodos de domínio como `addSkill(Skill skill)` e `removeSkill(UUID skillUuid)`.
- **Contratos no Domínio (`domain.repositories` e `domain.usecases`):**
  - `domain.repositories`: Interfaces estendendo `JpaRepository<Entity, UUID>` (ex.: `AboutRepository`, `SocialContactRepository`).
  - `domain.usecases`: Interfaces claras para cada caso de uso (ex.: `GetAboutUseCase`, `GetAllSocialContactsUseCase`).
- **Exceções de Domínio (`domain.exceptions`):**
  - `DomainValidationException`: Lançada quando invariantes de uma entidade falham (mapeada para HTTP 400 Bad Request).
  - `ResourceNotFoundException`: Lançada quando um recurso esperado não é encontrado no banco (mapeada para HTTP 404 Not Found).
- **Testes de Domínio:** Testes unitários JUnit 5 puros, ultrarrápidos, **sem subir contexto do Spring** e **sem mocks**.

### 2. Camada de Aplicação (`application`)
- Orquestra o fluxo de dados entre os repositórios e as regras de domínio.
- **Implementações de Casos de Uso (`application.usecases`):**
  - Classes anotadas com `@Service` que implementam sua respectiva interface de `domain.usecases` (ex.: `GetAboutUseCaseImpl implements GetAboutUseCase`).
  - Recebem dependências exclusivamente via injeção por construtor (`final` fields + construtor explícito ou `@RequiredArgsConstructor`).
  - **Sem Abreviações:** Variáveis, atributos e parâmetros utilizam sempre nomes completos (`aboutRepository`, `socialContactRepository`, `workExperience`).
- **DTOs (`application.dto`):**
  - Quando necessário desacoplar o contrato de entrada/saída da entidade, utilizam-se `record`s imutáveis do Java 25.
- **Testes de Casos de Uso:** Testados isoladamente com JUnit 5 + Mockito (`@ExtendWith(MockitoExtension.class)`), sem precisar subir o contexto completo do Spring (`@SpringBootTest`).

### 3. Camada de Infraestrutura (`infrastructure`)
- Contém clientes HTTP externos (ex.: integração com a API do GitHub usando `RestClient`) e classes `@Configuration` do Spring.
- Os repositórios de banco de dados são implementados automaticamente em tempo de execução pelo Spring Data JPA a partir das interfaces em `domain.repositories`.

### 4. Camada de Apresentação (`presentation`)
- **Controllers (`presentation.controllers`):**
  - Classes anotadas com `@RestController` e `@RequestMapping`.
  - **Regra de Ouro:** Controllers dependem **apenas das interfaces de Casos de Uso** (`domain.usecases`). Nunca injetam `Repository` diretamente.
  - Mantêm código limpo e enxuto: delegam a execução para `useCase.execute(...)` e retornam `ResponseEntity.ok(result)` (sem blocos `try/catch` manuais).
- **Tratamento Global de Erros (`presentation.exception.GlobalExceptionHandler`):**
  - Classe anotada com `@RestControllerAdvice` que captura exceções lançadas em qualquer camada e retorna respostas HTTP padronizadas:
    - `ResourceNotFoundException` → `404 Not Found`
    - `DomainValidationException` / `MethodArgumentNotValidException` → `400 Bad Request`
    - `Exception` (erros inesperados) → `500 Internal Server Error` com mensagem genérica segura.
- **Testes de Apresentação:** Testados com `@WebMvcTest(Controller.class)` + `MockMvc` (ou testes unitários diretos com `MockMvcBuilders.standaloneSetup`), mockando apenas o caso de uso com `@MockitoBean`.

## Flyway & PostgreSQL 18+ Standards

O projeto utiliza **PostgreSQL 18 (`postgres:18.6`)** e **Flyway** para versionamento do schema.

- **Localização das Migrations:** `src/main/resources/db/migration/` seguindo a nomenclatura padrão do Flyway:
  - `V1__create_about_table.sql`
  - `V2__create_social_contacts_table.sql`
  - `V3__create_work_experience_table.sql`
  - `V4__create_education_table.sql`
  - `V5__create_skills_table.sql`
  - `V6__create_projects_table.sql`
  - `V7__create_projects_skills_table.sql`
- **Validação de Schema JPA:** `spring.jpa.hibernate.ddl-auto=validate` — o Hibernate nunca altera o banco sozinho; apenas valida se as entidades JPA estão 100% fiéis às migrations do Flyway.
- **PostgreSQL 18 Volume Mount Path:** No `docker-compose.yaml`, o volume DEVE ser montado em `/var/lib/postgresql`, nunca em `/var/lib/postgresql/data`:
  ```yaml
  volumes:
    - postgres_data:/var/lib/postgresql
  ```
