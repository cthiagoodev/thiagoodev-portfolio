# Contrato HTMX

## Shell e fragments

`index.html` é o único documento completo. Endpoints de feature retornam
fragments destinados a `#content` e nunca incluem estrutura de documento.

Navegação padrão:

```html
hx-get="/about"
hx-target="#content"
hx-swap="innerHTML transition:true"
hx-push-url="true"
```

Atributos comuns podem ser herdados de containers somente quando o escopo for
óbvio e não afetar links externos.

## Ciclo de requisição

Toda navegação deve considerar idle, loading, success, empty, error e
unavailable quando aplicável. Feedback deve ser discreto, não bloquear o shell,
não criar atraso artificial e não provocar layout shift relevante.

- `aria-busy` deve refletir carregamentos significativos quando houver suporte;
- indicadores não podem depender apenas de animação;
- erros devem oferecer retry apenas quando a repetição for segura;
- respostas não devem expor erros internos.

## History

`hx-push-url="true"` registra o módulo carregado. Back e forward devem restaurar
conteúdo coerente. `hx-history-elt` limita o snapshot ao conteúdo principal.

## Contratos ainda não implementados

Deep links, refresh, active navigation, atualização de `<title>` e gerenciamento
de foco permanecem pendentes. Não criar workaround frontend sem aprovação.

Contrato recomendado ao backend para cada rota:

```text
Se HX-Request: true:
    retornar somente o fragment da feature

Se requisição normal:
    retornar o shell completo com a feature como conteúdo inicial,
    navegação coerente e metadata correspondente
```

O mecanismo futuro para active state e title deve ser escolhido entre resposta
OOB, headers/eventos HTMX ou JavaScript mínimo previamente aprovado.

## Foco e anúncios

Após swaps significativos, o usuário deve compreender que o conteúdo mudou. O
alvo preferido é o módulo ou seu heading, com foco visível e sem scroll
inesperado. A implementação está adiada até o contrato ser aprovado.

## Erros

O frontend pode fornecer o fragment visual de erro. O backend deve fornecer
status HTTP correto, contexto de retry e resposta adequada ao tipo de requisição.

## JavaScript

Não introduzir JavaScript para comportamento que HTML, CSS ou HTMX resolvam de
forma clara. Qualquer script deve ter responsabilidade única e documentada.

## Rotas atuais

Somente `/` e `/about` foram confirmadas no backend. `/experience`, `/projects`,
`/community` e `/contact` dependem de implementação do proprietário.
