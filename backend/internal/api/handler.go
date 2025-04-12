package api

import (
	"net/http"
)

// HealthHandler responds with "OK" when called
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK")) 
}
