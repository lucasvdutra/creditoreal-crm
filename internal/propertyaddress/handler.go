package propertyaddress

import (
	"errors"
	"net/http"

	"creditoreal-crm/internal/auth"
	"creditoreal-crm/internal/domain"
	apphttp "creditoreal-crm/internal/http/middleware"
	"creditoreal-crm/internal/http/respond"

	"github.com/google/uuid"
)

type Handler struct {
	service *Service
	tokens  *auth.TokenManager
}

func NewHandler(service *Service, tokens *auth.TokenManager) *Handler {
	return &Handler{service: service, tokens: tokens}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.Handle("GET /api/properties/{id}/address", apphttp.Auth(h.tokens, http.HandlerFunc(h.get)))
	mux.Handle("PUT /api/properties/{id}/address", apphttp.Auth(h.tokens, http.HandlerFunc(h.upsert)))
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	userID, tenantID, propertyID, ok := ids(w, r)
	if !ok {
		return
	}
	item, err := h.service.Get(r.Context(), userID, tenantID, propertyID)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, item)
}

func (h *Handler) upsert(w http.ResponseWriter, r *http.Request) {
	userID, tenantID, propertyID, ok := ids(w, r)
	if !ok {
		return
	}
	var req UpsertRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}
	item, err := h.service.Upsert(r.Context(), userID, tenantID, propertyID, req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, item)
}

func ids(w http.ResponseWriter, r *http.Request) (uuid.UUID, uuid.UUID, uuid.UUID, bool) {
	authCtx, ok := apphttp.AuthFromContext(r.Context())
	if !ok {
		respond.Error(w, http.StatusUnauthorized, "permissao.negada", "Permissao negada.")
		return uuid.UUID{}, uuid.UUID{}, uuid.UUID{}, false
	}
	userID, err := uuid.Parse(authCtx.UserID)
	if err != nil {
		respond.Error(w, http.StatusUnauthorized, "permissao.negada", "Permissao negada.")
		return uuid.UUID{}, uuid.UUID{}, uuid.UUID{}, false
	}
	tenantID, err := uuid.Parse(authCtx.TenantID)
	if err != nil {
		respond.Error(w, http.StatusUnauthorized, "tenant.atual_obrigatorio", "Tenant atual obrigatorio.")
		return uuid.UUID{}, uuid.UUID{}, uuid.UUID{}, false
	}
	propertyID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "id.invalido", "ID invalido.")
		return uuid.UUID{}, uuid.UUID{}, uuid.UUID{}, false
	}
	return userID, tenantID, propertyID, true
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidInput):
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
	case errors.Is(err, domain.ErrNotFound):
		respond.Error(w, http.StatusNotFound, "endereco.nao_encontrado", "Endereco nao encontrado.")
	case errors.Is(err, domain.ErrUnauthorized):
		respond.Error(w, http.StatusForbidden, "permissao.negada", "Permissao negada.")
	default:
		respond.Error(w, http.StatusInternalServerError, "erro.interno", "Erro interno.")
	}
}
