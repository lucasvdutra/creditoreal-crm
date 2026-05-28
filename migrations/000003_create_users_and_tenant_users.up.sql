CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome TEXT NOT NULL,
    email TEXT NOT NULL,
    senha_hash TEXT NOT NULL,
    telefone TEXT,
    status TEXT NOT NULL DEFAULT 'ativo',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT users_status_check CHECK (status IN ('ativo', 'inativo', 'pendente', 'bloqueado', 'excluido'))
);

CREATE UNIQUE INDEX users_email_unique ON users (lower(email)) WHERE deleted_at IS NULL;
CREATE INDEX users_status_idx ON users (status) WHERE deleted_at IS NULL;

CREATE TABLE tenant_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants (id),
    user_id UUID NOT NULL REFERENCES users (id),
    cargo TEXT,
    status TEXT NOT NULL DEFAULT 'ativo',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT tenant_users_status_check CHECK (status IN ('ativo', 'inativo', 'pendente', 'bloqueado', 'excluido'))
);

CREATE UNIQUE INDEX tenant_users_tenant_user_unique
    ON tenant_users (tenant_id, user_id)
    WHERE deleted_at IS NULL;

CREATE INDEX tenant_users_tenant_id_idx ON tenant_users (tenant_id) WHERE deleted_at IS NULL;
CREATE INDEX tenant_users_user_id_idx ON tenant_users (user_id) WHERE deleted_at IS NULL;
