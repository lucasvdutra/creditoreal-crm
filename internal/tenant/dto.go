package tenant

import (
	"database/sql"
	"encoding/json"
	"time"

	db "creditoreal-crm/pkg/database/queries"
)

type CreateRequest struct {
	Nome            string          `json:"nome"`
	NomeFantasia    *string         `json:"nome_fantasia"`
	TipoDocumento   *string         `json:"tipo_documento"`
	NumeroDocumento *string         `json:"numero_documento"`
	Email           *string         `json:"email"`
	Telefone        *string         `json:"telefone"`
	Site            *string         `json:"site"`
	LogoURL         *string         `json:"logo_url"`
	Status          string          `json:"status"`
	Tipo            string          `json:"tipo"`
	Configuracoes   json.RawMessage `json:"configuracoes"`
}

type UpdateRequest struct {
	Nome            *string         `json:"nome"`
	NomeFantasia    *string         `json:"nome_fantasia"`
	TipoDocumento   *string         `json:"tipo_documento"`
	NumeroDocumento *string         `json:"numero_documento"`
	Email           *string         `json:"email"`
	Telefone        *string         `json:"telefone"`
	Site            *string         `json:"site"`
	LogoURL         *string         `json:"logo_url"`
	Status          *string         `json:"status"`
	Tipo            *string         `json:"tipo"`
	Configuracoes   json.RawMessage `json:"configuracoes"`
}

type Response struct {
	ID              string          `json:"id"`
	Nome            string          `json:"nome"`
	NomeFantasia    *string         `json:"nome_fantasia,omitempty"`
	TipoDocumento   *string         `json:"tipo_documento,omitempty"`
	NumeroDocumento *string         `json:"numero_documento,omitempty"`
	Email           *string         `json:"email,omitempty"`
	Telefone        *string         `json:"telefone,omitempty"`
	Site            *string         `json:"site,omitempty"`
	LogoURL         *string         `json:"logo_url,omitempty"`
	Status          string          `json:"status"`
	Tipo            string          `json:"tipo"`
	Configuracoes   json.RawMessage `json:"configuracoes"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

func toResponse(item db.Tenant) Response {
	return Response{
		ID:              item.ID.String(),
		Nome:            item.Nome,
		NomeFantasia:    stringPtr(item.NomeFantasia),
		TipoDocumento:   stringPtr(item.TipoDocumento),
		NumeroDocumento: stringPtr(item.NumeroDocumento),
		Email:           stringPtr(item.Email),
		Telefone:        stringPtr(item.Telefone),
		Site:            stringPtr(item.Site),
		LogoURL:         stringPtr(item.LogoUrl),
		Status:          item.Status,
		Tipo:            item.Tipo,
		Configuracoes:   item.Configuracoes,
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
