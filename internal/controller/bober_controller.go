package controller

import (
	"gyad/internal/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type BoberController struct {
    BaseController
    BoberRepo repository.BobersRepository
}

func NewBoberController(repo repository.BobersRepository) *BoberController {
    return &BoberController{
        BoberRepo: repo,
    }
}

func (c *BoberController) GetAllBobers(w http.ResponseWriter, r *http.Request) {
    Bobers, err := c.BoberRepo.List(r.Context())
    if err != nil {
        c.SendJSONResponse(w, http.StatusInternalServerError, nil, err)
        return
    }
    c.SendJSONResponse(w, http.StatusOK, Bobers, nil)
}

func (c *BoberController) GetBoberByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"] 
    
    Bober, err := c.BoberRepo.FindByID(r.Context(), id)
    if err != nil {
        c.SendJSONResponse(w, http.StatusNotFound, nil, err)
        return
    }
    c.SendJSONResponse(w, http.StatusOK, Bober, nil)
}