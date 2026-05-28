package property

import (
	"context"
	"database/sql"
	"testing"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type fakeStore struct {
	created db.CreatePropertyParams
	item    db.Property
}

func (f *fakeStore) CreateProperty(_ context.Context, arg db.CreatePropertyParams) (db.Property, error) {
	f.created = arg
	return db.Property{ID: uuid.New(), TenantDonoID: arg.TenantDonoID, Titulo: arg.Titulo, Status: arg.Status, TipoTransacao: arg.TipoTransacao}, nil
}
func (f *fakeStore) GetPropertyByID(context.Context, uuid.UUID) (db.Property, error) {
	if f.item.ID == uuid.Nil {
		return db.Property{}, sql.ErrNoRows
	}
	return f.item, nil
}
func (f *fakeStore) ListPropertiesByTenant(context.Context, db.ListPropertiesByTenantParams) ([]db.Property, error) {
	return nil, nil
}
func (f *fakeStore) CountPropertiesByTenant(context.Context, db.CountPropertiesByTenantParams) (int64, error) {
	return 0, nil
}
func (f *fakeStore) UpdateProperty(context.Context, db.UpdatePropertyParams) (db.Property, error) {
	return db.Property{}, nil
}
func (f *fakeStore) SoftDeleteProperty(context.Context, uuid.UUID) error {
	return nil
}

type fakeAuthorizer struct {
	allowed bool
}

func (f fakeAuthorizer) Can(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, string) (bool, error) {
	return f.allowed, nil
}

func TestCreatePropertyRequiresAuthorization(t *testing.T) {
	t.Parallel()

	service := NewService(&fakeStore{}, fakeAuthorizer{allowed: false})
	_, err := service.Create(context.Background(), uuid.New(), uuid.New(), CreateRequest{
		Titulo:        "Apartamento",
		TipoTransacao: "venda",
	})
	if err == nil {
		t.Fatal("expected authorization error")
	}
}

func TestCreatePropertyDefaultsStatusRascunho(t *testing.T) {
	t.Parallel()

	store := &fakeStore{}
	service := NewService(store, fakeAuthorizer{allowed: true})
	item, err := service.Create(context.Background(), uuid.New(), uuid.New(), CreateRequest{
		Titulo:        "Apartamento",
		TipoTransacao: "venda",
	})
	if err != nil {
		t.Fatalf("create property: %v", err)
	}
	if item.Status != "rascunho" || store.created.Status != "rascunho" {
		t.Fatalf("expected status rascunho")
	}
}
