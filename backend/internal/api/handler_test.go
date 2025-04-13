package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	HealthHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	// Assert the status code is 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
