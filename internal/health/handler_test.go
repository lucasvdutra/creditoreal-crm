package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"creditoreal-crm/internal/http/middleware"
)

func TestHealthcheck(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", Handle)

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	middleware.RequestID(mux).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if response.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected content type application/json, got %q", response.Header().Get("Content-Type"))
	}

	if response.Header().Get("X-Request-Id") == "" {
		t.Fatal("expected request id header")
	}

	var body Response
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response body: %v", err)
	}

	if body.Status != "saudavel" {
		t.Fatalf("expected healthy status, got %q", body.Status)
	}

	if body.Time.IsZero() {
		t.Fatal("expected response time")
	}
}
