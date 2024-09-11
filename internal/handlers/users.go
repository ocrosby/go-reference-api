package handlers

import "net/http"

// CreateUser handles POST requests
// @Summary Post request
// @Description Handle POST request
// @Tags user
// @Produce plain
// @Success 200 {string} string "POST request"
// @Router /post [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User created"))
}
