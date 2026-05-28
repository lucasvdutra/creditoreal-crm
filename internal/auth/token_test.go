package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestAccessTokenRoundTrip(t *testing.T) {
	t.Parallel()

	manager := NewTokenManager("dev-secret")
	userID := uuid.New()
	tenantID := uuid.New()
	token, err := manager.AccessToken(userID, &tenantID, time.Minute)
	if err != nil {
		t.Fatalf("create token: %v", err)
	}

	claims, err := manager.VerifyAccess(token)
	if err != nil {
		t.Fatalf("verify token: %v", err)
	}
	if claims.Subject != userID.String() || claims.TenantID != tenantID.String() {
		t.Fatalf("unexpected claims: %+v", claims)
	}
}
