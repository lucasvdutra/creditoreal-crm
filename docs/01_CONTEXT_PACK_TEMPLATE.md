# 01 — Template de Context Pack para LLM / Antigravity

Use este modelo para cada tarefa. O objetivo é fornecer contexto suficiente sem sobrecarregar o agente.

---

## Context Pack — [NOME DA TAREFA]

### 1. Objetivo

Desenvolver **somente**:

```text
[descrever uma entrega pequena e verificável]
```

Exemplo:

```text
Implementar o CRUD backend de tenants, incluindo migration, queries sqlc, repository, service, handler REST, validações e testes de service.
```

### 2. Fora do escopo

Não implementar nesta tarefa:

```text
- frontend;
- relacionamento entre tenants;
- dashboard;
- VRSync;
- permissões avançadas;
- qualquer refatoração global.
```

### 3. Stack e padrões obrigatórios

```text
Backend: GoLang
API: REST
Banco: PostgreSQL
Queries: sqlc
Migrations: golang-migrate ou goose
Logs: zap ou zerolog
Testes: testing + testify
Valores internos: pt_BR, snake_case, sem acento
```

### 4. Invariantes do projeto

```text
- Todo dado de negócio pertence a um tenant ou tem regra explícita de escopo.
- Valores internos do domínio devem ser salvos em pt_BR.
- Valores externos VRSync nunca devem ser salvos nas entidades principais.
- Toda consulta sensível deve validar tenant_id e permissão efetiva.
- Ações sensíveis devem gerar auditoria.
- Endpoints devem retornar erros padronizados.
```

### 5. Arquivos relevantes

Ler antes de alterar:

```text
[lista de arquivos]
```

Criar ou modificar:

```text
[lista provável]
```

Não alterar:

```text
[arquivos críticos ou fora do escopo]
```

### 6. Modelo de dados esperado

```text
[campos da tabela, constraints, índices]
```

### 7. Endpoints esperados

```http
GET /api/[recurso]
POST /api/[recurso]
GET /api/[recurso]/{id}
PATCH /api/[recurso]/{id}
DELETE /api/[recurso]/{id}
```

### 8. Regras de validação

```text
- campo X obrigatório;
- status permitido: ativo, inativo, pendente, bloqueado, excluido;
- slug em português, snake_case, sem acento;
- não permitir duplicidade por tenant quando aplicável.
```

### 9. Testes obrigatórios

```text
- criar com sucesso;
- falhar com campo obrigatório ausente;
- listar filtrando por tenant;
- impedir acesso a tenant diferente;
- atualizar apenas campos permitidos;
- soft delete quando aplicável;
- erro padronizado em id inválido.
```

### 10. Critério de aceite

A tarefa está concluída quando:

```text
- migrations aplicam sem erro;
- queries sqlc geram código;
- testes focalizados passam;
- endpoints seguem contrato;
- nenhuma consulta ignora tenant_id;
- nenhum valor externo VRSync foi introduzido no domínio;
- o resumo final informa arquivos alterados e como testar.
```

---

## Prompt curto reutilizável

```text
Você é um agente de desenvolvimento Go/React/PostgreSQL neste projeto multi-tenant. Execute apenas a tarefa descrita no Context Pack. Não faça refatorações globais. Antes de codar, liste brevemente arquivos que serão alterados. Garanta tenant_id, permissões quando aplicável, valores internos em pt_BR e testes. Ao final, informe arquivos alterados, como testar, pendências e riscos.
```
