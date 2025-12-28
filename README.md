# 🚀 Portfólio de Projetos - Foco em Backend Java e Infraestrutura Self-Hosted

![Java](https://img.shields.io/badge/Java-25-ED8B00?style=for-the-badge&logo=java&logoColor=white)
![Spring Boot](https://img.shields.io/badge/Spring_Boot-3.x-6DB33F?style=for-the-badge&logo=spring&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Container-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Status](https://img.shields.io/badge/Status-Em_Produção-success?style=for-the-badge)

Este projeto atua como meu **laboratório prático** e **vitrine profissional**, com o objetivo principal de consolidar o conhecimento em **Desenvolvimento Backend com Java** e **Infraestrutura Própria (Self-Hosted)**.

Ele foi desenvolvido do zero para demonstrar a capacidade de gerenciar o ciclo completo de engenharia de software: do código à produção, incluindo a configuração de servidores, CI/CD e integração com APIs externas.

---

### 🎯 Funcionalidade Central: Sincronização Dinâmica com GitHub

Um dos diferenciais técnicos é a **sincronização automática** para manter a seção de projetos sempre atualizada, espelhando meus repositórios públicos do GitHub sem necessidade de *input* manual.

* **Integração:** Utiliza a **API Pública do GitHub** via `RestClient`.
* **Scheduler Inteligente:** Um `Schedule` no **Spring Boot** executa periodicamente a busca por atualizações.
* **Lógica de Filtragem:** O sistema busca todos os repositórios públicos, mas **filtra e exibe apenas** aqueles que possuem a *topic* (tag) `portfolio`. Isso permite controlar o que vai para a produção diretamente pelo painel do GitHub.

---

### 🛠️ Stack Tecnológica & Arquitetura

A arquitetura atual é **monolítica**, escolhida estrategicamente para reduzir a complexidade operacional nesta fase e focar na robustez do código e da infraestrutura.

#### 💻 Backend & Frontend

| Categoria | Tecnologia | Detalhe |
| :--- | :--- | :--- |
| **Backend** | **Java 25** | Utilização de recursos modernos da linguagem para alta performance. |
| **Framework** | **Spring Boot 3** | Base para criação de serviços RESTful, Agendamento de Tarefas e Injeção de Dependências. |
| **Template Engine** | **Thymeleaf** | Escolhido para **Server-Side Rendering (SSR)**, garantindo SEO amigável e renderização rápida. |
| **Interatividade** | **HTMX** | Adiciona dinamismo e atualizações parciais de página (AJAX) sem a complexidade de um SPA, mantendo o estado no servidor. |
| **Estilização** | **CSS3 / Tailwind** | Design responsivo com foco em performance de renderização (LCP/CLS otimizados). |

#### 💾 Dados & Armazenamento (Self-Hosted)

| Serviço | Tecnologia | Modo de Uso em Produção |
| :--- | :--- | :--- |
| **Banco de Dados** | **PostgreSQL** | Rodando em container Docker com **volumes persistentes** para durabilidade dos dados. |
| **Object Storage** | **MinIO** | Solução **S3-Compatible** self-hosted para armazenamento de imagens de perfil e assets dos projetos. |

---

### 🚀 Engenharia de Infraestrutura e Deployment

Este é o ponto focal de estudo do projeto: fugir das soluções "Serverless/PaaS" prontas e gerenciar a infraestrutura "no ferro" (VPS).

| Componente | Tecnologia | Propósito |
| :--- | :--- | :--- |
| **Servidor** | **Contabo VM** | Máquina virtual Linux (Ubuntu/Debian) dedicada ao hosting de Produção. |
| **Containerização** | **Docker & Compose** | Empacota a aplicação Spring Boot (Layered JAR) e serviços de suporte para isolamento total. |
| **Web Server** | **Nginx** | Atua como **Proxy Reverso**, gerenciando SSL (Certbot), compressão Gzip e roteamento de tráfego. |
| **CI/CD** | **GitHub Actions** | Pipeline automatizado que realiza build, testes e deploy contínuo na VPS via SSH. |

---

### 🗺️ Visão de Futuro e Próximos Passos

O roadmap de evolução visa aprimorar a arquitetura para cenários de maior escala:

1.  **Refatoração para Microserviços:** Migrar módulos específicos para serviços isolados para praticar comunicação assíncrona.
2.  **Otimização de Performance (Caching):** Implementação de **Redis** para cachear as respostas da API do GitHub e reduzir a latência.
3.  **Painel Administrativo:** Construção de um CMS simples para gestão de conteúdos estáticos do blog.

---
*Developed by [Thiago Sousa](https://www.linkedin.com/in/thiagoodev/)*