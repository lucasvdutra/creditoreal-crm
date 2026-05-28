package property

import (
	"context"
	"database/sql"
	"strings"

	"creditoreal-crm/internal/domain"
	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type Store interface {
	CreateProperty(context.Context, db.CreatePropertyParams) (db.Property, error)
	GetPropertyByID(context.Context, uuid.UUID) (db.Property, error)
	ListPropertiesByTenant(context.Context, db.ListPropertiesByTenantParams) ([]db.Property, error)
	CountPropertiesByTenant(context.Context, db.CountPropertiesByTenantParams) (int64, error)
	UpdateProperty(context.Context, db.UpdatePropertyParams) (db.Property, error)
	SoftDeleteProperty(context.Context, uuid.UUID) error
}

type Authorizer interface {
	Can(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, string) (bool, error)
}

type Service struct {
	store      Store
	authorizer Authorizer
}

func NewService(store Store, authorizer Authorizer) *Service {
	return &Service{store: store, authorizer: authorizer}
}

func (s *Service) Create(ctx context.Context, userID, tenantID uuid.UUID, req CreateRequest) (Response, error) {
	if req.Status == "" {
		req.Status = "rascunho"
	}
	if err := validateCreate(req); err != nil {
		return Response{}, err
	}
	if err := s.require(ctx, userID, tenantID, tenantID, "imovel.gerenciar"); err != nil {
		return Response{}, err
	}
	item, err := s.store.CreateProperty(ctx, db.CreatePropertyParams{
		TenantDonoID:         tenantID,
		PropertyTypeID:       nullUUID(req.PropertyTypeID),
		Titulo:               strings.TrimSpace(req.Titulo),
		Descricao:            nullString(req.Descricao),
		Status:               req.Status,
		TipoTransacao:        req.TipoTransacao,
		PrecoVendaCentavos:   nullInt64(req.PrecoVendaCentavos),
		PrecoAluguelCentavos: nullInt64(req.PrecoAluguelCentavos),
		AreaPrivativaM2:      nullInt32(req.AreaPrivativaM2),
		Quartos:              nullInt32(req.Quartos),
		Banheiros:            nullInt32(req.Banheiros),
		VagasGaragem:         nullInt32(req.VagasGaragem),
	})
	if err != nil {
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) List(ctx context.Context, userID, currentTenantID, resourceTenantID uuid.UUID, pageSize int, offset int32, status, tipoTransacao string, propertyTypeID *string) ([]Response, int64, error) {
	if err := s.require(ctx, userID, currentTenantID, resourceTenantID, "imovel.ler"); err != nil {
		return nil, 0, err
	}
	typeID, err := optionalUUID(propertyTypeID)
	if err != nil {
		return nil, 0, err
	}
	params := db.ListPropertiesByTenantParams{
		TenantDonoID:   resourceTenantID,
		Limit:          int32(pageSize),
		Offset:         offset,
		Status:         sql.NullString{String: status, Valid: status != ""},
		TipoTransacao:  sql.NullString{String: tipoTransacao, Valid: tipoTransacao != ""},
		PropertyTypeID: typeID,
	}
	items, err := s.store.ListPropertiesByTenant(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.store.CountPropertiesByTenant(ctx, db.CountPropertiesByTenantParams{
		TenantDonoID:   resourceTenantID,
		Status:         params.Status,
		TipoTransacao:  params.TipoTransacao,
		PropertyTypeID: typeID,
	})
	if err != nil {
		return nil, 0, err
	}
	out := make([]Response, 0, len(items))
	for _, item := range items {
		out = append(out, toResponse(item))
	}
	return out, total, nil
}

func (s *Service) Get(ctx context.Context, userID, currentTenantID, id uuid.UUID) (Response, error) {
	item, err := s.getAuthorized(ctx, userID, currentTenantID, id, "imovel.ler")
	if err != nil {
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) Update(ctx context.Context, userID, currentTenantID, id uuid.UUID, req UpdateRequest) (Response, error) {
	if err := validateUpdate(req); err != nil {
		return Response{}, err
	}
	if _, err := s.getAuthorized(ctx, userID, currentTenantID, id, "imovel.gerenciar"); err != nil {
		return Response{}, err
	}
	item, err := s.store.UpdateProperty(ctx, db.UpdatePropertyParams{
		ID:                   id,
		PropertyTypeID:       nullUUID(req.PropertyTypeID),
		Titulo:               nullString(req.Titulo),
		Descricao:            nullString(req.Descricao),
		Status:               nullString(req.Status),
		TipoTransacao:        nullString(req.TipoTransacao),
		PrecoVendaCentavos:   nullInt64(req.PrecoVendaCentavos),
		PrecoAluguelCentavos: nullInt64(req.PrecoAluguelCentavos),
		AreaPrivativaM2:      nullInt32(req.AreaPrivativaM2),
		Quartos:              nullInt32(req.Quartos),
		Banheiros:            nullInt32(req.Banheiros),
		VagasGaragem:         nullInt32(req.VagasGaragem),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return Response{}, domain.ErrNotFound
		}
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) Delete(ctx context.Context, userID, currentTenantID, id uuid.UUID) error {
	if _, err := s.getAuthorized(ctx, userID, currentTenantID, id, "imovel.gerenciar"); err != nil {
		return err
	}
	return s.store.SoftDeleteProperty(ctx, id)
}

func (s *Service) getAuthorized(ctx context.Context, userID, currentTenantID, id uuid.UUID, permission string) (db.Property, error) {
	item, err := s.store.GetPropertyByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.Property{}, domain.ErrNotFound
		}
		return db.Property{}, err
	}
	if err := s.require(ctx, userID, currentTenantID, item.TenantDonoID, permission); err != nil {
		return db.Property{}, err
	}
	return item, nil
}

func (s *Service) require(ctx context.Context, userID, currentTenantID, resourceTenantID uuid.UUID, permission string) error {
	allowed, err := s.authorizer.Can(ctx, userID, currentTenantID, resourceTenantID, permission)
	if err != nil {
		return err
	}
	if !allowed {
		return domain.ErrUnauthorized
	}
	return nil
}

func validateCreate(req CreateRequest) error {
	if strings.TrimSpace(req.Titulo) == "" || !validStatus(req.Status) || !validTipoTransacao(req.TipoTransacao) {
		return domain.ErrInvalidInput
	}
	if _, err := optionalUUID(req.PropertyTypeID); err != nil {
		return err
	}
	return nil
}

func validateUpdate(req UpdateRequest) error {
	if req.Titulo != nil && strings.TrimSpace(*req.Titulo) == "" {
		return domain.ErrInvalidInput
	}
	if req.Status != nil && !validStatus(*req.Status) {
		return domain.ErrInvalidInput
	}
	if req.TipoTransacao != nil && !validTipoTransacao(*req.TipoTransacao) {
		return domain.ErrInvalidInput
	}
	if _, err := optionalUUID(req.PropertyTypeID); err != nil {
		return err
	}
	return nil
}

func validStatus(value string) bool {
	switch value {
	case "rascunho", "ativo", "inativo", "vendido", "alugado", "excluido":
		return true
	default:
		return false
	}
}

func validTipoTransacao(value string) bool {
	switch value {
	case "venda", "aluguel", "venda_aluguel":
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

func nullInt64(value *int64) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *value, Valid: true}
}

func nullInt32(value *int32) sql.NullInt32 {
	if value == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *value, Valid: true}
}

func nullUUID(value *string) uuid.NullUUID {
	parsed, err := optionalUUID(value)
	if err != nil {
		return uuid.NullUUID{}
	}
	return parsed
}

func optionalUUID(value *string) (uuid.NullUUID, error) {
	if value == nil || *value == "" {
		return uuid.NullUUID{}, nil
	}
	id, err := uuid.Parse(*value)
	if err != nil {
		return uuid.NullUUID{}, domain.ErrInvalidInput
	}
	return uuid.NullUUID{UUID: id, Valid: true}, nil
}
