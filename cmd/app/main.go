package main

import (
	"fmt"
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
	mux.HandleFunc("GET /get", handlers.GetHandler)
	mux.HandleFunc("PUT /put", handlers.PutHandler)
	mux.HandleFunc("POST /post", handlers.PostHandler)
	mux.HandleFunc("DELETE /delete", handlers.DeleteHandler)

	// Register health probes
	mux.HandleFunc("/healthz", health.LivenessProbe)
	mux.HandleFunc("/readyz", health.ReadinessProbe)

	fmt.Println("Server listening on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}
