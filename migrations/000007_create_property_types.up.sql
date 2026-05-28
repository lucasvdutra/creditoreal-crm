CREATE TABLE property_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    codigo TEXT NOT NULL UNIQUE,
    nome TEXT NOT NULL,
    categoria TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'ativo',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT property_types_status_check CHECK (status IN ('ativo', 'inativo', 'excluido'))
);

CREATE INDEX property_types_status_idx ON property_types (status) WHERE deleted_at IS NULL;

INSERT INTO property_types (codigo, nome, categoria) VALUES
    ('apartamento', 'Apartamento', 'residencial'),
    ('casa', 'Casa', 'residencial'),
    ('casa_de_condominio', 'Casa de condominio', 'residencial'),
    ('terreno', 'Terreno', 'residencial'),
    ('sala_comercial', 'Sala comercial', 'comercial'),
    ('loja', 'Loja', 'comercial'),
    ('galpao', 'Galpao', 'comercial')
ON CONFLICT (codigo) DO NOTHING;
