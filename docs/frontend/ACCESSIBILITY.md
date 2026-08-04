# Acessibilidade

## Princípio

A identidade retro-computing nunca reduz legibilidade, operação por teclado ou
compreensão. Acessibilidade faz parte de cada feature.

## Semântica

- prefira elementos nativos a roles ARIA;
- cada view possui hierarquia lógica de headings sem saltos por motivo visual;
- links navegam; botões executam ações;
- listas, `dl`, artigos e seções devem refletir a estrutura do conteúdo;
- decoração deve ser ocultada de tecnologias assistivas.

## Teclado e foco

- toda interação funciona por teclado;
- foco é sempre visível;
- não remova outline sem substituição equivalente;
- o shell deve oferecer skip link para `#content`;
- target de toque deve permanecer confortável;
- gerenciamento de foco HTMX segue o contrato e não será improvisado.

## Conteúdo dinâmico

Loading, mudanças, erros e retry precisam ser compreensíveis sem visão ou
animação. Use `aria-live`, `aria-busy` e foco somente quando o comportamento for
necessário e testado; ARIA excessiva também cria ruído.

## Contraste e texto

- texto significativo deve atender WCAG AA;
- `--color-text-faint` não serve a informação essencial;
- estado não depende apenas de cor;
- labels decorativas podem ser pequenas, conteúdo útil não;
- texto deve tolerar zoom e aumento de fonte sem clipping ou perda.

## Imagens e links

- imagens significativas recebem alt contextual;
- imagens decorativas usam alt vazio ou são ocultadas;
- reserve largura e altura/aspect-ratio conhecidos;
- links externos com nova aba usam `rel="noopener noreferrer"` e seguem uma
  convenção de comunicação consistente.

## Movimento

Respeite `prefers-reduced-motion`. Funcionalidade e compreensão não dependem de
animação, flicker, scan ou transição.

## Verificação mínima

- landmarks e headings;
- Tab e Shift+Tab;
- foco visível;
- ativação por teclado;
- nomes acessíveis;
- contraste;
- zoom a 200%;
- movimento reduzido;
- estados HTMX;
- ausência de overflow ou conteúdo inacessível.
