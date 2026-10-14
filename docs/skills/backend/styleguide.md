# Backend Go Styleguide

Este guia define as convenções de estilo, padrões idiomáticos e práticas de engenharia Go adotadas no backend (`services/app-api`), baseando-se estritamente nas diretrizes oficiais do [Effective Go](https://go.dev/doc/effective_go) e nas evoluções contemporâneas da linguagem (Go 1.21+).

---

## 1. Formatação (Formatting)

No ecossistema Go, discussões sobre formatação são resolvidas por ferramentas automatizadas:

- **`gofmt` é a lei**: Todo código Go deve ser formatado via `gofmt` ou `go fmt ./...`.
- **Indentação com Tabs**: O Go utiliza tabs para indentação e espaços apenas para alinhamento vertical pontual gerado pelo próprio `gofmt`.
- **Sem limite artificial de linha**: Não quebre linhas desnecessariamente para atender a limites arbitrários (como 80 colunas). Se uma linha ficar longa demais, quebre-a de forma natural e indente com um tab extra.
- **Parênteses limpos**: Estruturas de controle (`if`, `for`, `switch`) **nunca** utilizam parênteses ao redor das condições.

---

## 2. Nomenclatura (Names)

A nomenclatura em Go comunica visibilidade, escopo e propósito de forma concisa.

### Pacotes (Package names)
- Nomes curtos, em minúsculas e de palavra única (`entities`, `postgres`, `http`).
- **Sem underscores** (`snake_case`) e **sem MixedCaps** nos nomes de pacotes.
- Evite redundâncias entre pacote e tipo:
  - Prefira `entities.Project` em vez de `entities.ProjectEntity`.
  - Prefira `http.Server` em vez de `http.HTTPServer`.

### Getters e Setters
- O Go não utiliza o prefixo `Get` para métodos leitores:
  - Correto: `project.Owner()` e `project.SetOwner(...)`.
  - Incorreto: `project.GetOwner()`.

### Interfaces
- Interfaces com um único método recebem o sufixo `-er` associado ao verbo da ação:
  - `Reader`, `Writer`, `Validator`, `Closer`, `Finder`.
- Mantenha interfaces pequenas e focadas no consumidor (regra: *accept interfaces, return structs*).

### MixedCaps e Siglas
- Identificadores utilizam `CamelCase` ou `mixedCaps` (nunca `snake_case`).
- Siglas e acrônimos mantêm a consistência de maiúsculas/minúsculas:
  - Correto: `UUID`, `HTTP`, `URL`, `ID`, `API`, `serveHTTP`, `appAPI`.
  - Incorreto: `Uuid`, `Http`, `Url`, `Id`, `Api`.

### Sem Abreviações (No Abbreviations)
- **Regra Estrita**: Não utilize abreviações no código para variáveis, campos de structs, parâmetros de funções ou dependências injetadas. Sempre utilize nomes completos e explícitos:
  - Correto: `aboutRepository`, `projectRepository`, `workExperience`, `education`, `index`, `databaseError`.
  - Incorreto: `aboutRepo`, `projectRepo`, `workExp`, `exp`, `edu`, `idx`, `dbErr`.
- **Exceções Permitidas (Padrão Go Standard Library)**:
  - `ctx` para `context.Context`.
  - `t` para `*testing.T`.
  - Receivers de métodos de structs de 1 ou 2 letras (`(p *Project)`, `(a *About)`).
  - Variáveis padrão de iteração curta (`err`, `ok`, `i`, `v`).

### Imports e Renomeação (Named Imports)
- **Código de Produção**: Evite renomear pacotes importados (*named imports* ou *aliases*). Sempre prefira o nome original e natural do pacote. Renomeações em código de produção são restritas a casos de **extrema necessidade**, especificamente para resolver colisões inevitáveis de nomes (por exemplo, quando o pacote corrente compartilha o mesmo nome de um pacote externo importado, como `usecases`, ou ao importar pacotes homônimos como `math/rand` e `crypto/rand`).
- **Código de Testes**: Imports nomeados são permitidos para enriquecer a semântica e expressividade dos cenários de teste (ex.: `mocks`, `domainerrors`).

---

## 3. Estruturas de Controle e Fluxo Idiomático

### Guard Clauses e "Happy Path" à Esquerda
Evite aninhamentos profundos de `if/else`. Trate erros e casos de borda imediatamente com retorno antecipado (*early return*), mantendo o fluxo principal alinhado à esquerda:

```go
// Incorreto: aninhamento desnecessário
if err == nil {
    // fluxo principal
    return data, nil
} else {
    return nil, err
}

// Correto: happy path alinhado à esquerda
if err != nil {
    return nil, err
}
return data, nil
```

### Inicialização em Declarações `if`
Utilize o escopo reduzido para variáveis temporárias de verificação:

```go
if err := entity.Validate(); err != nil {
    return fmt.Errorf("failed to validate entity: %w", err)
}
```

### `switch` e Type Switches
- O `switch` em Go não necessita de `break` explícito entre os casos (não há *fallthrough* automático).
- Casos podem ser listas separadas por vírgula (`case 'a', 'b', 'c':`).
- Utilize *type switches* para verificar tipos concretos em interfaces:

```go
switch v := value.(type) {
case string:
    return strings.TrimSpace(v)
case int:
    return strconv.Itoa(v)
default:
    return fmt.Sprintf("%v", v)
}
```

---

## 4. Funções e Retornos

### Múltiplos Retornos
- O padrão canônico para operações falíveis é retornar `(T, error)`.
- Quando uma função retornar erro, os valores de dados devem ser retornados em seu estado zero (`nil`, `""`, `0`):
  ```go
  return nil, err
  ```

### Named Result Parameters
- Use parâmetros nomeados de retorno apenas quando agregarem valor de documentação explícito em funções complexas.
- Evite retornos vazios (*naked returns*) em funções com mais de algumas linhas, pois prejudicam a legibilidade.

### `defer` para Limpeza Garantida
- Execute limpezas de recursos (`file.Close()`, `rows.Close()`, `tx.Rollback()`, destravamento de `sync.Mutex`) imediatamente após a inicialização bem-sucedida:
  ```go
  f, err := os.Open(filename)
  if err != nil {
      return err
  }
  defer f.Close()
  ```
- Lembre-se: `defer` segue ordem LIFO (Last-In, First-Out) e avalia seus argumentos imediatamente no ponto da declaração.

---

## 5. Estruturas de Dados e Alocação

### `new` vs. `make`
- `new(T)`: Aloca memória zerada para um tipo `T` e retorna seu ponteiro `*T`. Raramente necessário no dia a dia quando literais compostos estão disponíveis.
- `make(T, args)`: Cria e inicializa exclusivamente **slices**, **maps** e **channels**, retornando o valor inicializado (não um ponteiro).

### Literais Compostos (Composite Literals)
- Prefira literais compostos para construir structs e inicializar slices:
  ```go
  // Inicialização explícita e legível
  project := &entities.Project{
      UUID:        uuid.NewString(),
      Name:        name,
      Skills:      []entities.Skill{},
      IsHighlight: false,
  }
  ```

### Slices Modernos (Go 1.21+) e Prevenção de Memory Loitering
- **Pacote `slices`**: Sempre utilize o pacote nativo `slices` da standard library para manipulações funcionais e seguras:
  - `slices.ContainsFunc` para verificar pertinência por predicado.
  - `slices.IndexFunc` para localizar posições.
  - `slices.Delete` para remover elementos sem causar *Memory Loitering* (o pacote limpa as referências no array subjacente para liberar o Garbage Collector).
- **Defensive Slices**: Garanta que slices exportados em entidades ou DTOs nunca fiquem em estado `nil` quando representam listas vazias (`[]entities.Skill{}`). Isso garante previsibilidade na serialização JSON (`[]` em vez de `null`).

### Maps e Comma-ok Idiom
- Para testar a presença de uma chave em um map sem ambiguidade com o valor zero, utilize sempre a sintaxe *comma-ok*:
  ```go
  if val, ok := cache[key]; ok {
      return val, nil
  }
  ```

---

## 6. Métodos: Ponteiros vs. Valores (Pointers vs. Values)

A escolha entre *pointer receiver* (`*T`) e *value receiver* (`T`) deve seguir critérios consistentes:

1. **Mutação de Estado**: Se o método precisa alterar o estado interno do objeto, o receiver **deve** ser ponteiro.
2. **Tamanho e Cópia**: Se a struct for grande ou contiver campos pesados, use ponteiro para evitar cópias de memória a cada invocação.
3. **Consistência do Tipo**: Se qualquer método de um tipo exigir ponteiro, padronize todos os métodos desse tipo como ponteiros.
4. **Value Receivers**: Utilize receivers por valor apenas para pequenos tipos imutáveis (Value Objects pequenos, ex: `type Money struct { cents int64 }`).

---

## 7. Interfaces e Asserções de Tipo

- **Interfaces Pequenas**: Defina interfaces onde elas são consumidas (*consumer-driven*), não onde são implementadas.
- **Verificação de Implementação em Tempo de Compilação (Compile-Time Interface Check)**:
  Como o Go utiliza interfaces implícitas (sem a palavra-chave `implements`), declare sempre uma verificação com o identificador em branco (`_`) logo acima da struct concreta. Isso faz o compilador validar as assinaturas estaticamente, com custo de runtime zero:
  ```go
  var _ domainusecases.GetAboutUseCase = (*GetAboutUseCaseImpl)(nil)
  var _ repositories.ProjectRepository = (*PostgresProjectRepository)(nil)
  ```
  *(Sempre que este padrão for introduzido em novos componentes, o agente deve relembrar brevemente seu funcionamento ao proprietário do repositório).*
- **Type Assertion com Comma-ok**:
  Nunca execute type assertion simples (`x.(T)`) sem validação quando houver incerteza sobre o tipo subjacente, pois isso gera *panic*. Use sempre:
  ```go
  if str, ok := value.(string); ok {
      // uso seguro de str
  }
  ```

---

## 8. Tratamento de Erros (Errors as Values)

O Go trata erros como valores explícitos inspecionáveis, não como exceções mágicas.

### Sentinel Errors e Wrapping com `%w`
- Erros de domínio genéricos devem ser definidos como variáveis de sentinela em nível de pacote:
  ```go
  var ErrDomainEntityValidation = errors.New("domain entity validation error")
  ```
- Contextualize o erro ao propagá-lo utilizando `fmt.Errorf` com o verbo `%w`:
  ```go
  if err := validateTitle(title); err != nil {
      return fmt.Errorf("%w: title cannot be empty", ErrDomainEntityValidation)
  }
  ```
- Isso permite que camadas superiores (como handlers HTTP) inspecionem a raiz do erro com `errors.Is`:
  ```go
  if errors.Is(err, errors.ErrDomainEntityValidation) {
      // Mapeia deterministicamente para HTTP 400 Bad Request
  }
  ```

### Proibição de Panic em Regras de Negócio
- `panic` e `recover` são reservados estritamente para falhas catastróficas não recuperáveis de tempo de inicialização (ex: falha ao compilar regex estático com `regexp.MustCompile`).
- Código de negócio, casos de uso, repositórios e endpoints **nunca** devem disparar `panic`. Retorne erros tipados.

---

## 9. Comentários e Documentação

- **Godoc Idiomático**: Comentários de funções, structs e interfaces públicas devem iniciar exatamente com o nome do identificador e formar sentenças completas em inglês ou português coerente:
  ```go
  // Project representa a entidade raiz de um projeto do portfólio.
  type Project struct { ... }
  ```
- **Sem Código Comentado**: Não mantenha blocos de código comentados no repositório. O histórico de commits do Git é o local de preservação de versões anteriores.
- **Regra de Testes**: Código de teste (`*_test.go`) segue o padrão estrito de **zero comentários** (conforme [testing.md](testing.md)). A divisão AAA é feita unicamente através de blocos com linhas em branco.

