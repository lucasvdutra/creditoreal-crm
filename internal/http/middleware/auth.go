package middleware

import (
	"context"
	"net/http"
	"strings"

	"creditoreal-crm/internal/auth"
	"creditoreal-crm/internal/http/respond"
)

type authContextKey struct{}

type AuthContext struct {
	UserID   string
	TenantID string
}

func Auth(tokens *auth.TokenManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(strings.ToLower(header), "bearer ") {
			respond.Error(w, http.StatusUnauthorized, "permissao.negada", "Permissao negada.")
			return
		}
		claims, err := tokens.VerifyAccess(strings.TrimSpace(header[7:]))
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, "permissao.negada", "Permissao negada.")
			return
		}

		ctx := context.WithValue(r.Context(), authContextKey{}, AuthContext{
			UserID:   claims.Subject,
			TenantID: claims.TenantID,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthFromContext(ctx context.Context) (AuthContext, bool) {
	value, ok := ctx.Value(authContextKey{}).(AuthContext)
	return value, ok
}
