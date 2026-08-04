# Checklist Frontend

## Antes de implementar

- [ ] Documentação relevante lida.
- [ ] Backend tratado como somente leitura.
- [ ] Contrato de rota e dados confirmado.
- [ ] Padrões existentes inspecionados.
- [ ] Segunda ocorrência validou qualquer novo componente.
- [ ] Mudanças do proprietário identificadas e preservadas.

## Estrutura

- [ ] `index.html` continua sendo o único documento completo.
- [ ] Feature retorna fragment sem shell duplicado.
- [ ] Feature não depende do CSS de outra feature.
- [ ] Componentes compartilhados não conhecem detalhes de feature.
- [ ] Markup simples não foi escondido em abstração desnecessária.

## Visual e conteúdo

- [ ] Identidade existente preservada.
- [ ] Um foco visual primário por view.
- [ ] Conteúdo tem prioridade sobre decoração.
- [ ] Português humano e inglês de sistema estão consistentes.
- [ ] Estados e ações usam componentes existentes.

## Acessibilidade

- [ ] Semântica e landmarks corretos.
- [ ] Hierarquia de headings lógica.
- [ ] Teclado e foco visível verificados.
- [ ] Nomes acessíveis e alt texts adequados.
- [ ] Contraste de conteúdo significativo suficiente.
- [ ] Estado não depende apenas de cor.
- [ ] Zoom e movimento reduzido verificados.

## Responsividade

- [ ] 320px, 360px e 390px.
- [ ] Tablet portrait e landscape.
- [ ] 1440px e wide desktop.
- [ ] Sem overflow não intencional.
- [ ] Navegação completa e descobrível.
- [ ] Conteúdo longo e ausente testados.

## HTMX e estados

- [ ] Idle, loading, success, empty, error e unavailable considerados.
- [ ] Loading não cria delay artificial.
- [ ] Error não expõe detalhes internos.
- [ ] Retry é seguro e possui contrato.
- [ ] History, focus, active state, title e deep links seguem o contrato aprovado.

## Entrega

- [ ] Paridade visual de áreas não redesenhadas confirmada.
- [ ] Arquivos Go/backend não foram alterados.
- [ ] Arquivos modificados e decisões relatados.
- [ ] Dependências de backend documentadas e não implementadas.
- [ ] Documentação atualizada quando o padrão mudou.
