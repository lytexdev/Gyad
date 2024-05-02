package controller

import (
	"fmt"
	"net/http"

	"gyad/repository"
)

type BoberController struct {
    BaseController
    Repository repository.BoberRepository
}

func NewBoberController(repository repository.BoberRepository) *BoberController {
    return &BoberController{
        Repository: repository,
    }
}

func (test *BoberController) GetBober(w http.ResponseWriter, r *http.Request) {
    bobers, err := test.Repository.GetAll()

    fmt.Println(bobers)

    if err != nil {
        test.SendJSONResponse(w, http.StatusInternalServerError, nil, err)
        return
    }
    
    test.SendJSONResponse(w, http.StatusOK, "bobers", nil)
}
