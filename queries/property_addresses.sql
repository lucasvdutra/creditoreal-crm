-- name: UpsertPropertyAddress :one
INSERT INTO property_addresses (
    property_id,
    cep,
    estado,
    cidade,
    bairro,
    logradouro,
    numero,
    complemento,
    exibicao_endereco
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
ON CONFLICT (property_id) DO UPDATE SET
    cep = EXCLUDED.cep,
    estado = EXCLUDED.estado,
    cidade = EXCLUDED.cidade,
    bairro = EXCLUDED.bairro,
    logradouro = EXCLUDED.logradouro,
    numero = EXCLUDED.numero,
    complemento = EXCLUDED.complemento,
    exibicao_endereco = EXCLUDED.exibicao_endereco,
    updated_at = now()
RETURNING id, property_id, cep, estado, cidade, bairro, logradouro, numero, complemento, exibicao_endereco, created_at, updated_at;

-- name: GetPropertyAddress :one
SELECT id, property_id, cep, estado, cidade, bairro, logradouro, numero, complemento, exibicao_endereco, created_at, updated_at
FROM property_addresses
WHERE property_id = $1;
