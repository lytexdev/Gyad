package server

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	dbConn *sql.DB
}

func StartServer(db *sql.DB) *Server {
	fmt.Println("Starting HTTP server...")
	router := mux.NewRouter()
	server := &Server{
		Router: router,
		dbConn: db,
	}
	ConfigureRoutes(router)

	fmt.Println("Server started successfully!")
	return server
}
