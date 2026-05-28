package propertyaddress

import (
	"database/sql"
	"time"

	db "creditoreal-crm/pkg/database/queries"
)

type UpsertRequest struct {
	CEP              string  `json:"cep"`
	Estado           string  `json:"estado"`
	Cidade           string  `json:"cidade"`
	Bairro           string  `json:"bairro"`
	Logradouro       *string `json:"logradouro"`
	Numero           *string `json:"numero"`
	Complemento      *string `json:"complemento"`
	ExibicaoEndereco string  `json:"exibicao_endereco"`
}

type Response struct {
	ID               string    `json:"id"`
	PropertyID       string    `json:"property_id"`
	CEP              string    `json:"cep"`
	Estado           string    `json:"estado"`
	Cidade           string    `json:"cidade"`
	Bairro           string    `json:"bairro"`
	Logradouro       *string   `json:"logradouro,omitempty"`
	Numero           *string   `json:"numero,omitempty"`
	Complemento      *string   `json:"complemento,omitempty"`
	ExibicaoEndereco string    `json:"exibicao_endereco"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func toResponse(item db.PropertyAddress) Response {
	return Response{
		ID:               item.ID.String(),
		PropertyID:       item.PropertyID.String(),
		CEP:              item.Cep,
		Estado:           item.Estado,
		Cidade:           item.Cidade,
		Bairro:           item.Bairro,
		Logradouro:       stringPtr(item.Logradouro),
		Numero:           stringPtr(item.Numero),
		Complemento:      stringPtr(item.Complemento),
		ExibicaoEndereco: item.ExibicaoEndereco,
		CreatedAt:        item.CreatedAt,
		UpdatedAt:        item.UpdatedAt,
	}
}

func stringPtr(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}
