# Guia de Conteúdo

## Regra de idioma

Português do Brasil representa a camada humana. Inglês representa a interface da
máquina. A alternância deve ser intencional e previsível.

## Camada humana — português

Use português em navegação, headings, descrições, biografia, experiência,
projetos, comunidade, contato, mensagens de erro e instruções.

Exemplos: `Sobre`, `Experiência`, `Projetos`, `Não foi possível carregar esta
seção`, `Tentar novamente`.

## Camada de sistema — inglês

Use inglês em labels técnicas curtas, identificação de módulos, revisão, node,
status e diagnósticos.

Exemplos: `MODULE // 01`, `SYSTEM ONLINE`, `PROCESSOR UNIT`, `STATUS`, `REV. 02`.

Labels de sistema nunca substituem explicação humana necessária.

## Tom

- técnico, direto, maduro e preciso;
- profissional antes de teatral;
- primeira pessoa apenas quando representar Thiago;
- evitar jargão hacker, mensagens alarmistas e lore excessiva;
- evitar slogans genéricos e superlativos sem evidência.

## Navegação e ações

Rótulos devem descrever destino ou resultado. A metáfora pode complementar, mas
não esconder a função. Use verbos claros: `Explorar projetos`, `Abrir perfil`,
`Entrar em contato`, `Tentar novamente`.

## Estados

- loading: label curta de sistema;
- empty: explicação humana do conteúdo ausente;
- error: o que falhou e possível próximo passo;
- unavailable: diferença clara entre ausência intencional e falha.

Não exponha stack traces, códigos internos ou mensagens do backend.

## Consistência

- use `Thiago Sousa` como nome público principal;
- preserve nomes próprios de tecnologias;
- mantenha capitalização normal no conteúdo humano;
- uppercase é reservado à camada visual da máquina;
- datas, localidades e períodos devem seguir convenção brasileira consistente.
