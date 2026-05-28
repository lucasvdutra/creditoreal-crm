package propertyaddress

import (
	"context"
	"database/sql"
	"strings"

	"creditoreal-crm/internal/domain"
	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type Store interface {
	GetPropertyByID(context.Context, uuid.UUID) (db.Property, error)
	UpsertPropertyAddress(context.Context, db.UpsertPropertyAddressParams) (db.PropertyAddress, error)
	GetPropertyAddress(context.Context, uuid.UUID) (db.PropertyAddress, error)
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

func (s *Service) Upsert(ctx context.Context, userID, currentTenantID, propertyID uuid.UUID, req UpsertRequest) (Response, error) {
	if req.ExibicaoEndereco == "" {
		req.ExibicaoEndereco = "bairro"
	}
	if err := validate(req); err != nil {
		return Response{}, err
	}
	property, err := s.getProperty(ctx, userID, currentTenantID, propertyID, "imovel.gerenciar")
	if err != nil {
		return Response{}, err
	}
	item, err := s.store.UpsertPropertyAddress(ctx, db.UpsertPropertyAddressParams{
		PropertyID:       property.ID,
		Cep:              strings.TrimSpace(req.CEP),
		Estado:           strings.TrimSpace(req.Estado),
		Cidade:           strings.TrimSpace(req.Cidade),
		Bairro:           strings.TrimSpace(req.Bairro),
		Logradouro:       nullString(req.Logradouro),
		Numero:           nullString(req.Numero),
		Complemento:      nullString(req.Complemento),
		ExibicaoEndereco: req.ExibicaoEndereco,
	})
	if err != nil {
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) Get(ctx context.Context, userID, currentTenantID, propertyID uuid.UUID) (Response, error) {
	if _, err := s.getProperty(ctx, userID, currentTenantID, propertyID, "imovel.ler"); err != nil {
		return Response{}, err
	}
	item, err := s.store.GetPropertyAddress(ctx, propertyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return Response{}, domain.ErrNotFound
		}
		return Response{}, err
	}
	return toResponse(item), nil
}

func (s *Service) getProperty(ctx context.Context, userID, currentTenantID, propertyID uuid.UUID, permission string) (db.Property, error) {
	property, err := s.store.GetPropertyByID(ctx, propertyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.Property{}, domain.ErrNotFound
		}
		return db.Property{}, err
	}
	allowed, err := s.authorizer.Can(ctx, userID, currentTenantID, property.TenantDonoID, permission)
	if err != nil {
		return db.Property{}, err
	}
	if !allowed {
		return db.Property{}, domain.ErrUnauthorized
	}
	return property, nil
}

func validate(req UpsertRequest) error {
	if strings.TrimSpace(req.CEP) == "" || strings.TrimSpace(req.Estado) == "" || strings.TrimSpace(req.Cidade) == "" || strings.TrimSpace(req.Bairro) == "" {
		return domain.ErrInvalidInput
	}
	switch req.ExibicaoEndereco {
	case "completo", "bairro", "cidade", "oculto":
		return nil
	default:
		return domain.ErrInvalidInput
	}
}

func nullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: strings.TrimSpace(*value), Valid: true}
}
