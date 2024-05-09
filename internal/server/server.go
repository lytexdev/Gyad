package server

import (
	"fmt"
	"xorm.io/xorm"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Engine *xorm.Engine
}

// StartServer creates a new server instance
func StartServer(engine *xorm.Engine) *Server {
	fmt.Println("Starting HTTP server...")
	router := mux.NewRouter()
	server := &Server{
		Router: router,
		Engine: engine,
	}
	ConfigureRoutes(router, engine)

	fmt.Println("Server started successfully!")
	return server
}
