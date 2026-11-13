# Backend Java & Spring Boot Styleguide

Este guia define as convenções de estilo, padrões idiomáticos e práticas de engenharia em **Java 25** e **Spring Boot 4** adotadas no backend (`services/api`).

---

## 1. Formatação e Idioma Moderno (Java 25)

- **Indentação**: 4 espaços por nível de indentação (padrão oficial Java).
- **Java 25 Moderno**:
  - Utilize `record` para DTOs imutáveis e objetos de transferência de dados.
  - Utilize *Pattern Matching* para `instanceof` e `switch` expressions quando apropriado.
  - Utilize `var` apenas para variáveis locais cujo tipo seja imediatamente óbvio no lado direito da atribuição; prefira tipos explícitos quando melhorar a legibilidade.
  - Utilize `List.of()`, `Set.of()` e `Map.of()` para coleções imutáveis e `new ArrayList<>()` quando a entidade exigir mutabilidade controlada.

---

## 2. Nomenclatura (Naming Conventions)

A nomenclatura em Java deve ser expressiva, sem ambiguidades e revelar a intenção do componente.

### Pacotes (Packages)
- Sempre em letras minúsculas (`br.com.thiagoodev.api.domain.entities`).
- Evite underscores nos nomes de pacotes; prefira palavras compostas juntas ou subpacotes claros (`usecases`, `repositories`, `controllers`, `exception`).

### Classes, Interfaces e Métodos
- **Classes e Records**: `PascalCase` substantivo (`About`, `SocialContact`, `AboutController`, `GlobalExceptionHandler`).
- **Interfaces de Casos de Uso**: `PascalCase` descrevendo a ação (`GetAboutUseCase`, `GetAllSocialContactsUseCase`).
- **Implementações de Casos de Uso**: Sufixo `Impl` (`GetAboutUseCaseImpl`, `GetAllSocialContactsUseCaseImpl`).
- **Controllers**: Sufixo `Controller` (`AboutController`, `SocialContactController`).
- **Métodos e Variáveis**: `camelCase` (`getAll`, `execute`, `socialContactRepository`).
- **Constantes**: `UPPER_SNAKE_CASE` (`DEFAULT_PAGE_SIZE`).

### Sem Abreviações (No Abbreviations)
- **Regra Estrita**: Não utilize abreviações para variáveis locais, atributos de classe, parâmetros de métodos ou dependências injetadas. Sempre utilize nomes completos e explícitos:
  - Correto: `aboutRepository`, `socialContactRepository`, `workExperience`, `education`, `databaseException`.
  - Incorreto: `aboutRepo`, `scRepo`, `workExp`, `exp`, `edu`, `dbEx`.

---

## 3. Injeção de Dependências e Uso do Spring

### Injeção Exclusiva via Construtor (Constructor Injection)
- **Nunca utilize `@Autowired` em atributos (field injection)**.
- Todos os componentes (`@Service`, `@RestController`, `@Component`) devem declarar suas dependências como `private final` e recebê-las via construtor (explícito ou via `@RequiredArgsConstructor` do Lombok).
- **Por quê**: Garante imutabilidade das dependências, impede estado parcialmente inicializado e permite instanciar a classe diretamente em testes unitários puros sem precisar subir o container do Spring.

```java
@Service
@RequiredArgsConstructor
public class GetAboutUseCaseImpl implements GetAboutUseCase {

    private final AboutRepository aboutRepository;

    @Override
    public About execute() {
        return aboutRepository.findFirstByOrderByCreatedAtAsc()
                .orElseThrow(() -> new ResourceNotFoundException("About information not found"));
    }
}
```

### Uso Consciente do Lombok
- **Permitido**: `@Getter`, `@RequiredArgsConstructor`, `@AllArgsConstructor`, `@NoArgsConstructor(access = AccessLevel.PROTECTED)`, `@Builder` (quando adequado).
- **Proibido em Entidades JPA (`@Entity`)**:
  - **Nunca use `@Data`, `@EqualsAndHashCode` ou `@ToString` irrestrito em entidades JPA**, pois eles quebram o lazy loading do Hibernate, podem causar `LazyInitializationException` ou loops infinitos em relacionamentos bidirecionais, além de invalidar o contrato de `equals`/`hashCode` antes/depois da persistência.
  - Em entidades, prefira `@Getter` e implemente `equals`/`hashCode` baseados apenas no identificador (`uuid`) ou métodos de negócio explícitos.

