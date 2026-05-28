-- name: CreateTenant :one
INSERT INTO tenants (
    nome,
    nome_fantasia,
    tipo_documento,
    numero_documento,
    email,
    telefone,
    site,
    logo_url,
    status,
    tipo,
    configuracoes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, COALESCE($11, '{}'::jsonb)
)
RETURNING id, nome, nome_fantasia, tipo_documento, numero_documento, email, telefone, site, logo_url, status, tipo, configuracoes, created_at, updated_at, deleted_at;

-- name: GetTenantByID :one
SELECT id, nome, nome_fantasia, tipo_documento, numero_documento, email, telefone, site, logo_url, status, tipo, configuracoes, created_at, updated_at, deleted_at
FROM tenants
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListTenants :many
SELECT id, nome, nome_fantasia, tipo_documento, numero_documento, email, telefone, site, logo_url, status, tipo, configuracoes, created_at, updated_at, deleted_at
FROM tenants
WHERE deleted_at IS NULL
  AND (sqlc.narg('status')::text IS NULL OR status = sqlc.narg('status')::text)
  AND (sqlc.narg('tipo')::text IS NULL OR tipo = sqlc.narg('tipo')::text)
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountTenants :one
SELECT count(*)::bigint
FROM tenants
WHERE deleted_at IS NULL
  AND (sqlc.narg('status')::text IS NULL OR status = sqlc.narg('status')::text)
  AND (sqlc.narg('tipo')::text IS NULL OR tipo = sqlc.narg('tipo')::text);

-- name: UpdateTenant :one
UPDATE tenants
SET
    nome = COALESCE(sqlc.narg('nome'), nome),
    nome_fantasia = COALESCE(sqlc.narg('nome_fantasia'), nome_fantasia),
    tipo_documento = COALESCE(sqlc.narg('tipo_documento'), tipo_documento),
    numero_documento = COALESCE(sqlc.narg('numero_documento'), numero_documento),
    email = COALESCE(sqlc.narg('email'), email),
    telefone = COALESCE(sqlc.narg('telefone'), telefone),
    site = COALESCE(sqlc.narg('site'), site),
    logo_url = COALESCE(sqlc.narg('logo_url'), logo_url),
    status = COALESCE(sqlc.narg('status'), status),
    tipo = COALESCE(sqlc.narg('tipo'), tipo),
    configuracoes = COALESCE(sqlc.narg('configuracoes'), configuracoes),
    updated_at = now()
WHERE id = sqlc.arg('id') AND deleted_at IS NULL
RETURNING id, nome, nome_fantasia, tipo_documento, numero_documento, email, telefone, site, logo_url, status, tipo, configuracoes, created_at, updated_at, deleted_at;

-- name: SoftDeleteTenant :exec
UPDATE tenants
SET deleted_at = now(), status = 'excluido', updated_at = now()
WHERE id = $1 AND deleted_at IS NULL;
