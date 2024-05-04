package main

import (
	"gyad/internal/database"
	"gyad/internal/repository"
	"gyad/internal/server"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	printASCCI()

	db := database.NewDatabase()
	db.Connect()
	defer db.Close()

	repoFactory := repository.NewRepositoryFactory(db)
	serverInstance := server.StartServer(db.Conn, repoFactory)
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

// printASCCI prints the ASCII logo from lytexmedia
func printASCCI() {
	 content, err := ioutil.ReadFile("ascii-logo.txt")
	 if err != nil {
		 fmt.Println("LYTEX MEDIA")
		 return
	 }

	 fmt.Println(string(content))
}