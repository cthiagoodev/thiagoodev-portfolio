# Estados de Interface

## Modelo

Interfaces dinâmicas consideram, quando aplicável: idle, loading, success, empty,
error, unavailable e disabled. Nem todo componente precisa de todos os estados.

## Loading

- comunica atividade sem atraso artificial;
- preserva layout quando viável;
- não bloqueia o shell para requests de fragment;
- não depende apenas de movimento;
- respeita `prefers-reduced-motion`;
- usa vocabulário curto, como `LOADING MODULE`.

## Success

O conteúdo normalmente substitui o loading sem animação comemorativa. Feedback
explícito é reservado a ações como envio de formulário, não à navegação comum.

## Empty

Explica o que está vazio, se é esperado e o que o visitante pode fazer. Nunca
renderize um painel vazio sem mensagem humana.

## Error

Inclui label de sistema opcional, heading, explicação em português e retry quando
seguro. Não expõe detalhes internos. Status HTTP e contexto de retry pertencem ao
contrato backend.

## Unavailable

Indica indisponibilidade intencional ou temporária conhecida. Não deve ser usado
para mascarar erro inesperado.

## Disabled

Controles indisponíveis comunicam estado visual e semanticamente. Opacidade
reduzida isolada não é suficiente e não pode comprometer leitura.

## Consistência

Todos os módulos reutilizam `.state-panel` e `.loading-state`. Features podem
fornecer mensagens específicas, mas não inventam estruturas visuais diferentes.

## Anatomia de state panel

1. system label;
2. heading humano;
3. explicação curta;
4. ação opcional e segura;
5. informação técnica somente se útil ao visitante.
