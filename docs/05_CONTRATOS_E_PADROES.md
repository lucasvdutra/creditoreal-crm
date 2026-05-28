# 05 — Contratos e padrões de desenvolvimento

## 1. Convenção de idioma

Todos os valores internos de domínio devem estar em pt_BR, com slugs técnicos:

```text
snake_case
minúsculos
sem acento
sem espaços
```

Exemplos corretos:

```text
ativo
em_revisao
venda_aluguel
casa_de_condominio
seguro_fianca
publicacao_compartilhada
```

Exemplos incorretos:

```text
Active
For Sale
Residential / Apartment
Pool
SECURITY_DEPOSIT
Venda e Aluguel
```

## 2. Padrão de API REST

### Listagem

```http
GET /api/properties?page=1&page_size=20&status=ativo
```

Resposta:

```json
{
  "data": [],
  "pagination": {
    "page": 1,
    "page_size": 20,
    "total": 0
  }
}
```

### Criação

```http
POST /api/properties
```

Resposta esperada:

```http
201 Created
```

### Atualização parcial

```http
PATCH /api/properties/{id}
```

Não usar PUT para atualização parcial.

### Exclusão

Preferir soft delete quando a entidade tiver histórico relevante.

```http
DELETE /api/properties/{id}
```

## 3. Padrão de erro

```json
{
  "error": {
    "code": "imovel.nao_encontrado",
    "message": "Imóvel não encontrado.",
    "details": {}
  }
}
```

Códigos devem usar pt_BR técnico:

```text
tenant.nao_encontrado
permissao.negada
imovel.sem_imagem
vrsync.mapeamento_inexistente
lead.duplicado
```

## 4. Padrão de paginação

Toda listagem deve ter paginação.

Parâmetros:

```text
page
page_size
sort
order
```

Limite:

```text
page_size máximo: 100
padrão: 20
```

## 5. Padrão de filtros

Filtros devem ser explícitos e indexáveis.

Exemplo para imóveis:

```text
tenant_id
status
tipo_imovel
tipo_transacao
cidade
bairro
preco_min
preco_max
quartos
vagas_garagem
```

## 6. Padrão de banco

### IDs

Usar UUID ou ULID para identificadores expostos.

### Campos temporais

```text
created_at
updated_at
deleted_at
```

### Multi-tenancy

Tabelas de dados por tenant devem ter:

```text
tenant_id ou tenant_dono_id
índice por tenant_id
índices compostos com filtros frequentes
```

### Soft delete

Quando houver `deleted_at`, queries padrão devem excluir registros removidos.

## 7. Padrão de migrations

Cada migration deve:

- ter nome claro;
- criar constraints;
- criar índices necessários;
- evitar drop destrutivo no MVP;
- ter rollback quando a ferramenta permitir;
- ser pequena o suficiente para revisão.

Exemplo de nome:

```text
000004_create_properties.sql
```

## 8. Padrão sqlc

Queries devem ser nomeadas por caso de uso:

```sql
-- name: CreateTenant :one
-- name: GetTenantByID :one
-- name: ListTenants :many
-- name: UpdateTenant :one
-- name: SoftDeleteTenant :exec
```

Evitar queries genéricas demais.

## 9. Padrão Go

### Organização por módulo

```text
/internal/property/handler.go
/internal/property/service.go
/internal/property/repository.go
/internal/property/dto.go
/internal/property/validator.go
/internal/property/errors.go
```

### Contexto

Toda chamada de banco deve receber `context.Context`.

### Erros

Não retornar erros crus de banco diretamente ao cliente.

### Transactions

Operações compostas devem usar transaction explícita.

## 10. Padrão React

### Componentes

- componentes funcionais;
- TypeScript estrito;
- props tipadas;
- formulários com React Hook Form + Zod;
- dados remotos com TanStack Query;
- evitar lógica de negócio pesada no componente.

### Estrutura sugerida

```text
src/features/properties/
├── api.ts
├── hooks.ts
├── schemas.ts
├── types.ts
├── pages/
│   ├── PropertyListPage.tsx
│   └── PropertyFormPage.tsx
└── components/
    ├── PropertyForm.tsx
    └── PropertyTable.tsx
```

## 11. Padrão de OpenAPI

A API deve ser documentada progressivamente. Cada módulo novo deve atualizar o contrato.

Obrigatório documentar:

- request;
- response;
- códigos de erro;
- autenticação;
- paginação;
- filtros.

## 12. Padrão de logs

Logs devem ser estruturados.

Campos recomendados:

```text
request_id
user_id
tenant_id
resource
action
status_code
duration_ms
error_code
```

Nunca logar senhas, tokens, documentos sensíveis ou payloads pessoais completos.
