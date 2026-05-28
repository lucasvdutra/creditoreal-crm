package rbac

import (
	"context"
	"strings"

	"creditoreal-crm/internal/domain"
	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type Store interface {
	CreateRole(context.Context, db.CreateRoleParams) (db.Role, error)
	ListRolesByTenant(context.Context, uuid.UUID) ([]db.Role, error)
	ListPermissions(context.Context) ([]db.Permission, error)
	AddPermissionToRole(context.Context, db.AddPermissionToRoleParams) error
	AssignRoleToTenantUser(context.Context, db.AssignRoleToTenantUserParams) error
	ListUserPermissionsByTenant(context.Context, db.ListUserPermissionsByTenantParams) ([]string, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) CreateRole(ctx context.Context, req CreateRoleRequest) (RoleResponse, error) {
	if req.Status == "" {
		req.Status = "ativo"
	}
	if strings.TrimSpace(req.Nome) == "" || strings.TrimSpace(req.Codigo) == "" || !validStatus(req.Status) {
		return RoleResponse{}, domain.ErrInvalidInput
	}
	tenantID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return RoleResponse{}, domain.ErrInvalidInput
	}

	item, err := s.store.CreateRole(ctx, db.CreateRoleParams{
		TenantID: tenantID,
		Nome:     strings.TrimSpace(req.Nome),
		Codigo:   strings.TrimSpace(req.Codigo),
		Status:   req.Status,
	})
	if err != nil {
		return RoleResponse{}, err
	}
	return roleResponse(item), nil
}

func (s *Service) ListRoles(ctx context.Context, tenantID uuid.UUID) ([]RoleResponse, error) {
	items, err := s.store.ListRolesByTenant(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	out := make([]RoleResponse, 0, len(items))
	for _, item := range items {
		out = append(out, roleResponse(item))
	}
	return out, nil
}

func (s *Service) ListPermissions(ctx context.Context) ([]PermissionResponse, error) {
	items, err := s.store.ListPermissions(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]PermissionResponse, 0, len(items))
	for _, item := range items {
		out = append(out, permissionResponse(item))
	}
	return out, nil
}

func (s *Service) AddPermission(ctx context.Context, req AddPermissionRequest) error {
	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		return domain.ErrInvalidInput
	}
	permissionID, err := uuid.Parse(req.PermissionID)
	if err != nil {
		return domain.ErrInvalidInput
	}
	return s.store.AddPermissionToRole(ctx, db.AddPermissionToRoleParams{RoleID: roleID, PermissionID: permissionID})
}

func (s *Service) AssignRole(ctx context.Context, req AssignRoleRequest) error {
	tenantID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return domain.ErrInvalidInput
	}
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return domain.ErrInvalidInput
	}
	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		return domain.ErrInvalidInput
	}
	return s.store.AssignRoleToTenantUser(ctx, db.AssignRoleToTenantUserParams{
		TenantID: tenantID,
		UserID:   userID,
		RoleID:   uuid.NullUUID{UUID: roleID, Valid: true},
	})
}

func (s *Service) ListUserPermissions(ctx context.Context, userID, tenantID uuid.UUID) ([]string, error) {
	return s.store.ListUserPermissionsByTenant(ctx, db.ListUserPermissionsByTenantParams{
		UserID:   userID,
		TenantID: tenantID,
	})
}

func validStatus(value string) bool {
	switch value {
	case "ativo", "inativo", "pendente", "bloqueado", "excluido":
		return true
	default:
		return false
	}
}
