package controller

import (
	"net/http"
)

type BoberController struct {
    BaseController
}

func NewBoberController() *BoberController {
    return &BoberController{}
}

func (bc *BoberController) GetBober(w http.ResponseWriter, r *http.Request) {
    bc.SendJSONResponse(w, http.StatusOK, "bobers", nil)
}
