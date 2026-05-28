package tenant

import (
	"context"
	"database/sql"
	"testing"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type fakeTenantStore struct {
	created db.CreateTenantParams
}

func (f *fakeTenantStore) CreateTenant(_ context.Context, arg db.CreateTenantParams) (db.Tenant, error) {
	f.created = arg
	return db.Tenant{ID: uuid.New(), Nome: arg.Nome, Status: arg.Status, Tipo: arg.Tipo}, nil
}
func (f *fakeTenantStore) GetTenantByID(context.Context, uuid.UUID) (db.Tenant, error) {
	return db.Tenant{}, sql.ErrNoRows
}
func (f *fakeTenantStore) ListTenants(context.Context, db.ListTenantsParams) ([]db.Tenant, error) {
	return nil, nil
}
func (f *fakeTenantStore) CountTenants(context.Context, db.CountTenantsParams) (int64, error) {
	return 0, nil
}
func (f *fakeTenantStore) UpdateTenant(context.Context, db.UpdateTenantParams) (db.Tenant, error) {
	return db.Tenant{}, nil
}
func (f *fakeTenantStore) SoftDeleteTenant(context.Context, uuid.UUID) error {
	return nil
}

func TestCreateTenantValidatesRequiredFields(t *testing.T) {
	t.Parallel()

	service := NewService(&fakeTenantStore{})
	if _, err := service.Create(context.Background(), CreateRequest{Tipo: "imobiliaria"}); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestCreateTenantDefaultsStatusAtivo(t *testing.T) {
	t.Parallel()

	store := &fakeTenantStore{}
	service := NewService(store)
	item, err := service.Create(context.Background(), CreateRequest{
		Nome: "Credito Real",
		Tipo: "imobiliaria",
	})
	if err != nil {
		t.Fatalf("create tenant: %v", err)
	}
	if item.Status != "ativo" {
		t.Fatalf("expected status ativo, got %q", item.Status)
	}
	if store.created.Status != "ativo" {
		t.Fatalf("expected persisted status ativo, got %q", store.created.Status)
	}
}
