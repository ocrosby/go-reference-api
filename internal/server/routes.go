package server

import (
	"github.com/ocrosby/go-reference-api/internal/handlers"
	"github.com/ocrosby/go-reference-api/internal/health"
	"github.com/ocrosby/go-reference-api/pkg/middleware"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	// @title Example API
	// @description This is a sample server Petstore server.
	// @version 1.0
	// @termsOfService http://swagger.io/terms/
	// @contact

	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("GET /get", handlers.GetHandler)
	mux.HandleFunc("PUT /put", handlers.PutHandler)
	mux.HandleFunc("POST /post", handlers.PostHandler)
	mux.HandleFunc("DELETE /delete", handlers.DeleteHandler)

	mux.HandleFunc("POST /user", handlers.CreateUser)

	// Register health probes
	mux.Handle("/healthz", middleware.LoggingMiddleware(http.HandlerFunc(health.LivenessProbe)))
	mux.Handle("/readyz", middleware.LoggingMiddleware(http.HandlerFunc(health.ReadinessProbe)))

	return mux
}
