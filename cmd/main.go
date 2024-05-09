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
	"golang.org/x/crypto/acme/autocert"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	printASCII()

	db, err := database.NewEngine()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	serverInstance := server.StartServer(db.Engine)
	router := serverInstance.Router

	corsHandler := setupCors(router)

	sslEnabled := os.Getenv("SSL_ENABLED")
	domain := os.Getenv("DOMAIN")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	serverAddr := ":" + port

	if sslEnabled == "true" {
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("certs"),
			HostPolicy: autocert.HostWhitelist(domain),
		}

		httpsServer := &http.Server{
			Addr:    ":https",
			Handler: certManager.HTTPHandler(corsHandler),
		}

		go func() {
			http.ListenAndServe(":http", certManager.HTTPHandler(nil))
		}()

		log.Fatal(httpsServer.ListenAndServeTLS("", ""))
	} else {
		log.Fatal(http.ListenAndServe(serverAddr, corsHandler))
	}
}

// setupCors sets up the CORS middleware for the server
func setupCors(router *mux.Router) http.Handler {
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "FETCH", "DELETE", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-Requested-With", "Authorization", "Accept", "Origin", "X-CSRF-Token"},
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
