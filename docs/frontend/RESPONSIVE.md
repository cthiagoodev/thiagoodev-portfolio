# Design Responsivo

## Princípio

Responsividade é adaptação de conteúdo, não redução do desktop. A hierarquia deve
permanecer íntegra em qualquer largura.

Em espaço restrito, priorize: identidade, conteúdo principal, navegação, ações,
metadata e decoração. Simplifique efeitos antes de informação significativa.

## Estratégia de breakpoints

- use poucos breakpoints motivados por falha real de layout;
- mantenha media queries junto ao componente ou feature afetada;
- não crie breakpoint para cada componente;
- consolide valores apenas quando a repetição estiver comprovada;
- testes-alvo não são automaticamente breakpoints CSS.

## Navegação

Todos os destinos devem ser descobertos. Se houver scroll horizontal, deve
existir indicação perceptível de continuidade, suporte a teclado e touch targets
adequados. Ocultar scrollbar não pode ocultar a existência de outros links.

## Conteúdo e tipografia

- títulos não podem cortar nem causar overflow;
- texto de corpo mantém tamanho e line-height confortáveis;
- linhas não devem ficar excessivamente longas;
- labels significativas não podem se tornar microtipografia;
- ações principais permanecem alcançáveis.

## Decoração

Processor, meters, estrelas e diagnósticos podem reduzir ou reorganizar no mobile.
Nenhum conteúdo essencial pode desaparecer para preservar a composição desktop.

## Matriz de verificação

Verifique em torno de:

- 320px, 360px e 390px;
- tablet portrait e landscape;
- 1440px e wide desktop;
- zoom a 200%;
- conteúdo longo, ausente e nomes extensos;
- teclado, touch e movimento reduzido.

Critérios: sem overflow horizontal não intencional, targets alcançáveis, ordem de
leitura coerente, conteúdo prioritário visível e decoração controlada.
