-- name: CreateTenantRelationship :one
INSERT INTO tenant_relationships (
    tenant_origem_id,
    tenant_destino_id,
    tipo,
    status,
    permissoes,
    observacao
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, tenant_origem_id, tenant_destino_id, tipo, status, permissoes, observacao, created_at, updated_at, deleted_at;

-- name: GetTenantRelationshipByID :one
SELECT id, tenant_origem_id, tenant_destino_id, tipo, status, permissoes, observacao, created_at, updated_at, deleted_at
FROM tenant_relationships
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListTenantRelationships :many
SELECT id, tenant_origem_id, tenant_destino_id, tipo, status, permissoes, observacao, created_at, updated_at, deleted_at
FROM tenant_relationships
WHERE deleted_at IS NULL
  AND (tenant_origem_id = $1 OR tenant_destino_id = $1)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateTenantRelationship :one
UPDATE tenant_relationships
SET
    status = COALESCE(sqlc.narg('status'), status),
    permissoes = COALESCE(sqlc.narg('permissoes'), permissoes),
    observacao = COALESCE(sqlc.narg('observacao'), observacao),
    updated_at = now()
WHERE id = sqlc.arg('id') AND deleted_at IS NULL
RETURNING id, tenant_origem_id, tenant_destino_id, tipo, status, permissoes, observacao, created_at, updated_at, deleted_at;

-- name: SoftDeleteTenantRelationship :exec
UPDATE tenant_relationships
SET deleted_at = now(), status = 'excluido', updated_at = now()
WHERE id = $1 AND deleted_at IS NULL;
