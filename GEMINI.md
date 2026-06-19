# GEMINI.md — Guia do Projeto Portfolio (Go + Templates)

## Sobre este arquivo

Este arquivo contém todo o contexto do projeto de portfolio de Thiago, construído com Go no backend usando `html/template` para SSR. Foi gerado a partir de uma sessão de aprendizado profundo com o Claude e serve como guia para continuar o desenvolvimento.

**Instruções para o agente:**
- Responda sempre em **Português do Brasil**
- Sempre que possível, **pesquise a documentação atualizada** dos packages usados
- Oriente sobre **arquitetura, padrões de projeto** com foco em didática
- Aprofunde conceitos de computação: **memória, algoritmos, ponteiros, concorrência**
- Seja **didático e explique o porquê** de cada decisão, não apenas o que fazer
- Compare sintaxe Go com **Dart e Rust** sempre que possível — são as linguagens que Thiago domina
- Siga o plano de **migração para produção** descrito neste arquivo
- Continue o trabalho iniciado — não repita o que já foi feito, avance a partir daqui

---

## Perfil do Desenvolvedor

- **Nome:** Thiago
- **GitHub:** github.com/cthiagoodev
- **Experiência:** 5 anos de Flutter/Dart (nível sênior)
- **Aprendendo:** Go Lang (backend principal) e Rust (aposta de longo prazo)
- **Deixou de lado:** Java/Spring Boot — muito foco em abstrações do framework, pouco em fundamentos
- **Objetivo:** Ser full stack com Go + Dart, entendendo os fundamentos reais da computação
- **Perfil de aprendizado:** Quer entender o **porquê** de cada decisão, não apenas copiar código. Compara Go com Dart e Rust para solidificar conceitos.

---

## O Projeto

**Portfolio pessoal** sendo refatorado do zero.

- **Stack anterior:** Spring Boot + Java + Template Engine + Docker
- **Nova stack:** Go + `html/template` (SSR) — sem frontend separado
- **Banco de dados:** PostgreSQL via `pgx/v5` (interface nativa, não `database/sql`)
- **Router:** `go-chi/chi/v5`
- **Hospedagem:** GitHub + VM própria
- **Docker Compose:** usado **somente em produção**. Em desenvolvimento, roda direto na máquina.

### Por que SSR com templates e não API REST + Jaspr

Decisão consciente tomada durante o desenvolvimento:
- Portfolio é um projeto simples de vitrine — manutenção mínima após concluído
- Aprender duas tecnologias ao mesmo tempo (Go backend + Dart/Jaspr frontend) seria complexidade desnecessária
- `html/template` é nativo do Go — sem dependências externas
- Foco total em aprender Go profundamente

---

## Repositório

- **URL:** github.com/cthiagoodev/thiagoodev-portfolio
- **Estratégia de branches:**
    - `main` — v1 antiga (Spring Boot), site no ar, **não tocar**
    - `v2` — desenvolvimento da nova stack Go
    - Tag `v1.0.0` preserva o histórico da stack antiga
    - Quando pronto: merge da `v2` na `main`

---

## Módulo Go

```
module github.com/cthiagoodev/thiagoodev-portfolio

go 1.26.2
```

**Por que URL como nome do módulo:** convenção universal do ecossistema Go. Não há registry centralizado como pub.dev ou crates.io — a URL do repositório é o identificador único global.

---

## Dependências instaladas

```
require (
    github.com/jackc/pgx/v5 v5.9.2
    github.com/go-chi/chi/v5 v5.2.5
    github.com/joho/godotenv v1.5.1
)
```

---

## Estrutura de pastas

