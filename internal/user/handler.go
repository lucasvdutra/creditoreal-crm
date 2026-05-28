package user

import (
	"errors"
	"net/http"

	"creditoreal-crm/internal/domain"
	"creditoreal-crm/internal/http/respond"

	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", h.create)
	mux.HandleFunc("GET /api/users", h.list)
	mux.HandleFunc("POST /api/tenant-users", h.createTenantUser)
	mux.HandleFunc("GET /api/tenants/{tenant_id}/users", h.listTenantUsers)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}

	item, err := h.service.Create(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusCreated, item)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	page, pageSize, offset := respond.PaginationFromRequest(r)
	items, total, err := h.service.List(r.Context(), pageSize, offset, r.URL.Query().Get("status"))
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

func (h *Handler) createTenantUser(w http.ResponseWriter, r *http.Request) {
	var req CreateTenantUserRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}

	item, err := h.service.CreateTenantUser(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusCreated, item)
}

func (h *Handler) listTenantUsers(w http.ResponseWriter, r *http.Request) {
	tenantID, err := uuid.Parse(r.PathValue("tenant_id"))
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "id.invalido", "ID invalido.")
		return
	}
	_, pageSize, offset := respond.PaginationFromRequest(r)
	items, err := h.service.ListTenantUsers(r.Context(), tenantID, pageSize, offset)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, items)
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidInput):
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
	default:
		respond.Error(w, http.StatusInternalServerError, "erro.interno", "Erro interno.")
	}
}
