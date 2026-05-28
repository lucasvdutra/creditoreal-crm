CREATE TABLE property_addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    property_id UUID NOT NULL UNIQUE REFERENCES properties (id) ON DELETE CASCADE,
    cep TEXT NOT NULL,
    estado TEXT NOT NULL,
    cidade TEXT NOT NULL,
    bairro TEXT NOT NULL,
    logradouro TEXT,
    numero TEXT,
    complemento TEXT,
    exibicao_endereco TEXT NOT NULL DEFAULT 'bairro',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT property_addresses_exibicao_check CHECK (exibicao_endereco IN ('completo', 'bairro', 'cidade', 'oculto'))
);

CREATE INDEX property_addresses_city_idx ON property_addresses (cidade, bairro);
