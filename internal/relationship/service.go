package relationship

import (
	"context"
	"database/sql"
	"strings"

	"creditoreal-crm/internal/domain"
	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type Store interface {
	CreateTenantRelationship(context.Context, db.CreateTenantRelationshipParams) (db.TenantRelationship, error)
	GetTenantRelationshipByID(context.Context, uuid.UUID) (db.TenantRelationship, error)
	ListTenantRelationships(context.Context, db.ListTenantRelationshipsParams) ([]db.TenantRelationship, error)
	UpdateTenantRelationship(context.Context, db.UpdateTenantRelationshipParams) (db.TenantRelationship, error)
	SoftDeleteTenantRelationship(context.Context, uuid.UUID) error
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
	if !validTipo(req.Tipo) || !validStatus(req.Status) {
		return Response{}, domain.ErrInvalidInput
	}
	origemID, err := uuid.Parse(req.TenantOrigemID)
	if err != nil {
		return Response{}, domain.ErrInvalidInput
	}
	destinoID, err := uuid.Parse(req.TenantDestinoID)
	if err != nil || destinoID == origemID {
		return Response{}, domain.ErrInvalidInput
	}

	item, err := s.store.CreateTenantRelationship(ctx, db.CreateTenantRelationshipParams{
		TenantOrigemID:  origemID,
		TenantDestinoID: destinoID,
		Tipo:            req.Tipo,
		Status:          req.Status,
		Permissoes:      normalizePermissions(req.Permissoes),
		Observacao:      nullString(req.Observacao),
	})
	if err != nil {
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (Response, error) {
	item, err := s.store.GetTenantRelationshipByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return Response{}, domain.ErrNotFound
		}
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) List(ctx context.Context, tenantID uuid.UUID, pageSize int, offset int32) ([]Response, error) {
	items, err := s.store.ListTenantRelationships(ctx, db.ListTenantRelationshipsParams{
		TenantOrigemID: tenantID,
		Limit:          int32(pageSize),
		Offset:         offset,
	})
	if err != nil {
		return nil, err
	}
	out := make([]Response, 0, len(items))
	for _, item := range items {
		out = append(out, toResponse(item))
	}
	return out, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateRequest) (Response, error) {
	if req.Status != nil && !validStatus(*req.Status) {
		return Response{}, domain.ErrInvalidInput
	}
	item, err := s.store.UpdateTenantRelationship(ctx, db.UpdateTenantRelationshipParams{
		ID:         id,
		Status:     nullString(req.Status),
		Permissoes: normalizePermissions(req.Permissoes),
		Observacao: nullString(req.Observacao),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return Response{}, domain.ErrNotFound
		}
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.store.SoftDeleteTenantRelationship(ctx, id)
}

func validStatus(value string) bool {
	switch value {
	case "ativo", "inativo", "pendente", "bloqueado", "excluido":
		return true
	default:
		return false
	}
}

func validTipo(value string) bool {
	switch value {
	case "parceria", "administracao", "captacao", "repasse", "outro":
		return true
	default:
		return false
	}
}

func normalizePermissions(values []string) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func nullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: strings.TrimSpace(*value), Valid: true}
}
