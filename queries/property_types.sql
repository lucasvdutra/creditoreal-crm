-- name: ListPropertyTypes :many
SELECT id, codigo, nome, categoria, status, created_at, updated_at, deleted_at
FROM property_types
WHERE deleted_at IS NULL
  AND status = 'ativo'
ORDER BY categoria, nome;

-- name: GetPropertyTypeByID :one
SELECT id, codigo, nome, categoria, status, created_at, updated_at, deleted_at
FROM property_types
WHERE id = $1 AND deleted_at IS NULL;
