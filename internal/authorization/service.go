package authorization

import (
	"context"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type Store interface {
	UserHasTenantPermission(context.Context, db.UserHasTenantPermissionParams) (bool, error)
	RelationshipAllowsPermission(context.Context, db.RelationshipAllowsPermissionParams) (bool, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) Can(ctx context.Context, userID, currentTenantID, resourceTenantID uuid.UUID, permission string) (bool, error) {
	hasDirectPermission, err := s.store.UserHasTenantPermission(ctx, db.UserHasTenantPermissionParams{
		UserID:   userID,
		TenantID: resourceTenantID,
		Codigo:   permission,
	})
	if err != nil || hasDirectPermission || currentTenantID == resourceTenantID {
		return hasDirectPermission, err
	}

	hasActorPermission, err := s.store.UserHasTenantPermission(ctx, db.UserHasTenantPermissionParams{
		UserID:   userID,
		TenantID: currentTenantID,
		Codigo:   permission,
	})
	if err != nil || !hasActorPermission {
		return false, err
	}

	return s.store.RelationshipAllowsPermission(ctx, db.RelationshipAllowsPermissionParams{
		TenantOrigemID:  currentTenantID,
		TenantDestinoID: resourceTenantID,
		Column3:         permission,
	})
}