---

## 4. Entidades de Domínio e Inicialização Defensiva

Mesmo utilizando anotações do Jakarta Persistence (`@Entity`), nossas entidades são **entidades de domínio ricas**, e não sacos anêmicos de getters e setters:

1. **Construtor Protegido para o JPA**:
   - `protected About() {}` (via `@NoArgsConstructor(access = AccessLevel.PROTECTED)`) existe exclusivamente para o proxy do Hibernate.
2. **Higienização Centralizada (`Sanitizer.sanitize`)**:
   - Todas as entidades devem higienizar campos `String` no construtor utilizando a classe utilitária estática `br.com.thiagoodev.api.domain.utils.Sanitizer` (`import static br.com.thiagoodev.api.domain.utils.Sanitizer.sanitize;`), que converte `null` em `""` e aplica `.trim()`.
   - Classes utilitárias puras como `Sanitizer` devem ser declaradas como `public final class` com construtor `private` vazio para impedir instanciação e herança indevida.
3. **Construtor de Domínio Completo com Validação**:
   - Recebe os campos da entidade, aplica `sanitize(...)` nas strings, garante que listas nunca fiquem `null` (`this.skills = skills != null ? new ArrayList<>(skills) : new ArrayList<>()`), inicializa `createdAt`/`updatedAt` (`LocalDateTime.now()`) se nulos e chama `this.validate()` ao final.
4. **Validação Explícita (`validate()`)**:
   - Verifica as invariantes de negócio (campos obrigatórios, URLs, datas, limites de tamanho, etc.) e lança `DomainValidationException` com mensagem clara caso alguma regra seja violada.
5. **Encapsulamento de Coleções**:
   - Métodos que manipulam coleções (como `addSkill` e `removeSkill` em `Project`) devem encapsular a regra de não-duplicidade e remoção segura (`removeIf`).

---

## 5. Estruturas de Controle e Fluxo Limpo

### Guard Clauses e Early Return
Evite aninhamentos profundos de `if/else`. Valide pré-condições no topo do método e lance a exceção ou retorne imediatamente:

```java
// Correto: cláusula de guarda direta
if (name == null || name.isBlank()) {
    throw new DomainValidationException("name cannot be empty");
}
```

### Uso de `Optional<T>`
- Retorne `Optional<T>` em métodos de busca de repositório que podem não encontrar registro.
- Na camada de aplicação (`UseCase`), consuma o `Optional` de forma fluente com `.orElseThrow(() -> new ResourceNotFoundException(...))`.
- Nunca utilize `Optional` como parâmetro de método ou atributo de entidade.

---

## 6. Tratamento de Exceções e Respostas HTTP

- **Exceções de Domínio são Unchecked (`RuntimeException`)**:
  - `DomainValidationException` e `ResourceNotFoundException` estendem `RuntimeException`.
  - Isso mantém as assinaturas dos casos de uso e interfaces limpas e integra-se perfeitamente ao rollback transacional padrão do Spring (`@Transactional`).
- **Controllers Limpos (Sem `try/catch`)**:
  - Os métodos dos `@RestController` não devem conter blocos `try/catch` para mapear erros de negócio.
  - Toda tradução de exceção para status HTTP ocorre centralmente no `GlobalExceptionHandler` (`@RestControllerAdvice`).

---

## 7. Comentários e Código Limpo

- **Código Autodocumentado**: Nomes claros de classes, métodos e variáveis eliminam a necessidade de comentários explicativos óbvios.
- **Sem Código Comentado**: Não mantenha código morto comentado no repositório.
- **Regra Estrita em Testes**: Arquivos de teste (`*Test.java`) seguem a regra de **zero comentários** (conforme [testing.md](testing.md)). A separação entre Arrange, Act e Assert é feita exclusivamente por uma linha em branco entre os blocos.
