package auth

import (
	"errors"
	"net/http"

	"creditoreal-crm/internal/domain"
	"creditoreal-crm/internal/http/respond"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/login", h.login)
	mux.HandleFunc("POST /api/auth/refresh", h.refresh)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}
	tokens, err := h.service.Login(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, tokens)
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	if err := respond.Decode(r, &req); err != nil {
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
		return
	}
	tokens, err := h.service.Refresh(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	respond.JSON(w, http.StatusOK, tokens)
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidInput):
		respond.Error(w, http.StatusBadRequest, "requisicao.invalida", "Requisicao invalida.")
	case errors.Is(err, domain.ErrInvalidCredentials):
		respond.Error(w, http.StatusUnauthorized, "auth.credenciais_invalidas", "Credenciais invalidas.")
	case errors.Is(err, domain.ErrUnauthorized):
		respond.Error(w, http.StatusUnauthorized, "permissao.negada", "Permissao negada.")
	default:
		respond.Error(w, http.StatusInternalServerError, "erro.interno", "Erro interno.")
	}
}
