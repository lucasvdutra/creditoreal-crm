-- name: UserHasTenantPermission :one
SELECT EXISTS (
    SELECT 1
    FROM tenant_users tu
    JOIN roles r ON r.id = tu.role_id
    JOIN role_permissions rp ON rp.role_id = r.id
    JOIN permissions p ON p.id = rp.permission_id
    WHERE tu.user_id = $1
      AND tu.tenant_id = $2
      AND p.codigo = $3
      AND tu.status = 'ativo'
      AND tu.deleted_at IS NULL
      AND r.status = 'ativo'
      AND r.deleted_at IS NULL
)::bool;

-- name: RelationshipAllowsPermission :one
SELECT EXISTS (
    SELECT 1
    FROM tenant_relationships tr
    WHERE tr.tenant_origem_id = $1
      AND tr.tenant_destino_id = $2
      AND tr.status = 'ativo'
      AND tr.deleted_at IS NULL
      AND $3::text = ANY(tr.permissoes)
)::bool;
