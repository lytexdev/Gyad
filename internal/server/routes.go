package server

import (
	"net/http"

	"gyad/internal/controller"
)

// ConfigureRoutes sets up the routes for the server
func ConfigureRoutes(mux *http.ServeMux) {
	boberController := controller.BoberController{}
	mux.HandleFunc("/api/bober", boberController.GetBober)
}
