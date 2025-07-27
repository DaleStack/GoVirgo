package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"govirgo/server"
)

func TestServerHealth(t *testing.T) {
	handler := server.BuildHandler()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.Code)
	}
}
