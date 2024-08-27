package main

import (
	"github.com/ocrosby/go-reference-api/internal/handlers"
	"github.com/ocrosby/go-reference-api/internal/health"
	"net/http"
)

// @title Go Reference API
// @version 1.0
// @description This is a sample server for a Go reference API.
// @host localhost:8080
// @BasePath /

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email omar.crosby@gmail.com

func main() {
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/get", handlers.GetHandler)
	mux.HandleFunc("/put", handlers.PutHandler)
	mux.HandleFunc("/post", handlers.PostHandler)
	mux.HandleFunc("/delete", handlers.DeleteHandler)

	// Register health probes
	mux.HandleFunc("/healthz", health.LivenessProbe)
	mux.HandleFunc("/readyz", health.ReadinessProbe)

	http.ListenAndServe(":8080", mux)
}
