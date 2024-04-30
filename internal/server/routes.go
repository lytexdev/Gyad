package server

import (
	"net/http"

	"gyad/internal/controller"
)

// ConfigureRoutes sets up the routes for the server
func ConfigureRoutes(mux *http.ServeMux) {
	testController := controller.NewTestController()
	mux.HandleFunc("/api/test", testController.GetTest)
}
