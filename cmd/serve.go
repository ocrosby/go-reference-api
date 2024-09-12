package cmd

import (
	"fmt"
	"github.com/ocrosby/go-reference-api/internal/server"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server on :8080")
		mux := server.SetupRoutes()
		log.Fatal(http.ListenAndServe(":8080", mux))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
