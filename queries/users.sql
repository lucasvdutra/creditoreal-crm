-- name: CreateUser :one
INSERT INTO users (nome, email, senha_hash, telefone, status)
VALUES ($1, lower($2), $3, $4, $5)
RETURNING id, nome, email, senha_hash, telefone, status, created_at, updated_at, deleted_at;

-- name: GetUserByID :one
SELECT id, nome, email, senha_hash, telefone, status, created_at, updated_at, deleted_at
FROM users
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT id, nome, email, senha_hash, telefone, status, created_at, updated_at, deleted_at
FROM users
WHERE lower(email) = lower($1) AND deleted_at IS NULL;

-- name: ListUsers :many
SELECT id, nome, email, senha_hash, telefone, status, created_at, updated_at, deleted_at
FROM users
WHERE deleted_at IS NULL
  AND (sqlc.narg('status')::text IS NULL OR status = sqlc.narg('status')::text)
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT count(*)::bigint
FROM users
WHERE deleted_at IS NULL
  AND (sqlc.narg('status')::text IS NULL OR status = sqlc.narg('status')::text);

-- name: CreateTenantUser :one
INSERT INTO tenant_users (tenant_id, user_id, cargo, status)
VALUES ($1, $2, $3, $4)
RETURNING id, tenant_id, user_id, cargo, status, created_at, updated_at, deleted_at, role_id;

-- name: GetTenantUser :one
SELECT id, tenant_id, user_id, cargo, status, created_at, updated_at, deleted_at, role_id
FROM tenant_users
WHERE tenant_id = $1 AND user_id = $2 AND deleted_at IS NULL;

-- name: ListTenantUsers :many
SELECT id, tenant_id, user_id, cargo, status, created_at, updated_at, deleted_at, role_id
FROM tenant_users
WHERE tenant_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListUserTenants :many
SELECT id, tenant_id, user_id, cargo, status, created_at, updated_at, deleted_at, role_id
FROM tenant_users
WHERE user_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC;
