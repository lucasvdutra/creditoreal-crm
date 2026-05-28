-- name: CreateRole :one
INSERT INTO roles (tenant_id, nome, codigo, status)
VALUES ($1, $2, $3, $4)
RETURNING id, tenant_id, nome, codigo, status, created_at, updated_at, deleted_at;

-- name: ListRolesByTenant :many
SELECT id, tenant_id, nome, codigo, status, created_at, updated_at, deleted_at
FROM roles
WHERE tenant_id = $1 AND deleted_at IS NULL
ORDER BY nome;

-- name: ListPermissions :many
SELECT id, codigo, descricao, created_at
FROM permissions
ORDER BY codigo;

-- name: AddPermissionToRole :exec
INSERT INTO role_permissions (role_id, permission_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: AssignRoleToTenantUser :exec
UPDATE tenant_users
SET role_id = $3, updated_at = now()
WHERE tenant_id = $1 AND user_id = $2 AND deleted_at IS NULL;

-- name: ListUserPermissionsByTenant :many
SELECT p.codigo
FROM tenant_users tu
JOIN roles r ON r.id = tu.role_id
JOIN role_permissions rp ON rp.role_id = r.id
JOIN permissions p ON p.id = rp.permission_id
WHERE tu.user_id = $1
  AND tu.tenant_id = $2
  AND tu.status = 'ativo'
  AND tu.deleted_at IS NULL
  AND r.status = 'ativo'
  AND r.deleted_at IS NULL
ORDER BY p.codigo;
