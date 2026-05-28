package relationship

import (
	"context"
	"database/sql"
	"testing"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type fakeStore struct {
	created db.CreateTenantRelationshipParams
}

func (f *fakeStore) CreateTenantRelationship(_ context.Context, arg db.CreateTenantRelationshipParams) (db.TenantRelationship, error) {
	f.created = arg
	return db.TenantRelationship{ID: uuid.New(), TenantOrigemID: arg.TenantOrigemID, TenantDestinoID: arg.TenantDestinoID, Tipo: arg.Tipo, Status: arg.Status, Permissoes: arg.Permissoes}, nil
}
func (f *fakeStore) GetTenantRelationshipByID(context.Context, uuid.UUID) (db.TenantRelationship, error) {
	return db.TenantRelationship{}, sql.ErrNoRows
}
func (f *fakeStore) ListTenantRelationships(context.Context, db.ListTenantRelationshipsParams) ([]db.TenantRelationship, error) {
	return nil, nil
}
func (f *fakeStore) UpdateTenantRelationship(context.Context, db.UpdateTenantRelationshipParams) (db.TenantRelationship, error) {
	return db.TenantRelationship{}, nil
}
func (f *fakeStore) SoftDeleteTenantRelationship(context.Context, uuid.UUID) error {
	return nil
}

func TestCreateRelationshipRequiresDistinctTenants(t *testing.T) {
	t.Parallel()

	id := uuid.New().String()
	service := NewService(&fakeStore{})
	_, err := service.Create(context.Background(), CreateRequest{
		TenantOrigemID:  id,
		TenantDestinoID: id,
		Tipo:            "parceria",
	})
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestCreateRelationshipDoesNotGrantAccessWithoutPermissions(t *testing.T) {
	t.Parallel()

	store := &fakeStore{}
	service := NewService(store)
	_, err := service.Create(context.Background(), CreateRequest{
		TenantOrigemID:  uuid.New().String(),
		TenantDestinoID: uuid.New().String(),
		Tipo:            "parceria",
	})
	if err != nil {
		t.Fatalf("create relationship: %v", err)
	}
	if len(store.created.Permissoes) != 0 {
		t.Fatalf("expected no implicit permissions, got %v", store.created.Permissoes)
	}
}
