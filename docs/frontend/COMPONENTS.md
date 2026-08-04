# Componentes de Interface

## Regra de criação

Crie um componente quando ele tiver responsabilidade estável, uso repetido,
reduzir duplicação e continuar fácil de entender. A segunda ocorrência valida
novos componentes. Nem todo componente CSS precisa ser um Go template.

## Command Button

Classe base `.command-button`, com variantes `--primary` e `--secondary`.
Pode ser aplicada a `a` ou `button` conforme a semântica. Estados de hover e
focus devem ser equivalentes; o controle deve manter target confortável.

É somente CSS. Não criar template Go para botões.

## Module

`.module` é o container raiz visual de um fragment. Organiza superfície,
espaçamento e fluxo vertical sem impor a composição interna da feature.

## Module Header

`.module-header` identifica o módulo por label de sistema, título humano,
descrição e status opcional. Nesta fase é convenção HTML + CSS. Um Go template
só será criado quando dois fragments demonstrarem o mesmo contrato de dados.

## System Panel

`.system-panel` é uma superfície compartilhada para conteúdo agrupado. Variantes
devem surgir apenas por necessidade real; não é um card universal.

## System Label

`.system-label` apresenta metadata técnica curta em monospace. É visualmente
secundária e não pode carregar informação essencial em tamanho ou contraste
insuficientes.

## Status

`.status` combina indicador, texto e estado sem depender apenas de cor. Estados
iniciais: operational, warning, error e unavailable. Glow permanece sutil.

## Metadata

`.metadata` estiliza listas semânticas `dl`, com `dt` como label e `dd` como
valor. É convenção HTML + CSS; não criar partial para cada par.

## State Panel

`.state-panel` cobre empty, error e unavailable com label de sistema, heading,
explicação humana e ação opcional. Loading possui tratamento próprio para não
substituir conteúdo rápido desnecessariamente.

Nesta fase é convenção HTML + CSS. Um template Go depende de uso repetido e de
um contrato de dados claro.

## Loading Indicator

`.loading-state` comunica atividade de forma discreta, sem atraso artificial e
sem grande mudança de layout. Deve respeitar movimento reduzido.

## Específicos da homepage

Hero, atmosfera, processor, traces, pins, binary stream e animações associadas
são exclusivos de Home. Não devem ser promovidos a componentes globais.

## Não criar ainda

Não criar tag list, progress, section header, information group, card universal,
grid genérico, template de ícone ou template de link HTMX até existir uma
segunda ocorrência real que justifique a abstração.
