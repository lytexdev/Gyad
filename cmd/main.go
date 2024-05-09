package main

import (
	"github.com/ximmanuel/Gyad/internal/database"
	"github.com/ximmanuel/Gyad/internal/server"
	
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	printASCII()

	db, err := database.NewEngine()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	serverInstance := server.StartServer(db.Engine)
	router := serverInstance.Router

	corsHandler := setupCors(router)

	URL := os.Getenv("URL")
	log.Fatal(http.ListenAndServe(URL, corsHandler))
}

// setupCors sets up the CORS middleware for the server
func setupCors(router *mux.Router) http.Handler {
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "FETCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "X-Requested-With"},
		AllowedOrigins: []string{"*"},
	})
	return corsWrapper.Handler(router)
}

// printASCII prints the ASCII logo
func printASCII() {
	content, err := ioutil.ReadFile("ascii-logo.txt")
	if err != nil {
		fmt.Println("LYTEX MEDIA")
		return
	}

	fmt.Println(string(content))
}
