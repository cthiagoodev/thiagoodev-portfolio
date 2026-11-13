# ⚡ Thiago Sousa — Portfolio & Engineering Platform

[![Java Version](https://img.shields.io/badge/Java-25-ED8B00?style=flat&logo=openjdk&logoColor=white)](https://openjdk.org/)
[![Spring Boot](https://img.shields.io/badge/Spring_Boot-4.1-6DB33F?style=flat&logo=springboot&logoColor=white)](https://spring.io/projects/spring-boot)
[![Astro Version](https://img.shields.io/badge/Astro-5.x-BC52EE?style=flat&logo=astro&logoColor=white)](https://astro.build/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178C6?style=flat&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18.6-4169E1?style=flat&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Este repositório é um **monorepo de engenharia de software** que abriga meu portfólio profissional e uma plataforma de serviços backend em **Java 25 e Spring Boot 4**. Projetado do zero sem depender de soluções de BaaS (Backend-as-a-Service) fechadas, o projeto atua como laboratório prático para arquitetura limpa pragmática em camadas, **Test-Driven Development (TDD)** estrito e infraestrutura própria.

---

## 🏛️ Visão Geral da Arquitetura

O ecossistema é organizado em um monorepo limpo e modular:

```text
portfolio/
├── apps/
│   └── web/                                        # Frontend Astro 5 + React + TypeScript + Tailwind
├── services/
│   └── api/                                        # Backend em Java 25 + Spring Boot 4 (Gradle Kotlin DSL)
│       └── src/main/java/br/com/thiagoodev/api/
│           ├── domain/                             # Entidades de domínio/JPA, contratos (repositories/usecases) e exceções
│           ├── application/                        # Implementações de casos de uso (@Service) e DTOs (Records)
│           ├── infrastructure/                     # Integrações externas (GitHub RestClient) e configurações Spring
│           └── presentation/                       # Controllers REST (@RestController) e GlobalExceptionHandler
└── .agents/skills/                                 # Skills do agente, arquitetura e styleguide do repositório
```

---

## 🛠️ Stack Tecnológica

### Frontend (`apps/web`)
- **[Astro 5](https://astro.build/)**: Framework web focado em conteúdo e performance com arquitetura de ilhas (*Islands Architecture*).
- **[React](https://react.dev/) & [TypeScript](https://www.typescriptlang.org/)**: Componentes interativos de alta fidelidade e tipagem estrita de ponta a ponta.
- **[Tailwind CSS](https://tailwindcss.com/) & [Magic UI](https://magicui.design/)**: Design refinado com animações suaves (`BlurFade`, `FlickeringGrid`, `Dock`).

### Backend (`services/api`)
- **[Java 25](https://openjdk.org/)**: Linguagem moderna com suporte a Records, Pattern Matching, Virtual Threads e tipagem forte.
- **[Spring Boot 4.1](https://spring.io/projects/spring-boot)**: Ecossistema completo com `spring-boot-starter-webmvc`, injeção de dependências nativa e configuração modular.
- **[Spring Data JPA](https://spring.io/projects/spring-data-jpa) & Hibernate**: Persistência relacional produtiva com validação estrita de schema.
- **[JUnit 5](https://junit.org/junit5/) & [Mockito](https://site.mockito.org/)**: Suíte de testes unitários rápidos, isolados e determinísticos.

### Persistência & Dados
- **[PostgreSQL 18.6](https://www.postgresql.org/)**: Banco de dados relacional com suporte a arrays nativos e constraints de integridade referencial.
- **[Flyway](https://flywaydb.org/)**: Controle de versão automatizado do schema de banco via migrations SQL (`V1__*.sql`).

---

## 💎 Diferenciais e Padrões de Engenharia

### 1. Arquitetura em Camadas Pragmática (Spring Clean Architecture)
As dependências apontam estritamente em direção ao domínio:
$$\text{Presentation (Controllers)} \longrightarrow \text{Domain Contracts} \longleftarrow \text{Application (UseCases)}$$
- **Sem Reinventar a Roda**: Aproveita nativamente o container IoC do Spring, o roteamento declarativo do WebMVC, o pool HikariCP, o `@RestControllerAdvice` e os repositórios dinâmicos do Spring Data JPA.
- **Entidades Ricas com Invariantes**: As entidades em `domain.entities` combinam o mapeamento JPA com construtores de domínio validados (`validate()`), inicialização defensiva de coleções e métodos de negócio.

### 2. Test-Driven Development (TDD) Estrito
Todo o desenvolvimento de entidades, casos de uso e controllers segue o ciclo **Red-Green-Refactor**:
1. **Red**: Criação prévia do teste unitário (JUnit 5 / Mockito / `@WebMvcTest`) cobrindo cenários de sucesso e falha.
2. **Green**: Implementação mínima e suficiente para aprovação da suíte.
3. **Refactor**: Otimização contínua de legibilidade e design mantendo os testes verdes.
- **Regra de Zero Comentários em Testes**: Os arquivos `*Test.java` não contêm comentários. A estrutura AAA (Arrange, Act, Assert) é delimitada estritamente por blocos isolados com linhas em branco.

### 3. Styleguide Java 25 & Spring Boot 4
- **Injeção Exclusiva por Construtor**: Atributos `private final` sem `@Autowired` em campos, permitindo instanciamento direto em testes unitários sem subir o container do Spring.
- **Nomes Completos e Sem Abreviações**: Proibição de nomes crípticos no código (`aboutRepository` em vez de `aboutRepo`, `workExperience` em vez de `workExp`).
- **Tratamento Global de Exceções**: Mapeamento limpo via `@RestControllerAdvice` (`ResourceNotFoundException` → 404, `DomainValidationException` → 400).

---

## 🚀 Como Executar o Projeto Localmente

### Pré-requisitos
- [JDK 25](https://openjdk.org/)
- [Node.js 20+](https://nodejs.org/) & [pnpm](https://pnpm.io/)
- [Docker & Docker Compose](https://www.docker.com/)

### 1. Clonar o Repositório
```bash
git clone git@github.com:cthiagoodev/thiagoodev-portfolio.git
cd thiagoodev-portfolio
```

### 2. Subir o Banco de Dados (PostgreSQL 18)
```bash
docker compose up -d
```

### 3. Executar o Backend (`services/api`)
```bash
cd services/api
./gradlew bootRun
```
Para rodar a suíte de testes unitários do backend:
```bash
./gradlew test
```

### 4. Executar o Frontend (`apps/web`)
```bash
cd apps/web
pnpm install
pnpm dev
```
Acesse a aplicação em `http://localhost:4321`.

---

## 📦 Padrão de Commits

O repositório segue estritamente mensagens de commit atômicas e de linha única:
```text
[escopo] Ação executada
```
- `[portfolio-web]`: Modificações na interface e componentes Astro/React (`apps/web`).
- `[api]`: Modificações no serviço backend em Java/Spring Boot (`services/api`).
- `[repo]`: Modificações de orquestração global, documentação ou configurações raiz.

---

## 👨‍💻 Autor

Desenvolvido por **Thiago Sousa**  
- **LinkedIn**: [linkedin.com/in/thiagoodev](https://www.linkedin.com/in/thiagoodev/)  
- **GitHub**: [github.com/cthiagoodev](https://github.com/cthiagoodev)