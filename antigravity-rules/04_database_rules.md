# Database Rules — PostgreSQL/sqlc

## Regras obrigatórias

1. PostgreSQL é a fonte de verdade.
2. Dados por tenant exigem `tenant_id` ou `tenant_dono_id`.
3. Criar índices por tenant e filtros frequentes.
4. Usar constraints para integridade.
5. Usar slugs pt_BR para valores internos.
6. Não salvar valores externos VRSync no domínio.
7. Migrations devem ser pequenas e revisáveis.
8. Evitar alterações destrutivas.

## Campos comuns

```text
id
created_at
updated_at
deleted_at quando aplicável
```

## Índices mínimos

Para tabelas por tenant:

```sql
CREATE INDEX ... ON tabela (tenant_id);
```

Para imóveis:

```sql
CREATE INDEX ... ON properties (tenant_dono_id, status);
CREATE INDEX ... ON properties (tenant_dono_id, tipo_transacao);
CREATE INDEX ... ON properties (tenant_dono_id, tipo_imovel);
```

## sqlc

Nomear queries por caso de uso:

```sql
-- name: ListPropertiesByTenant :many
```

Não usar query genérica quando o caso exigir autorização específica.
