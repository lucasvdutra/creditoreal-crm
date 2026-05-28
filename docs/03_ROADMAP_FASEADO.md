# 03 — Roadmap faseado para desenvolvimento incremental

## Visão geral

O projeto deve ser executado em fases curtas, cada uma com entregas testáveis. A ordem abaixo reduz risco arquitetural e evita que o LLM tente implementar tudo ao mesmo tempo.

---

# Fase 0 — Fundação técnica

## Objetivo

Criar a base do repositório e os padrões técnicos.

## Entregas

- Monorepo ou estrutura separada `backend/` e `frontend/`.
- Docker Compose com PostgreSQL, Redis e MinIO/S3 local.
- Backend Go com healthcheck.
- Frontend React + TypeScript com layout base.
- Migrations configuradas.
- sqlc configurado.
- Logger estruturado.
- Configuração por `.env`.
- OpenAPI inicial.

## Critério de aceite

- `docker compose up` sobe serviços.
- API responde `/health`.
- Frontend abre tela inicial.
- Migrations rodam.
- Teste mínimo do backend passa.

---

# Fase 1 — Autenticação, tenants e usuários

## Objetivo

Permitir que tenants e usuários existam com escopo inicial de acesso.

## Entregas

- Tabela `tenants`.
- Tabela `users`.
- Tabela `tenant_users`.
- Autenticação JWT + refresh token.
- CRUD de tenants.
- Convite/criação de usuários.
- Associação de usuário a tenant.
- Status em pt_BR.
- Auditoria básica de criação/alteração.

## Não incluir

- Relacionamento entre tenants.
- VRSync.
- Imóveis.
- Dashboard.

## Critério de aceite

- Um usuário pode pertencer a múltiplos tenants.
- Listagens filtram por tenant.
- Status internos usam pt_BR.
- Endpoints de tenant têm testes.

---

# Fase 2 — RBAC e relacionamento entre tenants

## Objetivo

Implementar papéis, permissões e relação flexível entre tenants.

## Entregas

- Tabela `roles`.
- Tabela `permissions`.
- Tabela `role_permissions`.
- Tabela `tenant_relationships`.
- Serviço de autorização.
- Permissões efetivas por usuário, tenant e relacionamento.
- Endpoints de relacionamento.
- Testes de acesso cruzado entre tenants.

## Critério de aceite

- O sistema não infere acesso apenas por relacionamento.
- Acesso considera tenant, status, papel, permissão, relação e escopo.
- Testes cobrem permissão permitida e negada.

---

# Fase 3 — Cadastro de imóveis núcleo

## Objetivo

Criar o cadastro principal de imóveis com campos essenciais para operação e VRSync futuro.

## Entregas

- Tabela `properties`.
- Tabela `property_addresses`.
- Tabela opcional `property_prices` ou campos financeiros em `properties`.
- Tabela `property_types`.
- CRUD backend de imóveis.
- Filtros básicos por status, tipo, transação e tenant.
- Validações obrigatórias por tipo de transação.
- Tela React de listagem e formulário básico.

## Critério de aceite

- Imóvel sempre tem `tenant_dono_id`.
- Valores internos usam pt_BR.
- Listagens têm paginação.
- Queries filtram por tenant e autorização.
- Tela permite criar e editar imóvel sem mídias ainda.

---

# Fase 4 — Características, garantias, mídias e visibilidade

## Objetivo

Completar o cadastro operacional do imóvel.

## Entregas

- Tabela `features`.
- Tabela `property_features`.
- Tabela `property_warranties`.
- Tabela `property_media`.
- Tabela `property_visibility`.
- Upload para S3/MinIO.
- Hash de arquivo.
- Ordenação de imagens.
- Uma imagem principal.
- Regras de visibilidade por tenant.
- Seeds de catálogos.

## Critério de aceite

- Imóvel publicado deve ter ao menos uma imagem.
- Apenas uma imagem principal por imóvel.
- Visibilidade respeita relacionamento e permissão.
- Catálogos internos permanecem em pt_BR.

---

# Fase 5 — Validação e exportação VRSync

## Objetivo

Implementar VRSync como camada de integração, sem contaminar o domínio interno.

## Entregas

- Tabela `integration_mappings`.
- Tabela `vrsync_feed_configs`.
- Tabela `tenant_contact_profiles`.
- Tabela `property_publications`.
- `VRSyncMapper`.
- `VRSyncValidationService`.
- `VRSyncExportService`.
- Endpoint de validação.
- Endpoint/feed XML.
- Snapshot de XML gerado.
- Tela de erros de integração.

## Critério de aceite

- Feed XML válido por tenant.
- Mapper converte pt_BR para VRSync.
- Erros aparecem antes da exportação.
- Nenhum valor VRSync fica salvo nas entidades principais.
- XML gerado pode ser reproduzido por snapshot/hash.

---

# Fase 6 — Leads e atribuição manual

## Objetivo

Captar leads, deduplicar e atribuir manualmente.

## Entregas

- Tabela `leads`.
- Tabela `lead_assignments`.
- Captação por formulário público.
- Normalização de telefone/e-mail.
- Deduplicação inicial.
- Atribuição manual a usuário ou tenant.
- Listagem e detalhe de leads.
- Histórico básico.

## Critério de aceite

- Lead pertence a um tenant.
- Lead pode estar associado a imóvel.
- Duplicidade é detectada por telefone/e-mail/imóvel/janela temporal.
- Atribuição registra histórico.

---

# Fase 7 — Comunicação e CRM básico

## Objetivo

Registrar conversas, mensagens e evolução comercial básica.

## Entregas

- Tabela `conversations`.
- Tabela `messages`.
- Histórico manual de comunicação.
- Tarefas comerciais.
- Visitas.
- Propostas.
- Pipeline básico.
- Status e temperatura do lead.

## Critério de aceite

- Conversas vinculam tenant, lead, imóvel e canal.
- Mensagens são append-only sempre que possível.
- Lead tem histórico comercial consultável.

---

# Fase 8 — Dashboard e auditoria ampliada

## Objetivo

Gerar indicadores operacionais mínimos e ampliar rastreabilidade.

## Entregas

- Tabela `audit_logs` consolidada.
- Indicadores de imóveis.
- Indicadores de leads.
- Indicadores de VRSync.
- Métricas por tenant, canal e período.
- Views materializadas se necessário.

## Critério de aceite

- Dashboard respeita escopo de tenant.
- Ações sensíveis aparecem em auditoria.
- Métricas básicas são reproduzíveis.

---

# Fase 9 — Escala e otimização

## Objetivo

Melhorar performance com base em métricas reais.

## Entregas

- Jobs assíncronos com Redis/Asynq, NATS ou RabbitMQ.
- Particionamento de auditoria/mensagens se necessário.
- OpenSearch se a busca imobiliária exigir.
- Cache de consultas frequentes.
- Read replicas se necessário.
- Observabilidade.

## Critério de aceite

- Gargalos medidos antes da otimização.
- Otimizações documentadas por decisão técnica.
- Sem introduzir complexidade sem métrica justificadora.
