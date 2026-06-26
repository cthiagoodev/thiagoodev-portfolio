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
│   ├── api/
│   │   ├── main.go        ← entrypoint da API
│   │   └── router.go
│   └── migrate/
│       └── main.go        ← entrypoint da ferramenta de migração
├── internal/
│   ├── common/            ← compartilhado entre features
│   │   ├── errors.go
│   │   ├── handle_error.go
│   │   ├── parse_db_error.go
│   │   ├── db.go
│   │   └── config.go
│   ├── about/             ← feature: sobre mim
│   │   ├── entity.go
│   │   ├── repository.go
│   │   ├── database_repository.go
│   │   ├── usecase.go
│   │   ├── handler.go
│   │   └── routes.go
│   ├── contact/           ← feature: contato
│   ├── experience/        ← a implementar
│   ├── project/           ← a implementar
│   ├── talk/              ← a implementar
│   └── community/         ← a implementar
├── templates/             ← SSR com html/template
│   ├── base.html
│   ├── partials/
│   │   ├── header.html
│   │   ├── nav.html
│   │   └── footer.html
│   └── pages/
│       ├── about.html
│       └── ...
├── migrations/            ← SQL migrations (golang-migrate)
│   ├── 000001_create_about.up.sql
│   ├── 000001_create_about.down.sql
│   └── ...
├── .env                   ← variáveis de ambiente (não commitar)
├── go.mod
└── go.sum
```

---

## Arquivos implementados

(Seções de arquivos implementados permanecem as mesmas)

---

## Conceitos já aprendidos

(Seções de conceitos permanecem as mesmas, com a adição abaixo)

### Ferramenta de Migração (`cmd/migrate`)
- **Design:** Binário Go isolado (`cmd/migrate/main.go`), separado da API principal (`cmd/api/main.go`).
- **Motivo:** Separação de Responsabilidades (SoC) e Segurança. A lógica de manipulação do esquema do banco (especialmente `down`) não deve ser exposta ou sequer incluída no binário da aplicação web.
- **Funcionamento:** Lê argumentos da linha de comando (`up`, `down`) para decidir a ação.
- **`golang-migrate/migrate`**: Biblioteca principal que orquestra o processo.
- **Drivers (Source e Database):**
    -   `source/file`: Ensina o `migrate` a ler arquivos `.sql` do disco.
    -   `database/postgres` com `pgx/stdlib`: Atua como um "adaptador". O `migrate` espera uma conexão `database/sql`, e o `stdlib` permite que ele use uma conexão `pgx` por baixo dos panos.
-   **Blank Imports (`import _ "..."`):** Usado para registrar os drivers na inicialização do programa sem precisar chamar nenhuma função deles diretamente.
-   **Tratamento de Erro:**
    -   `log.Fatal` é usado para encerrar o programa com um *exit code* diferente de zero, sinalizando falha para scripts de CI/CD.
    -   `migrate.ErrNoChange` é um erro esperado e tratado como sucesso, pois significa que o banco de dados já está atualizado.

---

## Próximos passos — em ordem de prioridade

(Seções de próximos passos permanecem as mesmas)

---

## Padrões e decisões arquiteturais tomadas

(Seções de padrões permanecem as mesmas)

---

## Comparações Go vs Dart vs Rust (referência rápida)

(Seções de comparações permanecem as mesmas)

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

# migrations (com nossa ferramenta customizada)
go run ./cmd/migrate/main.go up
go run ./cmd/migrate/main.go down
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