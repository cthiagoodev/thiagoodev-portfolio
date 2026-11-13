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

Every test scenario must follow the **AAA** structure separated strictly by distinct blank-line blocks, ensuring clarity, maintainability, and fast root-cause identification:

- **No Comment Rule**: Test code (`*Test.java`) must contain **NO COMMENTS**. Do NOT add `// Arrange`, `// Act`, `// Assert`, or `// Given / When / Then` comments.
- **Blank-line Blocks**:
  1. **Arrange**: Initialize objects, test fixtures, initial state, and Mockito mock setup (`when(...).thenReturn(...)`).
  2. **Act**: Execute the single constructor, method, use case, or HTTP request under test.
  3. **Assert**: Verify the expected outcome (`assertNotNull`, `assertEquals`, `assertThrows`, `verify(...)`, `jsonPath(...)`).

## Test Naming Principles & Conventions

Test names must be self-explanatory and act as living documentation of system behavior.

### Core Naming Principles
1. **Describe Expected Behavior**: Clearly state what the system does under a specific condition (`shouldReturnOk_whenAboutExists`).
2. **Avoid Implementation Details**: Focus on observable behavior and use cases rather than internal variables, private methods, or internal logic.
3. **No Abbreviations**: Use full, descriptive words even if the method name is longer (`workExperience` instead of `workExp`, `repository` instead of `repo`).
4. **Avoid Generic or Vague Names**: Never use `test1()`, `testFunction()`, `validate()`, or `checkResult()`.
5. **Balance Clarity and Conciseness**: Avoid overly long names with multiple chained conditions. Split complex scenarios into separate, focused tests covering one behavior at a time.

### Approved Java Naming Patterns

Choose one of the following consistent conventions per test suite:

- **`should<ExpectedBehavior>_when<Condition>`** (Behavior-driven):
  - `shouldCreateAbout_whenAllFieldsAreValid()`
  - `shouldThrowDomainValidationException_whenNameIsBlank()`
  - `shouldReturnNotFound_whenAboutDoesNotExist()`
- **`given<InitialState>_when<Action>_then<ExpectedResult>`** (BDD Given-When-Then):
  - `givenValidFields_whenCreatingAbout_thenReturnInitializedEntity()`
  - `givenEmptyRepository_whenExecutingGetAbout_thenThrowResourceNotFoundException()`
- **`<methodName>_<stateOrCondition>_should<ExpectedBehavior>`** (Method-focused):
  - `execute_withExistingAbout_shouldReturnAbout()`
  - `addSkill_withDuplicateLabel_shouldIgnoreDuplicate()`

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
