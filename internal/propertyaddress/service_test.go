package propertyaddress

import (
	"context"
	"database/sql"
	"testing"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type fakeStore struct {
	upsert db.UpsertPropertyAddressParams
	item   db.Property
}

func (f *fakeStore) GetPropertyByID(context.Context, uuid.UUID) (db.Property, error) {
	if f.item.ID == uuid.Nil {
		return db.Property{}, sql.ErrNoRows
	}
	return f.item, nil
}
func (f *fakeStore) UpsertPropertyAddress(_ context.Context, arg db.UpsertPropertyAddressParams) (db.PropertyAddress, error) {
	f.upsert = arg
	return db.PropertyAddress{ID: uuid.New(), PropertyID: arg.PropertyID, Cep: arg.Cep, Estado: arg.Estado, Cidade: arg.Cidade, Bairro: arg.Bairro, ExibicaoEndereco: arg.ExibicaoEndereco}, nil
}
func (f *fakeStore) GetPropertyAddress(context.Context, uuid.UUID) (db.PropertyAddress, error) {
	return db.PropertyAddress{}, nil
}

type fakeAuthorizer struct {
	allowed bool
}

func (f fakeAuthorizer) Can(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, string) (bool, error) {
	return f.allowed, nil
}

func TestUpsertAddressDefaultsExibicaoEndereco(t *testing.T) {
	t.Parallel()

	propertyID := uuid.New()
	store := &fakeStore{item: db.Property{ID: propertyID, TenantDonoID: uuid.New()}}
	service := NewService(store, fakeAuthorizer{allowed: true})
	item, err := service.Upsert(context.Background(), uuid.New(), uuid.New(), propertyID, UpsertRequest{
		CEP:    "90000-000",
		Estado: "RS",
		Cidade: "Porto Alegre",
		Bairro: "Centro",
	})
	if err != nil {
		t.Fatalf("upsert address: %v", err)
	}
	if item.ExibicaoEndereco != "bairro" || store.upsert.ExibicaoEndereco != "bairro" {
		t.Fatalf("expected default exibicao_endereco bairro")
	}
}
