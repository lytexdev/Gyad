package server

import (
	"database/sql"
	"fmt"
	"gyad/internal/repository"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	dbConn *sql.DB
}

func StartServer(db *sql.DB, repoFactory *repository.RepositoryFactory) *Server {
	fmt.Println("Starting HTTP server...")
	router := mux.NewRouter()
	server := &Server{
		Router: router,
		dbConn: db,
	}
	ConfigureRoutes(router, repoFactory)

	fmt.Println("Server started successfully!")
	return server
}
