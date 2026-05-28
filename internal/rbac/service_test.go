package rbac

import (
	"context"
	"testing"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type fakeStore struct {
	role db.CreateRoleParams
}

func (f *fakeStore) CreateRole(_ context.Context, arg db.CreateRoleParams) (db.Role, error) {
	f.role = arg
	return db.Role{ID: uuid.New(), TenantID: arg.TenantID, Nome: arg.Nome, Codigo: arg.Codigo, Status: arg.Status}, nil
}
func (f *fakeStore) ListRolesByTenant(context.Context, uuid.UUID) ([]db.Role, error) {
	return nil, nil
}
func (f *fakeStore) ListPermissions(context.Context) ([]db.Permission, error) {
	return nil, nil
}
func (f *fakeStore) AddPermissionToRole(context.Context, db.AddPermissionToRoleParams) error {
	return nil
}
func (f *fakeStore) AssignRoleToTenantUser(context.Context, db.AssignRoleToTenantUserParams) error {
	return nil
}
func (f *fakeStore) ListUserPermissionsByTenant(context.Context, db.ListUserPermissionsByTenantParams) ([]string, error) {
	return nil, nil
}

func TestCreateRoleDefaultsStatusAtivo(t *testing.T) {
	t.Parallel()

	store := &fakeStore{}
	service := NewService(store)
	item, err := service.CreateRole(context.Background(), CreateRoleRequest{
		TenantID: uuid.New().String(),
		Nome:     "Administrador",
		Codigo:   "administrador",
	})
	if err != nil {
		t.Fatalf("create role: %v", err)
	}
	if item.Status != "ativo" || store.role.Status != "ativo" {
		t.Fatalf("expected status ativo")
	}
}
