# 04 — Backlog MVP em task cards para agentes

Cada task card deve ser executado isoladamente. Não misturar cards sem necessidade.

---

## Card 0.1 — Bootstrap do backend

**Objetivo:** criar base Go do backend.

**Entregas:**

- `/cmd/api/main.go`.
- Configuração `.env`.
- Logger estruturado.
- `/health`.
- Middleware request_id.
- Teste de healthcheck.

**Aceite:** API sobe localmente e responde status saudável.

---

## Card 0.2 — Docker Compose local

**Objetivo:** criar ambiente local reproduzível.

**Entregas:**

- PostgreSQL.
- Redis.
- MinIO/S3 local.
- Backend.
- Frontend.
- Volumes nomeados.

**Aceite:** `docker compose up` sobe todos os serviços.

---

## Card 0.3 — Migrations e sqlc

**Objetivo:** preparar banco e geração de queries.

**Entregas:**

- Configuração de migrations.
- `sqlc.yaml`.
- Migration inicial de extensões.
- Query de teste.

**Aceite:** migrations rodam e sqlc gera código.

---

## Card 1.1 — Tabela e CRUD de tenants

**Objetivo:** implementar tenants.

**Entregas:**

- Migration `tenants`.
- Queries sqlc.
- Repository.
- Service.
- Handler REST.
- Testes de service.

**Campos principais:**

```text
id, nome, nome_fantasia, tipo_documento, numero_documento, email, telefone, site, logo_url, status, tipo, configuracoes, created_at, updated_at, deleted_at
```

**Aceite:** CRUD funciona com status e tipo em pt_BR.

---

## Card 1.2 — Usuários e vínculo com tenant

**Objetivo:** permitir usuário em múltiplos tenants.

**Entregas:**

- Migration `users`.
- Migration `tenant_users`.
- Queries.
- Service.
- Handlers.
- Testes.

**Aceite:** mesmo usuário pode estar em mais de um tenant com cargo/status distintos.

---

## Card 1.3 — Autenticação JWT

**Objetivo:** login e refresh token.

**Entregas:**

- Hash de senha.
- Login.
- Refresh.
- Middleware auth.
- Contexto com user_id e tenant atual.

**Aceite:** endpoints protegidos exigem token válido.

---

## Card 2.1 — Roles e permissions

**Objetivo:** criar RBAC básico.

**Entregas:**

- Migrations `roles`, `permissions`, `role_permissions`.
- Seeds de permissões em pt_BR.
- Service de consulta de permissões.
- Testes.

**Aceite:** usuário recebe permissões via role no tenant.

---

## Card 2.2 — Relacionamento entre tenants

**Objetivo:** criar grafo de relacionamento entre tenants.

**Entregas:**

- Migration `tenant_relationships`.
- CRUD backend.
- Validações de status.
- Permissões por relacionamento.
- Testes.

**Aceite:** relacionamento ativo não concede acesso sem permissão explícita.

---

## Card 2.3 — Authorization service

**Objetivo:** centralizar autorização.

**Entregas:**

- Interface `Can`.
- Verificação tenant usuário.
- Verificação relacionamento.
- Verificação role/permission.
- Testes de acesso permitido/negado.

**Aceite:** handlers usam authorization service, não lógica duplicada.

---

## Card 3.1 — Catálogo de tipos de imóvel

**Objetivo:** criar catálogo interno pt_BR.

**Entregas:**

- Migration `property_types`.
- Seed inicial.
- Endpoint de listagem.
- Teste de seed/listagem.

**Aceite:** catálogo não armazena valores VRSync como domínio principal.

---

## Card 3.2 — Cadastro principal de imóveis

**Objetivo:** criar `properties`.

**Entregas:**

- Migration.
- Queries.
- Service.
- Handler.
- Testes.

**Aceite:** todo imóvel tem `tenant_dono_id` e filtros por tenant.

---

## Card 3.3 — Endereço de imóvel

**Objetivo:** criar endereço completo e regra de exibição.

**Entregas:**

- Migration `property_addresses`.
- Service de upsert.
- Validações de CEP/cidade/bairro para exportação futura.
- Testes.

**Aceite:** endereço completo é salvo internamente; exibição pública é controlada por `exibicao_endereco`.

