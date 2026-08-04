# Design System

## Princípios

O sistema visual representa uma máquina de engenharia retro-computing. Conteúdo
e legibilidade têm prioridade sobre decoração. Reutilize tokens existentes antes
de criar novos valores.

## Cores semânticas

- near-black: background e void;
- soft white: texto principal;
- muted gray: texto secundário;
- violet: identidade, seleção e estrutura técnica;
- green: operação, disponibilidade e ação principal;
- warning/danger: somente estados correspondentes.

Não introduza nova cor de destaque sem justificativa. Estado nunca depende só de
cor. Texto significativo deve atender contraste adequado; `text-faint` é
reservado a informação realmente secundária ou decorativa.

## Tipografia

- display: identidade principal e uso raro;
- monospace: comandos, labels, status e metadata curta;
- sans-serif: parágrafos e conteúdo longo.

Não use monospace em textos extensos. Microtipografia não pode carregar conteúdo
essencial. Fontes externas ou locais exigem decisão posterior de assets,
licenciamento, preload e fallback.

## Espaçamento e layout

Use a escala de spacing existente. O visual retrô não justifica interfaces
apertadas. Comprimento de linha, agrupamento e espaço negativo devem sustentar a
hierarquia.

## Bordas, cantos e superfícies

- bordas normalmente têm 1px e baixo contraste;
- bordas brilhantes ficam reservadas a foco, seleção, sinal ou ação importante;
- cantos são quadrados, minimamente arredondados ou recortados;
- evitar radius amplo e aparência de SaaS;
- superfícies podem usar gradientes discretos, nunca glassmorphism generalizado.

## Glow e efeitos

Glow sinaliza atividade, foco ou seleção. Processor e indicadores podem usar
mais efeito que conteúdo normal. Grids, scanlines, estrelas e código decorativo
devem permanecer secundários e concentrados na homepage.

## Movimento

Movimento comunica atividade ou transição. Não criar delays fictícios. Toda
animação deve tolerar `prefers-reduced-motion`; funcionalidade não pode depender
dela.

## Breakpoints

Breakpoints respondem a falhas de layout, não a nomes de dispositivos. Mantenha
regras próximas do componente ou feature que alteram e use poucos pontos de
ruptura compartilhados quando houver repetição comprovada.