```
server/
├── cmd/
│   └── api/
│       ├── main.go        ← entrypoint
│       └── router.go      ← montagem do router
├── internal/
│   ├── common/            ← compartilhado entre features
│   │   ├── errors.go      ← sentinel errors globais
│   │   ├── handle_error.go← helper de resposta de erro
│   │   ├── parse_db_error.go ← traduz erros do pgx para domínio
│   │   ├── db.go          ← inicialização do pgxpool
│   │   └── config.go      ← leitura de variáveis de ambiente
│   ├── about/             ← feature: sobre mim
│   │   ├── entity.go
│   │   ├── repository.go
│   │   ├── database_repository.go
│   │   ├── usecase.go
│   │   ├── handler.go
│   │   └── routes.go
│   ├── contact/           ← feature: contato
│   │   ├── entity.go
│   │   ├── repository.go
│   │   ├── database_repository.go
│   │   ├── usecase.go
│   │   ├── handler.go
│   │   └── routes.go
│   ├── experience/        ← a implementar
│   ├── project/           ← a implementar
│   ├── talk/              ← a implementar
│   └── community/         ← a implementar
├── templates/             ← SSR com html/template
│   ├── base.html          ← layout base
│   ├── partials/
│   │   ├── header.html
│   │   ├── nav.html
│   │   └── footer.html
│   └── pages/
│       ├── about.html
│       ├── experience.html
│       ├── projects.html
│       ├── talks.html
│       ├── community.html
│       └── contact.html
├── migrations/            ← SQL migrations (golang-migrate)
│   ├── 000001_create_about.up.sql
│   ├── 000001_create_about.down.sql
│   ├── 000002_create_technology.up.sql
│   ├── 000002_create_technology.down.sql
│   └── 000003_create_contact.up.sql
├── .env                   ← variáveis de ambiente (não commitar)
├── go.mod
└── go.sum
```

---

## Arquivos implementados

### `internal/common/errors.go`

```go
package common

import "errors"

var ErrNotFound         = errors.New("not found")
var ErrInvalidData      = errors.New("invalid data")
var ErrConnection       = errors.New("connection error")
var ErrAlreadyExists    = errors.New("already exists")
var ErrInvalidReference = errors.New("invalid reference")
```

### `internal/common/parse_db_error.go`

```go
package common

import (
    "errors"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgconn"
)

func ParseDbError(err error) error {
    var connErr *pgconn.ConnectError
    if errors.As(err, &connErr) {
        return ErrConnection
    }

    if errors.Is(err, pgx.ErrNoRows) {
        return ErrNotFound
    }

    return err
}
```

### `internal/common/handle_error.go`

```go
package common

import (
    "encoding/json"
    "errors"
    "net/http"
)

func HandleError(w http.ResponseWriter, err error) {
    w.Header().Set("Content-Type", "application/json")

    statusCode := http.StatusInternalServerError

    switch {
    case errors.Is(err, ErrNotFound):
        statusCode = http.StatusNotFound
    case errors.Is(err, ErrInvalidData):
        statusCode = http.StatusBadRequest
    case errors.Is(err, ErrConnection):
        statusCode = http.StatusServiceUnavailable
    default:
        statusCode = http.StatusInternalServerError
    }

    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(map[string]any{
        "status": statusCode,
        "error":  err.Error(),
    })
}
```

### `internal/common/db.go`

```go
package common

import (
    "context"

    "github.com/jackc/pgx/v5/pgxpool"
)

func NewDb(url string) (*pgxpool.Pool, error) {
    pool, err := pgxpool.New(context.Background(), url)
    if err != nil {
        return nil, err
    }
    return pool, nil
}
```

### `internal/common/config.go`

```go
package common

import (
    "log"
    "os"
)

type Config struct {
    DbUrl string
    Port  string
}

func NewConfig() *Config {
    dbUrl := os.Getenv("DATABASE_URL")

    if dbUrl == "" {
        log.Fatal("DATABASE_URL is required")
    }

    return &Config{
        DbUrl: dbUrl,
        Port:  "8000",
    }
}
```

### `internal/about/entity.go`

