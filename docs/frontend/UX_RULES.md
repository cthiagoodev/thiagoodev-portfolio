# Regras de UX

## Modelo de experiência

O visitante navega por módulos de uma máquina de engenharia. A metáfora deve
melhorar continuidade e identidade sem exigir que o visitante a compreenda.

- Home → System / Profile
- About → Identity Module
- Experience → Career Log
- Projects → Project Registry
- Community → Network Node
- Contact → Communication Terminal

Usabilidade, conteúdo e clareza têm prioridade sobre teatralidade.

## Hierarquia

Cada view possui um foco primário. Em módulos: identidade do módulo, conteúdo,
ações importantes, metadata e por último decoração. Evite múltiplos elementos
competindo por atenção.

## Navegação

A navegação deve ser compreensível em português mesmo sem a metáfora de sistema,
funcionar com mouse, teclado e touch e futuramente suportar back, forward, active
state e deep links conforme `HTMX_CONTRACT.md`.

Active state não pode depender apenas de cor. Navegação móvel deve tornar todos
os destinos descobríveis.

## Requisições

Estados de loading são discretos, sem bloquear o shell e sem atrasos falsos.
Success rotineiro não exige celebração. Empty, error e unavailable sempre incluem
explicação humana; vocabulário de sistema pode complementar, nunca substituir.

## Ações

Inclua ação primária somente quando necessária. Ação primária usa sinal verde;
ações secundárias são mais silenciosas. Estados destrutivos nunca reutilizam o
verde operacional.

## Mobile

Mobile reorganiza prioridades, não apenas reduz o desktop. Conteúdo, navegação e
ações permanecem acessíveis; processor e diagnósticos decorativos podem ser
simplificados antes de informação significativa.

## Movimento

Concentre movimento ambiente no processor. Módulos de conteúdo permanecem mais
calmos. Respeite movimento reduzido.

## Idioma

Conteúdo humano usa português do Brasil. Inglês é reservado à interface da
máquina. A convenção completa está em `CONTENT_GUIDE.md`.

## Regra de decisão

Entre uma interação impressionante e uma alternativa clara, escolha a clara,
salvo quando o efeito comunicar estado, hierarquia ou identidade relevante.
