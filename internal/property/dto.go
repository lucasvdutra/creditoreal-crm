package property

import (
	"database/sql"
	"time"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type CreateRequest struct {
	PropertyTypeID       *string `json:"property_type_id"`
	Titulo               string  `json:"titulo"`
	Descricao            *string `json:"descricao"`
	Status               string  `json:"status"`
	TipoTransacao        string  `json:"tipo_transacao"`
	PrecoVendaCentavos   *int64  `json:"preco_venda_centavos"`
	PrecoAluguelCentavos *int64  `json:"preco_aluguel_centavos"`
	AreaPrivativaM2      *int32  `json:"area_privativa_m2"`
	Quartos              *int32  `json:"quartos"`
	Banheiros            *int32  `json:"banheiros"`
	VagasGaragem         *int32  `json:"vagas_garagem"`
}

type UpdateRequest struct {
	PropertyTypeID       *string `json:"property_type_id"`
	Titulo               *string `json:"titulo"`
	Descricao            *string `json:"descricao"`
	Status               *string `json:"status"`
	TipoTransacao        *string `json:"tipo_transacao"`
	PrecoVendaCentavos   *int64  `json:"preco_venda_centavos"`
	PrecoAluguelCentavos *int64  `json:"preco_aluguel_centavos"`
	AreaPrivativaM2      *int32  `json:"area_privativa_m2"`
	Quartos              *int32  `json:"quartos"`
	Banheiros            *int32  `json:"banheiros"`
	VagasGaragem         *int32  `json:"vagas_garagem"`
}

type Response struct {
	ID                   string    `json:"id"`
	TenantDonoID         string    `json:"tenant_dono_id"`
	PropertyTypeID       *string   `json:"property_type_id,omitempty"`
	Titulo               string    `json:"titulo"`
	Descricao            *string   `json:"descricao,omitempty"`
	Status               string    `json:"status"`
	TipoTransacao        string    `json:"tipo_transacao"`
	PrecoVendaCentavos   *int64    `json:"preco_venda_centavos,omitempty"`
	PrecoAluguelCentavos *int64    `json:"preco_aluguel_centavos,omitempty"`
	AreaPrivativaM2      *int32    `json:"area_privativa_m2,omitempty"`
	Quartos              *int32    `json:"quartos,omitempty"`
	Banheiros            *int32    `json:"banheiros,omitempty"`
	VagasGaragem         *int32    `json:"vagas_garagem,omitempty"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func toResponse(item db.Property) Response {
	return Response{
		ID:                   item.ID.String(),
		TenantDonoID:         item.TenantDonoID.String(),
		PropertyTypeID:       uuidPtr(item.PropertyTypeID),
		Titulo:               item.Titulo,
		Descricao:            stringPtr(item.Descricao),
		Status:               item.Status,
		TipoTransacao:        item.TipoTransacao,
		PrecoVendaCentavos:   int64Ptr(item.PrecoVendaCentavos),
		PrecoAluguelCentavos: int64Ptr(item.PrecoAluguelCentavos),
		AreaPrivativaM2:      int32Ptr(item.AreaPrivativaM2),
		Quartos:              int32Ptr(item.Quartos),
		Banheiros:            int32Ptr(item.Banheiros),
		VagasGaragem:         int32Ptr(item.VagasGaragem),
		CreatedAt:            item.CreatedAt,
		UpdatedAt:            item.UpdatedAt,
	}
}

func stringPtr(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}

func uuidPtr(value uuid.NullUUID) *string {
	if !value.Valid {
		return nil
	}
	id := value.UUID.String()
	return &id
}

func int64Ptr(value sql.NullInt64) *int64 {
	if !value.Valid {
		return nil
	}
	return &value.Int64
}

func int32Ptr(value sql.NullInt32) *int32 {
	if !value.Valid {
		return nil
	}
	return &value.Int32
}
