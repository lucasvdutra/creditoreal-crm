CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome TEXT NOT NULL,
    nome_fantasia TEXT,
    tipo_documento TEXT,
    numero_documento TEXT,
    email TEXT,
    telefone TEXT,
    site TEXT,
    logo_url TEXT,
    status TEXT NOT NULL DEFAULT 'ativo',
    tipo TEXT NOT NULL,
    configuracoes JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT tenants_status_check CHECK (status IN ('ativo', 'inativo', 'pendente', 'bloqueado', 'excluido')),
    CONSTRAINT tenants_tipo_check CHECK (tipo IN ('imobiliaria', 'incorporadora', 'corretor_autonomo', 'administradora', 'parceiro')),
    CONSTRAINT tenants_tipo_documento_check CHECK (tipo_documento IS NULL OR tipo_documento IN ('cpf', 'cnpj'))
);

CREATE UNIQUE INDEX tenants_numero_documento_unique
    ON tenants (numero_documento)
    WHERE numero_documento IS NOT NULL AND deleted_at IS NULL;

CREATE INDEX tenants_status_idx ON tenants (status) WHERE deleted_at IS NULL;
CREATE INDEX tenants_tipo_idx ON tenants (tipo) WHERE deleted_at IS NULL;
