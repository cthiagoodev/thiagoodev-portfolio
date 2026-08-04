# Arquitetura Frontend

## Modelo da aplicação

`templates/index.html` é o único documento HTML completo. Ele contém metadata,
navegação global, `#content`, footer e dependências globais.

As features são fragments HTML inseridos em `#content`:

- `/about` → `about.html`
- `/experience` → `experience.html`
- `/projects` → `projects.html`
- `/community` → `community.html`
- `/contact` → `contact.html`

Fragments nunca contêm `doctype`, `html`, `head` ou `body`.

## Camadas CSS

### Foundation

Arquivos globais existentes: reset, tokens e base. Eles controlam normalização,
fundamentos visuais, shell, foco, movimento reduzido e comportamento global.

### Components

Padrões reutilizáveis e independentes de feature, em
`templates/static/css/components/`. Um componente não pode conhecer detalhes de
uma feature.

### Features

Composição específica de cada módulo, em `templates/static/css/features/`.
Uma feature não pode depender do CSS específico de outra.

`main.css` é o manifesto de imports, na ordem foundation → shell/components →
features. O projeto não exige bundler nesta fase.

## Níveis de reutilização

1. **Padrão CSS:** quando apenas aparência/comportamento visual é compartilhado.
2. **Convenção HTML:** quando o markup é curto e explícito.
3. **Go template:** somente quando markup significativo se repete, possui contrato
   de dados claro e a abstração melhora a leitura.

A segunda ocorrência real valida a criação de novos componentes. Os componentes
estruturais aprovados nesta fase são a base mínima para os módulos planejados.
Não reproduza uma arquitetura de SPA ou React em `html/template`.

## Organização de templates

Os globs atuais reconhecem apenas `*.html` e `components/*.html`. Portanto,
fragments e eventuais templates compartilhados permanecem diretamente em
`templates/components/`, sem subpastas. Alterar essa restrição depende de Go e
está fora da fronteira frontend.

Convenções:

- o nome definido pelo fragment corresponde ao nome executado pelo backend;
- o fragment possui um único elemento raiz com heading associado;
- IDs servem a âncoras, ARIA e alvos; classes servem a estilos;
- lógica de negócio não entra em templates;
- atributos HTMX permanecem explícitos ou são herdados de um container cujo
  escopo seja evidente.

## Nomenclatura

- BEM pragmático: `.module-header__title`, `.command-button--primary`;
- estados transitórios: `.is-active`, `.is-loading`, `.is-unavailable`;
- classes específicas recebem prefixo da feature: `.hero`, `.about-profile`;
- nomes de classes e templates em inglês;
- conteúdo segue `CONTENT_GUIDE.md`.

## JavaScript

HTML, CSS e HTMX são as ferramentas padrão. JavaScript só pode ser introduzido
para uma responsabilidade específica, documentada e previamente aprovada.

## Fronteira

O frontend adapta-se aos contratos existentes. Quando faltar rota, dado, header
ou comportamento do servidor, registre a necessidade conforme
`BACKEND_BOUNDARY.md` e não altere o backend.
