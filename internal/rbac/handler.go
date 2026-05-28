package rbac

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
	mux.HandleFunc("POST /api/rbac/roles", h.createRole)
	mux.HandleFunc("GET /api/rbac/roles", h.listRoles)
	mux.HandleFunc("GET /api/rbac/permissions", h.listPermissions)
	mux.HandleFunc("POST /api/rbac/role-permissions", h.addPermission)
	mux.HandleFunc("POST /api/rbac/assign-role", h.assignRole)
	mux.HandleFunc("GET /api/rbac/user-permissions", h.listUserPermissions)
}

func (h *Handler) createRole(w http.ResponseWriter, r *http.Request) {
	var req CreateRoleRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}
	item, err := h.service.CreateRole(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusCreated, item)
}

func (h *Handler) listRoles(w http.ResponseWriter, r *http.Request) {
	tenantID, err := uuid.Parse(r.URL.Query().Get("tenant_id"))
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "id.invalido", "ID invalido.")
		return
	}
	items, err := h.service.ListRoles(r.Context(), tenantID)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, items)
}

func (h *Handler) listPermissions(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.ListPermissions(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, items)
}

func (h *Handler) addPermission(w http.ResponseWriter, r *http.Request) {
	var req AddPermissionRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}
	if err := h.service.AddPermission(r.Context(), req); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) assignRole(w http.ResponseWriter, r *http.Request) {
	var req AssignRoleRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}
	if err := h.service.AssignRole(r.Context(), req); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) listUserPermissions(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "id.invalido", "ID invalido.")
		return
	}
	tenantID, err := uuid.Parse(r.URL.Query().Get("tenant_id"))
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "id.invalido", "ID invalido.")
		return
	}
	items, err := h.service.ListUserPermissions(r.Context(), userID, tenantID)
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
