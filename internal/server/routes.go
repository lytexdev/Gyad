package server

import (
	"gyad/internal/controller"
	"github.com/gorilla/mux"
)

// ConfigureRoutes sets up the routes for the server using Gorilla Mux
func ConfigureRoutes(router *mux.Router) {
	boberController := controller.NewBoberController()
	router.HandleFunc("/api/bober", boberController.GetBober).Methods("GET")
}