```go
package about

type About struct {
    Uuid       string       `json:"uuid"`
    Name       string       `json:"name"`
    Bio        string       `json:"bio"`
    Photo      string       `json:"photo"`
    Curriculum string       `json:"curriculum"`
    Linkedin   string       `json:"linkedin"`
    GitHub     string       `json:"github"`
    Stack      []Technology `json:"stack"`
    City       string       `json:"city"`
    State      string       `json:"state"`
}

func (a About) IsValid() bool {
    return a.Name != "" &&
        a.Bio != "" &&
        a.Photo != "" &&
        a.Linkedin != "" &&
        a.GitHub != "" &&
        a.City != "" &&
        a.State != ""
}

func (a About) HasStack() bool {
    return len(a.Stack) > 0
}

type Technology struct {
    Uuid string `json:"uuid"`
    Name string `json:"name"`
}
```

### `internal/about/repository.go`

```go
package about

import "context"

type Repository interface {
    Find(ctx context.Context) (About, error)
}
```

### `internal/about/database_repository.go`

```go
package about

import (
    "context"

    "github.com/cthiagoodev/thiagoodev-portfolio/internal/common"
    "github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseRepository struct {
    pool *pgxpool.Pool
}

func NewDatabaseRepository(pool *pgxpool.Pool) *DatabaseRepository {
    return &DatabaseRepository{pool}
}

func (r *DatabaseRepository) Find(ctx context.Context) (About, error) {
    query := `
        SELECT uuid, name, bio, photo, curriculum,
               linkedin, github, city, state
        FROM about
        LIMIT 1
    `
    row := r.pool.QueryRow(ctx, query)

    about := About{}
    err := row.Scan(
        &about.Uuid,
        &about.Name,
        &about.Bio,
        &about.Photo,
        &about.Curriculum,
        &about.Linkedin,
        &about.GitHub,
        &about.City,
        &about.State,
    )

    if err != nil {
        return About{}, common.ParseDbError(err)
    }

    return about, nil
}
```

### `internal/about/usecase.go`

```go
package about

import "context"

type UseCase struct {
    repo Repository
}

func NewUseCase(repo Repository) *UseCase {
    return &UseCase{repo: repo}
}

func (u *UseCase) Get(ctx context.Context) (About, error) {
    return u.repo.Find(ctx)
}
```

### `internal/about/handler.go`

**Atenção:** o handler ainda retorna JSON. Precisa ser migrado para `html/template`. Ver seção "Próximos passos".

```go
package about

import (
    "encoding/json"
    "net/http"

    "github.com/cthiagoodev/thiagoodev-portfolio/internal/common"
)

type Handler struct {
    useCase *UseCase
}

func NewHandler(useCase *UseCase) *Handler {
    return &Handler{useCase}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    about, err := h.useCase.Get(r.Context())

    if err != nil {
        common.HandleError(w, err)
        return
    }

    json.NewEncoder(w).Encode(about)
}
```

### `internal/about/routes.go`

```go
package about

import (
    "github.com/go-chi/chi/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(mx *chi.Mux, pool *pgxpool.Pool) {
    repository := NewDatabaseRepository(pool)
    useCase := NewUseCase(repository)
    handler := NewHandler(useCase)

    mx.Get("/about", handler.Get)
}
```

### `internal/contact/entity.go`

```go
package contact

type Contact struct {
    Uuid  string `json:"uuid"`
    Phone string `json:"phone"`
    Email string `json:"email"`
}

func (c Contact) IsValid() bool {
    return c.Phone != "" && c.Email != ""
}
```

### `internal/contact/repository.go`

```go
package contact

import "context"

type Repository interface {
    Find(ctx context.Context) (Contact, error)
}
```

### `cmd/api/main.go`

```go
package main

import (
    "log"
    "net/http"

    "github.com/cthiagoodev/thiagoodev-portfolio/internal/common"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load(".env")

    config := common.NewConfig()
    pool, err := common.NewDb(config.DbUrl)

    if err != nil {
        log.Fatal("error on init server: ", err)
    }

    router := NewRouter(pool)

    serverError := http.ListenAndServe(":"+config.Port, router)
    if serverError != nil {
        log.Fatal("error on init server: ", serverError)
    }
}
```

