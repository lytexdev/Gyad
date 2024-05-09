package server

import (
	"github.com/ximmanuel/Gyad/internal/controller"

	"github.com/gorilla/mux"
	"xorm.io/xorm"
)

// ConfigureRoutes sets up the routes for the server
func ConfigureRoutes(router *mux.Router, engine *xorm.Engine) {
	router.Use(RateLimit)

	//* Bober example routes
	boberController := controller.NewBoberController(engine)
	router.HandleFunc("/api/bober", boberController.GetAllBobers).Methods("GET")
	router.HandleFunc("/api/bober/{id}", boberController.GetBoberByID).Methods("GET")
	router.HandleFunc("/api/bober/create", boberController.CreateTestBober).Methods("POST")
}
