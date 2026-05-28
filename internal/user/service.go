package user

import (
	"context"
	"database/sql"
	"strings"

	"creditoreal-crm/internal/auth/password"
	"creditoreal-crm/internal/domain"
	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type Store interface {
	CreateUser(context.Context, db.CreateUserParams) (db.User, error)
	GetUserByID(context.Context, uuid.UUID) (db.User, error)
	ListUsers(context.Context, db.ListUsersParams) ([]db.User, error)
	CountUsers(context.Context, sql.NullString) (int64, error)
	CreateTenantUser(context.Context, db.CreateTenantUserParams) (db.TenantUser, error)
	ListTenantUsers(context.Context, db.ListTenantUsersParams) ([]db.TenantUser, error)
	ListUserTenants(context.Context, uuid.UUID) ([]db.TenantUser, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (Response, error) {
	if req.Status == "" {
		req.Status = "ativo"
	}
	if err := validateCreate(req); err != nil {
		return Response{}, err
	}

	hash, err := password.Hash(req.Senha)
	if err != nil {
		return Response{}, err
	}

	item, err := s.store.CreateUser(ctx, db.CreateUserParams{
		Nome:      strings.TrimSpace(req.Nome),
		Lower:     strings.TrimSpace(req.Email),
		SenhaHash: hash,
		Telefone:  nullString(req.Telefone),
		Status:    req.Status,
	})
	if err != nil {
		return Response{}, err
	}

	return toResponse(item), nil
}

func (s *Service) List(ctx context.Context, pageSize int, offset int32, status string) ([]Response, int64, error) {
	statusParam := sql.NullString{String: status, Valid: status != ""}
	items, err := s.store.ListUsers(ctx, db.ListUsersParams{
		Limit:  int32(pageSize),
		Offset: offset,
		Status: statusParam,
	})
	if err != nil {
		return nil, 0, err
	}
	total, err := s.store.CountUsers(ctx, statusParam)
	if err != nil {
		return nil, 0, err
	}

	out := make([]Response, 0, len(items))
	for _, item := range items {
		out = append(out, toResponse(item))
	}
	return out, total, nil
}

func (s *Service) CreateTenantUser(ctx context.Context, req CreateTenantUserRequest) (TenantUserResponse, error) {
	if req.Status == "" {
		req.Status = "ativo"
	}
	if !validStatus(req.Status) {
		return TenantUserResponse{}, domain.ErrInvalidInput
	}
	tenantID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return TenantUserResponse{}, domain.ErrInvalidInput
	}
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return TenantUserResponse{}, domain.ErrInvalidInput
	}

	item, err := s.store.CreateTenantUser(ctx, db.CreateTenantUserParams{
		TenantID: tenantID,
		UserID:   userID,
		Cargo:    nullString(req.Cargo),
		Status:   req.Status,
	})
	if err != nil {
		return TenantUserResponse{}, err
	}
	return toTenantUserResponse(item), nil
}

func (s *Service) ListTenantUsers(ctx context.Context, tenantID uuid.UUID, pageSize int, offset int32) ([]TenantUserResponse, error) {
	items, err := s.store.ListTenantUsers(ctx, db.ListTenantUsersParams{
		TenantID: tenantID,
		Limit:    int32(pageSize),
		Offset:   offset,
	})
	if err != nil {
		return nil, err
	}
	out := make([]TenantUserResponse, 0, len(items))
	for _, item := range items {
		out = append(out, toTenantUserResponse(item))
	}
	return out, nil
}

func validateCreate(req CreateRequest) error {
	if strings.TrimSpace(req.Nome) == "" || strings.TrimSpace(req.Email) == "" || len(req.Senha) < 8 {
		return domain.ErrInvalidInput
	}
	if !validStatus(req.Status) {
		return domain.ErrInvalidInput
	}
	return nil
}

func validStatus(value string) bool {
	switch value {
	case "ativo", "inativo", "pendente", "bloqueado", "excluido":
		return true
	default:
		return false
	}
}

func nullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: strings.TrimSpace(*value), Valid: true}
}