---

## Card 3.4 — Tela de imóveis

**Objetivo:** criar listagem e formulário inicial.

**Entregas:**

- Página de listagem.
- Filtros básicos.
- Formulário criar/editar.
- Hooks com TanStack Query.
- Validação com React Hook Form + Zod.

**Aceite:** usuário cria e edita imóvel pelo frontend.

---

## Card 4.1 — Features e garantias

**Objetivo:** adicionar características e garantias.

**Entregas:**

- Migrations `features`, `property_features`, `property_warranties`.
- Seeds.
- Endpoints de associação.
- Testes.

**Aceite:** valores internos em pt_BR, sem valores VRSync nas entidades principais.

---

## Card 4.2 — Mídias de imóveis

**Objetivo:** upload e ordenação de imagens.

**Entregas:**

- Migration `property_media`.
- Upload para S3.
- Hash do arquivo.
- Ordenação.
- Marcação de principal.
- Testes.

**Aceite:** apenas uma mídia principal por imóvel e URL muda quando arquivo muda.

---

## Card 4.3 — Visibilidade de imóveis

**Objetivo:** permitir compartilhamento controlado.

**Entregas:**

- Migration `property_visibility`.
- Service de visibilidade.
- Integração com authorization service.
- Testes.

**Aceite:** tenant relacionado só acessa imóvel com visibilidade e permissão compatíveis.

---

## Card 5.1 — Integration mappings

**Objetivo:** isolar mapeamentos externos.

**Entregas:**

- Migration `integration_mappings`.
- Seeds VRSync.
- Repository.
- Testes.

**Aceite:** mapper consulta mapeamento sem contaminar domínio.

---

## Card 5.2 — VRSyncMapper

**Objetivo:** converter pt_BR para VRSync.

**Entregas:**

- Mapper de tipo de transação.
- Mapper de tipo de imóvel.
- Mapper de características.
- Mapper de garantias.
- Mapper de publicação.
- Testes de conversão.

**Aceite:** conversão é coberta por testes e falha para código sem mapeamento.

---

## Card 5.3 — VRSyncValidationService

**Objetivo:** validar imóvel antes de exportar.

**Entregas:**

- Validação de campos obrigatórios.
- Validação de mídia mínima.
- Validação de contato.
- Validação por tipo de transação.
- Endpoint `/validate-vrsync`.

**Aceite:** erros são apresentados antes da exportação.

---

## Card 5.4 — VRSyncExportService

**Objetivo:** gerar XML por tenant.

**Entregas:**

- Feed XML.
- Header por tenant.
- Listings.
- Snapshot do XML.
- Hash do XML.
- Endpoint/feed.

**Aceite:** XML válido pode ser reproduzido e auditado.

---

## Card 6.1 — Leads

**Objetivo:** cadastrar e listar leads.

**Entregas:**

- Migration `leads`.
- Endpoint público de criação.
- Endpoint autenticado de listagem.
- Normalização de contato.
- Deduplicação inicial.

**Aceite:** lead pertence a tenant e duplicidade é identificada.

---

## Card 6.2 — Atribuição manual de lead

**Objetivo:** atribuir lead a usuário ou tenant.

**Entregas:**

- Migration `lead_assignments`.
- Endpoint de atribuição.
- Histórico.
- Testes.

**Aceite:** atribuição respeita permissão e registra auditoria.

---

## Card 7.1 — Conversas e mensagens manuais

**Objetivo:** registrar comunicação.

**Entregas:**

- Migrations `conversations`, `messages`.
- Endpoints.
- UI básica.
- Testes.

**Aceite:** mensagens são vinculadas a conversa, lead, imóvel e tenant.

---

## Card 8.1 — Auditoria básica

**Objetivo:** consolidar logs auditáveis.

**Entregas:**

- Migration `audit_logs`.
- Audit service.
- Integração em tenants, imóveis, leads e VRSync.

**Aceite:** ações sensíveis aparecem na auditoria.

---

## Card 8.2 — Dashboard operacional

**Objetivo:** mostrar indicadores mínimos.

**Entregas:**

- Endpoints de métricas.
- Cards de dashboard.
- Filtros por período.

**Aceite:** métricas respeitam escopo do tenant.