### `cmd/api/router.go`

```go
package main

import (
    "github.com/cthiagoodev/thiagoodev-portfolio/internal/about"
    "github.com/go-chi/chi/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {
    router := chi.NewRouter()

    router.Route("/", func(r chi.Router) {
        about.RegisterRoutes(router, pool)
    })

    return router
}
```

---

## Arquivo `.env`

```
DATABASE_URL=postgres://admin:admin@localhost:5432/portfolio_db
PORT=8000
```

---

## Migrations — SQL

### `000001_create_about.up.sql`

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE about (
    uuid        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(255) NOT NULL,
    bio         TEXT NOT NULL,
    photo       VARCHAR(500) NOT NULL,
    curriculum  VARCHAR(500),
    linkedin    VARCHAR(255) NOT NULL,
    github      VARCHAR(255) NOT NULL,
    city        VARCHAR(100) NOT NULL,
    state       VARCHAR(2)   NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);
```

### `000002_create_technology.up.sql`

```sql
CREATE TABLE technology (
    uuid        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    about_uuid  UUID NOT NULL REFERENCES about(uuid) ON DELETE CASCADE,
    name        VARCHAR(100) NOT NULL
);
```

### `000003_create_contact.up.sql`

```sql
CREATE TABLE contact (
    uuid   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    phone  VARCHAR(20)  NOT NULL,
    email  VARCHAR(255) NOT NULL
);
```

---

## Conceitos já aprendidos

### Módulo vs Package
- **Módulo** — projeto inteiro, tem `go.mod`, unidade de versionamento
- **Package** — cada pasta com arquivos `.go`, declarado com `package nome`
- Imports usam o caminho completo: `github.com/user/repo/internal/about`

### Tipos e Zero Values
- Todo tipo tem zero value: `string` → `""`, `int` → `0`, `bool` → `false`, ponteiro → `nil`
- Go não tem `null safety` como Dart nem `Option<T>` como Rust
- Ponteiro `*T` é a forma de representar ausência — mas usar com cautela em entidades de domínio

### Ponteiros
- `&valor` — pega o endereço de memória (cria ponteiro)
- `*ptr` — acessa o valor apontado (derreferencia)
- `*Tipo` na declaração — declara que é um ponteiro para aquele tipo
- Structs são auto-derreferenciadas: `ptr.Name` equivale a `(*ptr).Name`
- Diferença de Rust: Go não tem ownership/borrow checker — GC cuida da memória em runtime

### Interfaces
- Implícitas — sem `implements`. Se o tipo tem os métodos, satisfaz a interface
- A interface fica no lado de **quem consome**, não de quem implementa
- Interfaces já são referências internamente — não usar `*Interface`
- Permite Clean Architecture mais limpa que Java/Dart

### Erros
- Erros são valores retornados — sem `try/catch`
- Sempre o último valor de retorno: `(T, error)`
- `errors.Is` — compara valor (sentinel errors)
- `errors.As` — compara tipo em runtime, extrai dados
- `fmt.Errorf("contexto: %w", err)` — encapsula preservando o original
- `var ErrAlgo = errors.New("msg")` — sentinel error (prefixo `Err`)
- `type AlgoError struct{}` — erro com dados (sufixo `Error`)

### Context
- `context.Background()` — raiz, sem cancelamento, usado no startup
- `context.TODO()` — placeholder temporário
- `context.WithCancel()` — cancelamento manual
- `context.WithTimeout()` — cancela após duração
- Sempre primeiro parâmetro, sempre chamado `ctx`
- Propaga cancelamento por toda a cadeia: handler → usecase → repository → banco

### Goroutines e Concorrência
- Go não tem `async/await` — toda função parece síncrona
- Goroutines começam com ~2KB de stack (vs ~1MB de thread do SO)
- `go func()` — dispara goroutine
- Channels para comunicação entre goroutines
- Context para cancelamento entre goroutines

### Structs e Métodos
- Receiver de valor `(a About)` — recebe cópia, só leitura
- Receiver de ponteiro `(a *About)` — recebe referência, pode modificar
- Construtor por convenção: `func NewAlgo(...) *Algo`
- Campos privados com letra minúscula, públicos com maiúscula

### Collections
- `[]T` — slice (equivalente a `List<T>` do Dart)
- `map[K]V` — map (equivalente a `Map<K,V>` do Dart)
- `map[T]struct{}` — set (Go não tem Set nativo)
- `for i, v := range slice` — iteração
- `len()` — tamanho (funciona em nil slice sem panic)

### Packages importantes
- `encoding/json` — Marshal/Unmarshal, struct tags `json:"nome"`
- `net/http` — servidor HTTP, Handler interface
- `context` — gerenciamento de ciclo de vida
- `errors` — Is, As, New
- `fmt` — Errorf com %w
- `os` — Getenv
- `log` — Fatal (imprime + os.Exit(1))

---

## Próximos passos — em ordem de prioridade

### 1. Migrar handlers para `html/template`

O handler atual retorna JSON. Precisa ser migrado para renderizar HTML.

**Conceitos a ensinar:**
- `html/template` vs `text/template`
- `{{define "nome"}}` — declara template nomeado
- `{{template "nome" .}}` — inclui template
- `{{block "nome" .}}` — bloco sobrescrevível (como abstract)
- `{{range .Stack}}` — iteração
- `{{if .Curriculum}}` — condição
- `template.ParseGlob` — carrega múltiplos arquivos
- `tmpl.ExecuteTemplate(w, "base", data)` — renderiza template específico

**Estrutura do sistema de templates:**

```
templates/
├── base.html           ← layout base com {{block "content" .}}
├── partials/
│   ├── header.html
│   ├── nav.html
│   └── footer.html
└── pages/
    ├── about.html      ← {{define "content"}}
    ├── experience.html
    └── ...
