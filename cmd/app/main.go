package main

import (
	"fmt"
	"github.com/ocrosby/go-reference-api/internal"
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
	mux := internal.SetupRoutes()

	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
