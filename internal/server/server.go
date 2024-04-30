package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
    dbConn *sql.DB
}

func NewServer(db *sql.DB) *Server {
	fmt.Println("Creating new server")
	server := &Server{
		Router: http.NewServeMux(),
		dbConn: db,
	}
	ConfigureRoutes(server.Router)
	return server
}

func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}