```

### 2. Arquivos estáticos

Servir CSS, JS, imagens:

```go
// no router
fileServer := http.FileServer(http.Dir("./static"))
r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
```

```
server/
└── static/
    ├── css/
    │   └── main.css
    └── img/
```

### 3. Implementar features restantes

Em ordem de prioridade definida por Thiago:

1. `contact/` — database_repository, usecase, handler, routes
2. `experience/` — entity, repository, database_repository, usecase, handler, routes
3. `project/` — mesma estrutura
4. `talk/` — mesma estrutura
5. `community/` — mesma estrutura

### 4. Buscar tecnologias do `about`

A tabela `technology` tem FK para `about`. O `Find` do about precisa de uma segunda query para buscar as tecnologias e montar o slice `Stack`:

```go
// após buscar o about
rows, err := r.pool.Query(ctx,
    "SELECT uuid, name FROM technology WHERE about_uuid = $1",
    about.Uuid,
)
// iterar com rows.Next() e rows.Scan()
```

### 5. Migrations

Instalar `golang-migrate` e criar os arquivos SQL:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate create -ext sql -dir migrations -seq create_about
migrate -path migrations -database $DATABASE_URL up
```

### 6. Testes unitários

- `entity_test.go` — testa `IsValid()` e `HasStack()`
- `usecase_test.go` — mock do repository, testa orquestração
- `handler_test.go` — testa status codes e resposta HTML

### 7. Makefile

```makefile
run:
    go run ./cmd/api/

build:
    go build -o bin/server ./cmd/api/

migrate-up:
    migrate -path migrations -database $(DATABASE_URL) up

migrate-down:
    migrate -path migrations -database $(DATABASE_URL) down 1

test:
    go test ./...
```

### 8. Docker para produção

```dockerfile
# Dockerfile
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/api/

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
EXPOSE 8000
CMD ["./server"]
```

