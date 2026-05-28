# Credito Real CRM

## Desenvolvimento local

Subir os servicos locais:

```powershell
docker compose up
```

Aplicar migrations:

```powershell
docker compose --profile tools run --rm migrate
```

Gerar codigo sqlc:

```powershell
docker compose --profile tools run --rm sqlc
```

Rodar testes Go:

```powershell
go test ./...
```

Endpoints iniciais:

```text
GET    /health
POST   /api/tenants
GET    /api/tenants
GET    /api/tenants/{id}
PATCH  /api/tenants/{id}
DELETE /api/tenants/{id}
POST   /api/users
GET    /api/users
POST   /api/tenant-users
GET    /api/tenants/{tenant_id}/users
POST   /api/auth/login
POST   /api/auth/refresh
GET    /api/auth/me
POST   /api/rbac/roles
GET    /api/rbac/roles?tenant_id={tenant_id}
GET    /api/rbac/permissions
POST   /api/rbac/role-permissions
POST   /api/rbac/assign-role
GET    /api/rbac/user-permissions?user_id={user_id}&tenant_id={tenant_id}
POST   /api/tenant-relationships
GET    /api/tenant-relationships?tenant_id={tenant_id}
GET    /api/tenant-relationships/{id}
PATCH  /api/tenant-relationships/{id}
DELETE /api/tenant-relationships/{id}
GET    /api/property-types
POST   /api/properties
GET    /api/properties
GET    /api/properties/{id}
PATCH  /api/properties/{id}
DELETE /api/properties/{id}
GET    /api/properties/{id}/address
PUT    /api/properties/{id}/address
```
