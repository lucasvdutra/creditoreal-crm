CREATE TABLE tenant_relationships (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_origem_id UUID NOT NULL REFERENCES tenants (id),
    tenant_destino_id UUID NOT NULL REFERENCES tenants (id),
    tipo TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'ativo',
    permissoes TEXT[] NOT NULL DEFAULT '{}'::text[],
    observacao TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT tenant_relationships_distinct_check CHECK (tenant_origem_id <> tenant_destino_id),
    CONSTRAINT tenant_relationships_status_check CHECK (status IN ('ativo', 'inativo', 'pendente', 'bloqueado', 'excluido')),
    CONSTRAINT tenant_relationships_tipo_check CHECK (tipo IN ('parceria', 'administracao', 'captacao', 'repasse', 'outro'))
);

CREATE UNIQUE INDEX tenant_relationships_pair_unique
    ON tenant_relationships (tenant_origem_id, tenant_destino_id, tipo)
    WHERE deleted_at IS NULL;

CREATE INDEX tenant_relationships_origem_idx ON tenant_relationships (tenant_origem_id) WHERE deleted_at IS NULL;
CREATE INDEX tenant_relationships_destino_idx ON tenant_relationships (tenant_destino_id) WHERE deleted_at IS NULL;
