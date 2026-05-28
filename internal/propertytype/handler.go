package propertytype

import (
	"net/http"

	"creditoreal-crm/internal/http/respond"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/property-types", h.list)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.List(r.Context())
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "erro.interno", "Erro interno.")
		return
	}
	respond.JSON(w, http.StatusOK, items)
}
