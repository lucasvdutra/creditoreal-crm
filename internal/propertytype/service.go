package propertytype

import (
	"context"

	db "creditoreal-crm/pkg/database/queries"
)

type Store interface {
	ListPropertyTypes(context.Context) ([]db.PropertyType, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) List(ctx context.Context) ([]Response, error) {
	items, err := s.store.ListPropertyTypes(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]Response, 0, len(items))
	for _, item := range items {
		out = append(out, toResponse(item))
	}
	return out, nil
}
