package propertytype

import (
	"time"

	db "creditoreal-crm/pkg/database/queries"
)

type Response struct {
	ID        string    `json:"id"`
	Codigo    string    `json:"codigo"`
	Nome      string    `json:"nome"`
	Categoria string    `json:"categoria"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toResponse(item db.PropertyType) Response {
	return Response{
		ID:        item.ID.String(),
		Codigo:    item.Codigo,
		Nome:      item.Nome,
		Categoria: item.Categoria,
		Status:    item.Status,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}
