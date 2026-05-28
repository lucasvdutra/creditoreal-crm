package authorization

import (
	"context"
	"testing"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type fakeStore struct {
	permissions   map[uuid.UUID]bool
	relationship  bool
	relationshipQ db.RelationshipAllowsPermissionParams
}

func (f *fakeStore) UserHasTenantPermission(_ context.Context, arg db.UserHasTenantPermissionParams) (bool, error) {
	return f.permissions[arg.TenantID], nil
}

func (f *fakeStore) RelationshipAllowsPermission(_ context.Context, arg db.RelationshipAllowsPermissionParams) (bool, error) {
	f.relationshipQ = arg
	return f.relationship, nil
}

func TestCanAllowsDirectTenantPermission(t *testing.T) {
	t.Parallel()

	userID := uuid.New()
	tenantID := uuid.New()
	service := NewService(&fakeStore{permissions: map[uuid.UUID]bool{tenantID: true}})

	allowed, err := service.Can(context.Background(), userID, tenantID, tenantID, "tenant.ler")
	if err != nil {
		t.Fatalf("can: %v", err)
	}
	if !allowed {
		t.Fatal("expected direct permission to allow access")
	}
}

func TestCanRequiresRelationshipPermissionForCrossTenant(t *testing.T) {
	t.Parallel()

	currentTenantID := uuid.New()
	resourceTenantID := uuid.New()
	store := &fakeStore{
		permissions:  map[uuid.UUID]bool{currentTenantID: true},
		relationship: true,
	}
	service := NewService(store)

	allowed, err := service.Can(context.Background(), uuid.New(), currentTenantID, resourceTenantID, "imovel.ler")
	if err != nil {
		t.Fatalf("can: %v", err)
	}
	if !allowed {
		t.Fatal("expected relationship permission to allow access")
	}
	if store.relationshipQ.TenantDestinoID != resourceTenantID {
		t.Fatalf("expected resource tenant in relationship check")
	}
}

func TestCanDeniesRelationshipWithoutActorPermission(t *testing.T) {
	t.Parallel()

	service := NewService(&fakeStore{permissions: map[uuid.UUID]bool{}, relationship: true})
	allowed, err := service.Can(context.Background(), uuid.New(), uuid.New(), uuid.New(), "imovel.ler")
	if err != nil {
		t.Fatalf("can: %v", err)
	}
	if allowed {
		t.Fatal("expected access denied without actor permission")
	}
}
