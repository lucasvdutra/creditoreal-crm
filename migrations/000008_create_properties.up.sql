CREATE TABLE properties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_dono_id UUID NOT NULL REFERENCES tenants (id),
    property_type_id UUID REFERENCES property_types (id),
    titulo TEXT NOT NULL,
    descricao TEXT,
    status TEXT NOT NULL DEFAULT 'rascunho',
    tipo_transacao TEXT NOT NULL,
    preco_venda_centavos BIGINT,
    preco_aluguel_centavos BIGINT,
    area_privativa_m2 INTEGER,
    quartos INTEGER,
    banheiros INTEGER,
    vagas_garagem INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT properties_status_check CHECK (status IN ('rascunho', 'ativo', 'inativo', 'vendido', 'alugado', 'excluido')),
    CONSTRAINT properties_tipo_transacao_check CHECK (tipo_transacao IN ('venda', 'aluguel', 'venda_aluguel')),
    CONSTRAINT properties_preco_venda_check CHECK (preco_venda_centavos IS NULL OR preco_venda_centavos >= 0),
    CONSTRAINT properties_preco_aluguel_check CHECK (preco_aluguel_centavos IS NULL OR preco_aluguel_centavos >= 0)
);

CREATE INDEX properties_tenant_dono_id_idx ON properties (tenant_dono_id) WHERE deleted_at IS NULL;
CREATE INDEX properties_tenant_status_idx ON properties (tenant_dono_id, status) WHERE deleted_at IS NULL;
CREATE INDEX properties_tenant_transaction_idx ON properties (tenant_dono_id, tipo_transacao) WHERE deleted_at IS NULL;
CREATE INDEX properties_property_type_idx ON properties (property_type_id) WHERE deleted_at IS NULL;
