---
name: git
description: >-
  Git workflow and commit conventions for the portfolio monorepo.
  Defines commit message format, scopes, and permission boundaries between
  frontend and backend services.
---

# Git Skill & Monorepo Workflow

Esta skill define as diretrizes para operações de Git, padronização de mensagens de commit e limites de autonomia no monorepo.

## 1. Padrão de Mensagem de Commit

O formato obrigatório para mensagens de commit é:

```text
[escopo] Ação executada
```

### Regras de Formatação:
- **Linha única:** Mensagem direta, concisa e em uma única linha (sem corpos extensos).
- **Escopos do monorepo:**
  - `[portfolio-web]`: Alterações em `apps/web`.
  - `[app-api]`: Alterações em `services/app-api`.
  - `[repo]`: Alterações na raiz do monorepo, orquestração global (`docker-compose.yml`) ou documentações (`docs/`).
- **Exemplos:**
  - `[portfolio-web] Cria arquivo Dockerfile`
  - `[portfolio-web] Configura porta do preview`
  - `[app-api] Cria entidade Project e testes de unidade`
  - `[repo] Atualiza documentação de arquitetura`

## 2. Limites de Permissão e Autonomia

### Frontend (`apps/web`)
- O agente tem **autonomia para commitar e dar push** para a branch principal (`main`), mantendo o ciclo de entrega do frontend ágil conforme solicitado.
- Seguir estritamente o padrão `[portfolio-web] Ação executada`.

### Serviços Backend (`services/`)
- O agente **NUNCA deve commitar, dar push ou executar comandos de mutação git de forma autônoma** nesta área.
- Qualquer comando git em `services/` só pode ser executado mediante:
  1. **Pedido explícito** do proprietário do repositório.
  2. **Revisão prévia** das alterações pelo usuário.
