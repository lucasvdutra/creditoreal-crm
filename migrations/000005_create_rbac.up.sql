CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    codigo TEXT NOT NULL UNIQUE,
    descricao TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants (id),
    nome TEXT NOT NULL,
    codigo TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'ativo',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT roles_status_check CHECK (status IN ('ativo', 'inativo', 'pendente', 'bloqueado', 'excluido'))
);

CREATE UNIQUE INDEX roles_tenant_codigo_unique
    ON roles (tenant_id, codigo)
    WHERE deleted_at IS NULL;

CREATE INDEX roles_tenant_id_idx ON roles (tenant_id) WHERE deleted_at IS NULL;

CREATE TABLE role_permissions (
    role_id UUID NOT NULL REFERENCES roles (id) ON DELETE CASCADE,
    permission_id UUID NOT NULL REFERENCES permissions (id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (role_id, permission_id)
);

ALTER TABLE tenant_users
    ADD COLUMN role_id UUID REFERENCES roles (id);

CREATE INDEX tenant_users_role_id_idx ON tenant_users (role_id) WHERE deleted_at IS NULL;

INSERT INTO permissions (codigo, descricao) VALUES
    ('tenant.ler', 'Consultar tenants'),
    ('tenant.criar', 'Criar tenants'),
    ('tenant.atualizar', 'Atualizar tenants'),
    ('tenant.excluir', 'Excluir tenants'),
    ('usuario.gerenciar', 'Gerenciar usuarios e vinculos'),
    ('relacionamento.gerenciar', 'Gerenciar relacionamentos entre tenants'),
    ('imovel.ler', 'Consultar imoveis'),
    ('imovel.gerenciar', 'Gerenciar imoveis')
ON CONFLICT (codigo) DO NOTHING;