```yaml
# docker-compose.yml (na raiz do monorepo)
services:
  server:
    build:
      context: ./server
    ports:
      - "8000:8000"
    environment:
      - DATABASE_URL=${DATABASE_URL}
      - PORT=8000
    depends_on:
      db:
        condition: service_healthy

  db:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER}"]
      interval: 5s
      timeout: 5s
      retries: 5

volumes:
  postgres_data:
```

---

## Padrões e decisões arquiteturais tomadas

| Decisão | Escolha | Motivo |
|---|---|---|
| Estrutura | Feature-based | Mais natural para portfolio — pensa em "o que tem" não em "qual camada" |
| ORM | Nenhum — SQL puro com pgx | Go é direto, sem magic, você controla tudo |
| Interface pgx | Nativa (pgxpool) | Projeto só usa Postgres, sem necessidade de abstração database/sql |
| Injeção de dependência | Manual via construtor | Sem framework — Go não precisa, é explícito e legível |
| Erros de banco | ParseDbError centralizado | Traduz erros técnicos para domínio — Clean Architecture |
| Frontend | SSR com html/template | Portfolio simples, manutenção mínima, sem stack dupla |
| Router | chi | Leve, idiomático, baseado em net/http, suporte a middleware |
| Migrations | golang-migrate | Padrão do ecossistema Go |
| Entidades | Valor, não ponteiro | Ponteiro em entidade vaza detalhe técnico para o domínio |
| Interfaces | No lado do consumidor | Dependency Inversion real — domain não conhece infrastructure |

---

## Comparações Go vs Dart vs Rust (referência rápida)

| Conceito | Dart | Rust | Go |
|---|---|---|---|
| Nulo | `String?` | `Option<String>` | `*string` ou zero value |
| Erro | `try/catch` | `Result<T,E>` | `(T, error)` |
| Async | `async/await` | `async/await + tokio` | goroutines (sem keyword) |
| Classe | `class` | `struct + impl` | `struct + methods` |
| Interface | `abstract class` | `trait` | `interface` (implícita) |
| Herança | `extends` | não existe | não existe — embedding |
| Memória | GC | ownership + borrow | GC + escape analysis |
| Referência | sempre (objetos) | `&T` / `&mut T` | `*T` + `&valor` |
| Construtor | `ClassName(...)` | `Type::new(...)` | `func NewTipo(...)` |
| Privado | `_nome` | `nome` (minúsculo) | `nome` (minúsculo) |
| Público | `Nome` | `Nome` (maiúsculo) | `Nome` (maiúsculo) |
| Lista | `List<T>` | `Vec<T>` | `[]T` |
| Map | `Map<K,V>` | `HashMap<K,V>` | `map[K]V` |
| Set | `Set<T>` | `HashSet<T>` | `map[T]struct{}` |

---

## Comandos úteis

```bash
# desenvolvimento
go run ./cmd/api/

# verificar se compila
go build ./...

# adicionar dependência
go get github.com/alguma/lib

# sincronizar go.mod
go mod tidy

# baixar dependências
go mod download

# rodar testes
go test ./...
go test -v ./internal/about/

# migrations
migrate -path migrations -database $DATABASE_URL up
migrate -path migrations -database $DATABASE_URL down 1
```

---

## Filosofia Go — princípios que guiam as decisões

- **Explícito > implícito** — sem magic, sem anotações, sem reflexão desnecessária
- **Simples > complexo** — não crie abstração onde não adiciona valor
- **Composição > herança** — embedding e interfaces resolvem tudo
- **Erros são valores** — trate explicitamente, não esconda em exceptions
- **Comece simples, cresça conforme necessidade** — YAGNI aplicado à estrutura
- **Leia o código como prosa** — nomes claros, estrutura óbvia

---

*Gerado a partir de sessão de aprendizado com Claude — continue o trabalho a partir dos próximos passos descritos acima.*