# Backend Testing

Test-Driven Development (TDD) and comprehensive unit testing are the core development strategies for all backend code under `services/api`. Tests define the contract, structure, and design before implementation begins.

## TDD Workflow: Red-Green-Refactor

All domain entities, use cases, controllers, and exception handlers must be implemented following the TDD cycle:

```text
1. RED      → Escreva o teste unitário primeiro, definindo o comportamento esperado. Execute `./gradlew test` e veja-o falhar.
2. GREEN    → Escreva a implementação mínima necessária para fazer o teste passar.
3. REFACTOR → Melhore a legibilidade, arquitetura e performance sem alterar o comportamento, mantendo a suíte verde.
```

## AAA Pattern (Arrange, Act, Assert)

Every test scenario must follow the **AAA** structure separated strictly by distinct blank-line blocks.

- **No Comment Rule**: Test code (`*Test.java`) must contain **NO COMMENTS**. Do NOT add `// Arrange`, `// Act`, `// Assert`, or `// Given / When / Then` comments.
- **Blank-line Blocks**:
  1. **Arrange**: Input variables, test fixtures, and Mockito mock setup (`when(...).thenReturn(...)`).
  2. **Act**: Execution of the constructor, method, use case, or HTTP request under test.
  3. **Assert**: Verifications (`assertNotNull`, `assertEquals`, `assertThrows`, `verify(...)`, etc.).

## Test Naming Convention

Tests must use descriptive, behavior-driven method names in `camelCase` or `snake_case` following the pattern:

```text
should<ExpectedBehavior>_when<Condition>
```
or
```text
<methodOrTarget>_<expectedBehavior>[_<condition>]
```

### Examples:
- `newAbout_success()`
- `newAbout_validationError_emptyName()`
- `execute_success()`
- `execute_throwsResourceNotFoundException_whenEmpty()`
- `get_returnsOkWithAboutPayload()`
- `get_returnsNotFound_whenAboutDoesNotExist()`

## Unit Testing Strategy by Layer (Fast & Isolated)

Prioritize fast, isolated unit tests that do not boot the full database or entire application context unless strictly testing a slice:

1. **Domain Entities (`domain/entities/*Test.java`)**:
   - **Pure JUnit 5 tests** (`org.junit.jupiter.api.Test`), zero Spring context, zero mocks.
   - Verify valid construction, defensive collection initialization, string trimming, default timestamps, domain methods (`addSkill`, `removeSkill`), and `DomainValidationException` rules using `assertThrows`.
2. **Application Use Cases (`application/usecases/*Test.java`)**:
   - **Pure JUnit 5 + Mockito** (`@ExtendWith(MockitoExtension.class)`), using `@Mock` for repositories and `@InjectMocks` for the UseCase implementation.
   - Verify business orchestration, repository interactions (`verify(repository, times(1))`), and exception propagation without starting Spring.
3. **Presentation Controllers & Global Exception Handler (`presentation/controllers/*Test.java`, `presentation/exception/*Test.java`)**:
   - Tested using `@WebMvcTest(TargetController.class)` with `MockMvc` and `@MockitoBean` for the UseCase interface (or standalone `MockMvcBuilders.standaloneSetup` with `GlobalExceptionHandler`).
   - Verify HTTP status codes (`200 OK`, `400 Bad Request`, `404 Not Found`, `500 Internal Server Error`), `Content-Type` headers, and JSON response bodies (`jsonPath`).
4. **Database Repositories (`domain/repositories`)**:
   - Standard Spring Data JPA interfaces do not require unit tests unless custom `@Query` or complex specifications are introduced.

## Assertions and Determinism

- Use JUnit 5 (`org.junit.jupiter.api.Assertions.*`) or AssertJ (`org.assertj.core.api.Assertions.assertThat`) consistently.
- Verify mock interactions explicitly with Mockito (`verify(...)`, `verifyNoMoreInteractions(...)`) when testing use cases and controllers.
- Keep tests 100% deterministic: no `Thread.sleep`, no external network calls, no shared mutable static state.

## Required Validation

To run the test suite from `services/api`:

```bash
./gradlew test
```

Do not weaken or delete a valid test to make the suite pass.
