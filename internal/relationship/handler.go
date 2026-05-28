package relationship

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
	mux.HandleFunc("POST /api/tenant-relationships", h.create)
	mux.HandleFunc("GET /api/tenant-relationships", h.list)
	mux.HandleFunc("GET /api/tenant-relationships/{id}", h.get)
	mux.HandleFunc("PATCH /api/tenant-relationships/{id}", h.update)
	mux.HandleFunc("DELETE /api/tenant-relationships/{id}", h.delete)
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
	tenantID, err := uuid.Parse(r.URL.Query().Get("tenant_id"))
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "id.invalido", "ID invalido.")
		return
	}
	_, pageSize, offset := respond.PaginationFromRequest(r)
	items, err := h.service.List(r.Context(), tenantID, pageSize, offset)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, items)
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
		respond.Error(w, http.StatusNotFound, "relacionamento.nao_encontrado", "Relacionamento nao encontrado.")
	default:
		respond.Error(w, http.StatusInternalServerError, "erro.interno", "Erro interno.")
	}
}
