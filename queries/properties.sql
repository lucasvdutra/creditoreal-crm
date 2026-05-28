-- name: CreateProperty :one
INSERT INTO properties (
    tenant_dono_id,
    property_type_id,
    titulo,
    descricao,
    status,
    tipo_transacao,
    preco_venda_centavos,
    preco_aluguel_centavos,
    area_privativa_m2,
    quartos,
    banheiros,
    vagas_garagem
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id, tenant_dono_id, property_type_id, titulo, descricao, status, tipo_transacao, preco_venda_centavos, preco_aluguel_centavos, area_privativa_m2, quartos, banheiros, vagas_garagem, created_at, updated_at, deleted_at;

-- name: GetPropertyByID :one
SELECT id, tenant_dono_id, property_type_id, titulo, descricao, status, tipo_transacao, preco_venda_centavos, preco_aluguel_centavos, area_privativa_m2, quartos, banheiros, vagas_garagem, created_at, updated_at, deleted_at
FROM properties
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListPropertiesByTenant :many
SELECT id, tenant_dono_id, property_type_id, titulo, descricao, status, tipo_transacao, preco_venda_centavos, preco_aluguel_centavos, area_privativa_m2, quartos, banheiros, vagas_garagem, created_at, updated_at, deleted_at
FROM properties
WHERE tenant_dono_id = $1
  AND deleted_at IS NULL
  AND (sqlc.narg('status')::text IS NULL OR status = sqlc.narg('status')::text)
  AND (sqlc.narg('tipo_transacao')::text IS NULL OR tipo_transacao = sqlc.narg('tipo_transacao')::text)
  AND (sqlc.narg('property_type_id')::uuid IS NULL OR property_type_id = sqlc.narg('property_type_id')::uuid)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountPropertiesByTenant :one
SELECT count(*)::bigint
FROM properties
WHERE tenant_dono_id = $1
  AND deleted_at IS NULL
  AND (sqlc.narg('status')::text IS NULL OR status = sqlc.narg('status')::text)
  AND (sqlc.narg('tipo_transacao')::text IS NULL OR tipo_transacao = sqlc.narg('tipo_transacao')::text)
  AND (sqlc.narg('property_type_id')::uuid IS NULL OR property_type_id = sqlc.narg('property_type_id')::uuid);

-- name: UpdateProperty :one
UPDATE properties
SET
    property_type_id = COALESCE(sqlc.narg('property_type_id'), property_type_id),
    titulo = COALESCE(sqlc.narg('titulo'), titulo),
    descricao = COALESCE(sqlc.narg('descricao'), descricao),
    status = COALESCE(sqlc.narg('status'), status),
    tipo_transacao = COALESCE(sqlc.narg('tipo_transacao'), tipo_transacao),
    preco_venda_centavos = COALESCE(sqlc.narg('preco_venda_centavos'), preco_venda_centavos),
    preco_aluguel_centavos = COALESCE(sqlc.narg('preco_aluguel_centavos'), preco_aluguel_centavos),
    area_privativa_m2 = COALESCE(sqlc.narg('area_privativa_m2'), area_privativa_m2),
    quartos = COALESCE(sqlc.narg('quartos'), quartos),
    banheiros = COALESCE(sqlc.narg('banheiros'), banheiros),
    vagas_garagem = COALESCE(sqlc.narg('vagas_garagem'), vagas_garagem),
    updated_at = now()
WHERE id = sqlc.arg('id') AND deleted_at IS NULL
RETURNING id, tenant_dono_id, property_type_id, titulo, descricao, status, tipo_transacao, preco_venda_centavos, preco_aluguel_centavos, area_privativa_m2, quartos, banheiros, vagas_garagem, created_at, updated_at, deleted_at;

-- name: SoftDeleteProperty :exec
UPDATE properties
SET deleted_at = now(), status = 'excluido', updated_at = now()
WHERE id = $1 AND deleted_at IS NULL;
