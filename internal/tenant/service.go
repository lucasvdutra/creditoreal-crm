package tenant

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"creditoreal-crm/internal/domain"
	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

type Store interface {
	CreateTenant(context.Context, db.CreateTenantParams) (db.Tenant, error)
	GetTenantByID(context.Context, uuid.UUID) (db.Tenant, error)
	ListTenants(context.Context, db.ListTenantsParams) ([]db.Tenant, error)
	CountTenants(context.Context, db.CountTenantsParams) (int64, error)
	UpdateTenant(context.Context, db.UpdateTenantParams) (db.Tenant, error)
	SoftDeleteTenant(context.Context, uuid.UUID) error
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

	config := req.Configuracoes
	if len(config) == 0 {
		config = json.RawMessage(`{}`)
	}

	item, err := s.store.CreateTenant(ctx, db.CreateTenantParams{
		Nome:            strings.TrimSpace(req.Nome),
		NomeFantasia:    nullString(req.NomeFantasia),
		TipoDocumento:   nullString(req.TipoDocumento),
		NumeroDocumento: nullString(req.NumeroDocumento),
		Email:           nullString(req.Email),
		Telefone:        nullString(req.Telefone),
		Site:            nullString(req.Site),
		LogoUrl:         nullString(req.LogoURL),
		Status:          req.Status,
		Tipo:            req.Tipo,
		Column11:        config,
	})
	if err != nil {
		return Response{}, err
	}

	return toResponse(item), nil
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (Response, error) {
	item, err := s.store.GetTenantByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return Response{}, domain.ErrNotFound
		}
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) List(ctx context.Context, pageSize int, offset int32, status, tipo string) ([]Response, int64, error) {
	params := db.ListTenantsParams{
		Limit:  int32(pageSize),
		Offset: offset,
		Status: sql.NullString{String: status, Valid: status != ""},
		Tipo:   sql.NullString{String: tipo, Valid: tipo != ""},
	}
	items, err := s.store.ListTenants(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.store.CountTenants(ctx, db.CountTenantsParams{Status: params.Status, Tipo: params.Tipo})
	if err != nil {
		return nil, 0, err
	}
	out := make([]Response, 0, len(items))
	for _, item := range items {
		out = append(out, toResponse(item))
	}
	return out, total, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateRequest) (Response, error) {
	if err := validateUpdate(req); err != nil {
		return Response{}, err
	}

	item, err := s.store.UpdateTenant(ctx, db.UpdateTenantParams{
		ID:              id,
		Nome:            nullString(req.Nome),
		NomeFantasia:    nullString(req.NomeFantasia),
		TipoDocumento:   nullString(req.TipoDocumento),
		NumeroDocumento: nullString(req.NumeroDocumento),
		Email:           nullString(req.Email),
		Telefone:        nullString(req.Telefone),
		Site:            nullString(req.Site),
		LogoUrl:         nullString(req.LogoURL),
		Status:          nullString(req.Status),
		Tipo:            nullString(req.Tipo),
		Configuracoes:   nullRaw(req.Configuracoes),
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
	return s.store.SoftDeleteTenant(ctx, id)
}

func validateCreate(req CreateRequest) error {
	if strings.TrimSpace(req.Nome) == "" || req.Tipo == "" {
		return domain.ErrInvalidInput
	}
	if !validStatus(req.Status) || !validTipo(req.Tipo) {
		return domain.ErrInvalidInput
	}
	if req.TipoDocumento != nil && !validTipoDocumento(*req.TipoDocumento) {
		return domain.ErrInvalidInput
	}
	if len(req.Configuracoes) > 0 && !json.Valid(req.Configuracoes) {
		return domain.ErrInvalidInput
	}
	return nil
}

func validateUpdate(req UpdateRequest) error {
	if req.Nome != nil && strings.TrimSpace(*req.Nome) == "" {
		return domain.ErrInvalidInput
	}
	if req.Status != nil && !validStatus(*req.Status) {
		return domain.ErrInvalidInput
	}
	if req.Tipo != nil && !validTipo(*req.Tipo) {
		return domain.ErrInvalidInput
	}
	if req.TipoDocumento != nil && !validTipoDocumento(*req.TipoDocumento) {
		return domain.ErrInvalidInput
	}
	if len(req.Configuracoes) > 0 && !json.Valid(req.Configuracoes) {
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

func validTipo(value string) bool {
	switch value {
	case "imobiliaria", "incorporadora", "corretor_autonomo", "administradora", "parceiro":
		return true
	default:
		return false
	}
}

func validTipoDocumento(value string) bool {
	return value == "cpf" || value == "cnpj"
}

func nullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: strings.TrimSpace(*value), Valid: true}
}

func nullRaw(value json.RawMessage) pqtype.NullRawMessage {
	if len(value) == 0 {
		return pqtype.NullRawMessage{}
	}
	return pqtype.NullRawMessage{RawMessage: value, Valid: true}
}
