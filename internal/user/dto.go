package user

import (
	"database/sql"
	"time"

	db "creditoreal-crm/pkg/database/queries"
)

type CreateRequest struct {
	Nome     string  `json:"nome"`
	Email    string  `json:"email"`
	Senha    string  `json:"senha"`
	Telefone *string `json:"telefone"`
	Status   string  `json:"status"`
}

type CreateTenantUserRequest struct {
	TenantID string  `json:"tenant_id"`
	UserID   string  `json:"user_id"`
	Cargo    *string `json:"cargo"`
	Status   string  `json:"status"`
}

type Response struct {
	ID        string    `json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Telefone  *string   `json:"telefone,omitempty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TenantUserResponse struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	UserID    string    `json:"user_id"`
	Cargo     *string   `json:"cargo,omitempty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toResponse(item db.User) Response {
	return Response{
		ID:        item.ID.String(),
		Nome:      item.Nome,
		Email:     item.Email,
		Telefone:  stringPtr(item.Telefone),
		Status:    item.Status,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func toTenantUserResponse(item db.TenantUser) TenantUserResponse {
	return TenantUserResponse{
		ID:        item.ID.String(),
		TenantID:  item.TenantID.String(),
		UserID:    item.UserID.String(),
		Cargo:     stringPtr(item.Cargo),
		Status:    item.Status,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func stringPtr(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}
