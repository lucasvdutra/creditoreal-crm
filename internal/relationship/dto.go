package relationship

import (
	"database/sql"
	"time"

	db "creditoreal-crm/pkg/database/queries"
)

type CreateRequest struct {
	TenantOrigemID  string   `json:"tenant_origem_id"`
	TenantDestinoID string   `json:"tenant_destino_id"`
	Tipo            string   `json:"tipo"`
	Status          string   `json:"status"`
	Permissoes      []string `json:"permissoes"`
	Observacao      *string  `json:"observacao"`
}

type UpdateRequest struct {
	Status     *string  `json:"status"`
	Permissoes []string `json:"permissoes"`
	Observacao *string  `json:"observacao"`
}

type Response struct {
	ID              string    `json:"id"`
	TenantOrigemID  string    `json:"tenant_origem_id"`
	TenantDestinoID string    `json:"tenant_destino_id"`
	Tipo            string    `json:"tipo"`
	Status          string    `json:"status"`
	Permissoes      []string  `json:"permissoes"`
	Observacao      *string   `json:"observacao,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func toResponse(item db.TenantRelationship) Response {
	return Response{
		ID:              item.ID.String(),
		TenantOrigemID:  item.TenantOrigemID.String(),
		TenantDestinoID: item.TenantDestinoID.String(),
		Tipo:            item.Tipo,
		Status:          item.Status,
		Permissoes:      item.Permissoes,
		Observacao:      stringPtr(item.Observacao),
		CreatedAt:       item.CreatedAt,
		UpdatedAt:       item.UpdatedAt,
	}
}

func stringPtr(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}
