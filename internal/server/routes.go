package server

import (
	"gyad/internal/controller"
	"gyad/internal/repository"
	"github.com/gorilla/mux"
)

func ConfigureRoutes(router *mux.Router, factory *repository.RepositoryFactory) {
    boberController := controller.NewBoberController(factory.CreateBoberRepository())
    router.HandleFunc("/api/bobers", boberController.GetAllBobers).Methods("GET")
	router.HandleFunc("/api/bober/{id}", boberController.GetBoberByID).Methods("GET")

}
