package main

import (
	"gyad/internal/database"
	"gyad/internal/server"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	db := database.NewDatabase()
	db.Connect()
	defer db.Close()

	serverInstance := server.NewServer(db.Conn)
	corsHandler := setupCors(serverInstance.Router)
	URL := os.Getenv("URL")

	log.Fatal(http.ListenAndServe(URL, corsHandler))
}

func setupCors(router *http.ServeMux) http.Handler {
	corsOptions := cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "X-Requested-With"},
		AllowedOrigins: []string{"*"},
	}
	return cors.New(corsOptions).Handler(router)
}
