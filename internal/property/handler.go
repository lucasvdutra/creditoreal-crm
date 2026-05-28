package property

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
	mux.Handle("POST /api/properties", apphttp.Auth(h.tokens, http.HandlerFunc(h.create)))
	mux.Handle("GET /api/properties", apphttp.Auth(h.tokens, http.HandlerFunc(h.list)))
	mux.Handle("GET /api/properties/{id}", apphttp.Auth(h.tokens, http.HandlerFunc(h.get)))
	mux.Handle("PATCH /api/properties/{id}", apphttp.Auth(h.tokens, http.HandlerFunc(h.update)))
	mux.Handle("DELETE /api/properties/{id}", apphttp.Auth(h.tokens, http.HandlerFunc(h.delete)))
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	userID, tenantID, ok := authIDs(w, r)
	if !ok {
		return
	}
	var req CreateRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}
	item, err := h.service.Create(r.Context(), userID, tenantID, req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusCreated, item)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	userID, currentTenantID, ok := authIDs(w, r)
	if !ok {
		return
	}
	resourceTenantID := currentTenantID
	if raw := r.URL.Query().Get("tenant_id"); raw != "" {
		parsed, err := uuid.Parse(raw)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, "id.invalido", "ID invalido.")
			return
		}
		resourceTenantID = parsed
	}
	page, pageSize, offset := respond.PaginationFromRequest(r)
	var typeID *string
	if raw := r.URL.Query().Get("property_type_id"); raw != "" {
		typeID = &raw
	}
	items, total, err := h.service.List(r.Context(), userID, currentTenantID, resourceTenantID, pageSize, offset, r.URL.Query().Get("status"), r.URL.Query().Get("tipo_transacao"), typeID)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, respond.ListResponse[Response]{
		Data: items,
		Pagination: respond.Pagination{
			Page:     page,
			PageSize: pageSize,
			Total:    total,
		},
	})
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	userID, tenantID, ok := authIDs(w, r)
	if !ok {
		return
	}
	id, ok := parseID(w, r.PathValue("id"))
	if !ok {
		return
	}
	item, err := h.service.Get(r.Context(), userID, tenantID, id)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, item)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	userID, tenantID, ok := authIDs(w, r)
	if !ok {
		return
	}
	id, ok := parseID(w, r.PathValue("id"))
	if !ok {
		return
	}
	var req UpdateRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}
	item, err := h.service.Update(r.Context(), userID, tenantID, id, req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, item)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	userID, tenantID, ok := authIDs(w, r)
	if !ok {
		return
	}
	id, ok := parseID(w, r.PathValue("id"))
	if !ok {
		return
	}
	if err := h.service.Delete(r.Context(), userID, tenantID, id); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func authIDs(w http.ResponseWriter, r *http.Request) (uuid.UUID, uuid.UUID, bool) {
	authCtx, ok := apphttp.AuthFromContext(r.Context())
	if !ok {
		respond.Error(w, http.StatusUnauthorized, "permissao.negada", "Permissao negada.")
		return uuid.UUID{}, uuid.UUID{}, false
	}
	userID, err := uuid.Parse(authCtx.UserID)
	if err != nil {
		respond.Error(w, http.StatusUnauthorized, "permissao.negada", "Permissao negada.")
		return uuid.UUID{}, uuid.UUID{}, false
	}
	tenantID, err := uuid.Parse(authCtx.TenantID)
	if err != nil {
		respond.Error(w, http.StatusUnauthorized, "tenant.atual_obrigatorio", "Tenant atual obrigatorio.")
		return uuid.UUID{}, uuid.UUID{}, false
	}
	return userID, tenantID, true
}

func parseID(w http.ResponseWriter, raw string) (uuid.UUID, bool) {
	id, err := uuid.Parse(raw)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "id.invalido", "ID invalido.")
		return uuid.UUID{}, false
	}
	return id, true
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidInput):
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
	case errors.Is(err, domain.ErrNotFound):
		respond.Error(w, http.StatusNotFound, "imovel.nao_encontrado", "Imovel nao encontrado.")
	case errors.Is(err, domain.ErrUnauthorized):
		respond.Error(w, http.StatusForbidden, "permissao.negada", "Permissao negada.")
	default:
		respond.Error(w, http.StatusInternalServerError, "erro.interno", "Erro interno.")
	}
}
