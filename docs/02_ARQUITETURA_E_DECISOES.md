# 02 — Arquitetura e decisões técnicas obrigatórias

## 1. Arquitetura geral

A plataforma deve seguir arquitetura modular, API-first e orientada a domínio.

Estrutura sugerida:

```text
/cmd/api
/internal/auth
/internal/tenant
/internal/relationship
/internal/user
/internal/rbac
/internal/property
/internal/media
/internal/publication
/internal/vrsync
/internal/lead
/internal/conversation
/internal/crm
/internal/dashboard
/internal/audit
/internal/integration
/pkg/database
/pkg/http
/pkg/queue
/pkg/storage
/pkg/logger
```

## 2. Decisões obrigatórias

### ADR-001 — PostgreSQL como fonte de verdade

O domínio é relacional, transacional e multi-tenant. Portanto, PostgreSQL é a fonte de verdade para tenants, usuários, permissões, imóveis, leads, publicações, conversas, auditoria e integração.

Não usar NoSQL como banco principal no MVP.

### ADR-002 — Redis como suporte, não fonte de verdade

Redis pode ser usado para:

- cache;
- rate limit;
- locks distribuídos;
- filas leves;
- sessões efêmeras.

Redis não deve armazenar dados críticos como fonte primária.

### ADR-003 — S3 compatível para arquivos

Imagens, vídeos, plantas e documentos devem ser armazenados em S3 compatível. O banco deve guardar metadados, URLs, hash e storage_key.

### ADR-004 — OpenSearch somente em fase de escala

No MVP, busca deve ser feita com PostgreSQL. OpenSearch deve entrar apenas quando houver necessidade real de busca textual, geográfica, ranqueamento e filtros facetados em escala.

### ADR-005 — Domínio interno em pt_BR

Todos os valores internos de negócio devem ser salvos em português brasileiro, usando slugs técnicos:

```text
minúsculos
sem acento
snake_case
```

Exemplos:

```text
ativo
venda_aluguel
apartamento
casa_de_condominio
portaria_24h
seguro_fianca
```

### ADR-006 — VRSync como camada de integração

VRSync não é modelo de domínio. Valores como `For Sale`, `Residential / Apartment`, `Pool`, `STANDARD` e `SECURITY_DEPOSIT` só podem existir em:

```text
VRSyncMapper
VRSyncValidationService
VRSyncExportService
VRSyncImportService
integration_mappings
snapshots de payload/XML
logs de integração
```

Nunca gravar valores VRSync diretamente em `properties`, `features`, `property_types`, `leads` ou entidades principais.

### ADR-007 — Multi-tenancy por relações configuráveis

Não existe subtenant fixo. Todo participante é tenant autônomo.

Acesso depende de:

```text
tenant do usuário
tenant dono do recurso
relacionamento ativo
permissões do relacionamento
papel do usuário
permissões efetivas
visibilidade do recurso
```

### ADR-008 — Autorização no backend

O frontend pode esconder botões, mas a autorização real deve estar no backend.

Toda operação sensível deve passar por um serviço de autorização, por exemplo:

```go
authorization.Can(ctx, userID, tenantID, resource, action)
```

### ADR-009 — Auditoria append-only

Ações sensíveis devem registrar auditoria. A tabela de auditoria não deve ser alterada retroativamente por fluxos comuns.

### ADR-010 — sqlc recomendado

Usar sqlc para reduzir overhead de ORM, melhorar previsibilidade de queries e facilitar otimização. Queries devem ser explícitas e filtradas por tenant sempre que necessário.

## 3. Camadas recomendadas no backend

```text
Handler REST
  ↓
DTO / validação de entrada
  ↓
Application service / use case
  ↓
Authorization service
  ↓
Repository / sqlc queries
  ↓
PostgreSQL
```

## 4. Regras de dependência

- `internal/property` pode depender de `internal/tenant`, `internal/rbac`, `internal/audit` e `pkg/database`.
- `internal/vrsync` pode depender de `internal/property`, `internal/media`, `internal/publication` e `internal/integration`.
- Domínio não deve depender de handlers HTTP.
- Services não devem conhecer componentes React.
- Frontend deve depender do contrato de API, não do banco.

## 5. Estratégia incremental

Cada módulo deve nascer com:

```text
migration
queries sqlc
repository
service
handler REST
DTOs
testes mínimos
OpenAPI parcial
```

Só depois criar telas React.

## 6. Estrutura de erro padronizada

Formato sugerido:

```json
{
  "error": {
    "code": "tenant.nao_encontrado",
    "message": "Tenant não encontrado.",
    "details": {}
  }
}
```

## 7. Performance desde o início

Obrigatório:

- índice por `tenant_id` nas tabelas com dados por tenant;
- índice composto para filtros frequentes;
- paginação em listagens;
- evitar `SELECT *` em listagens densas;
- usar transactions para operações compostas;
- evitar N+1 no backend;
- usar TanStack Query no frontend para cache de leitura;
- usar jobs assíncronos para mídia, VRSync, notificações e dashboards.
