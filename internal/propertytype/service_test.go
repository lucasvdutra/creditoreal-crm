package propertytype

import (
	"context"
	"testing"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type fakeStore struct{}

func (fakeStore) ListPropertyTypes(context.Context) ([]db.PropertyType, error) {
	return []db.PropertyType{{ID: uuid.New(), Codigo: "apartamento", Nome: "Apartamento", Categoria: "residencial", Status: "ativo"}}, nil
}

func TestListPropertyTypes(t *testing.T) {
	t.Parallel()

	items, err := NewService(fakeStore{}).List(context.Background())
	if err != nil {
		t.Fatalf("list property types: %v", err)
	}
	if len(items) != 1 || items[0].Codigo != "apartamento" {
		t.Fatalf("unexpected property types: %+v", items)
	}
}
