package handlers

import (
	"fmt"
	"net/http"
)

// GetHandler handles GET requests
// @Summary Get request
// @Description Handle GET request
// @Tags example
// @Produce plain
// @Success 200 {string} string "GET request"
// @Router /get [get]
func GetHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprint(w, "GET request")
	_, _ = w.Write([]byte("GET request"))
}

// PutHandler handles PUT requests
// @Summary Put request
// @Description Handle PUT request
// @Tags example
// @Produce plain
// @Success 200 {string} string "PUT request"
// @Router /put [put]
func PutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("PUT request"))
}

// PostHandler handles POST requests
// @Summary Post request
// @Description Handle POST request
// @Tags example
// @Produce plain
// @Success 200 {string} string "POST request"
// @Router /post [post]
func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("POST request"))
}

// DeleteHandler handles DELETE requests
// @Summary Delete request
// @Description Handle DELETE request
// @Tags example
// @Produce plain
// @Success 200 {string} string "DELETE request"
// @Router /delete [delete]
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("DELETE request"))
}
