package tenant

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
	mux.HandleFunc("POST /api/tenants", h.create)
	mux.HandleFunc("GET /api/tenants", h.list)
	mux.HandleFunc("GET /api/tenants/{id}", h.get)
	mux.HandleFunc("PATCH /api/tenants/{id}", h.update)
	mux.HandleFunc("DELETE /api/tenants/{id}", h.delete)
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
	items, total, err := h.service.List(r.Context(), pageSize, offset, r.URL.Query().Get("status"), r.URL.Query().Get("tipo"))
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
	id, ok := parseID(w, r.PathValue("id"))
	if !ok {
		return
	}

	item, err := h.service.Get(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, item)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r.PathValue("id"))
	if !ok {
		return
	}

	var req UpdateRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}

	item, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, item)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r.PathValue("id"))
	if !ok {
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
		respond.Error(w, http.StatusNotFound, "tenant.nao_encontrado", "Tenant nao encontrado.")
	default:
		respond.Error(w, http.StatusInternalServerError, "erro.interno", "Erro interno.")
	}
}
