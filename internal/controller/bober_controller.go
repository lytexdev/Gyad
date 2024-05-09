package controller

import (
	"github.com/ximmanuel/Gyad/models"

	"log"
	"net/http"

	"xorm.io/xorm"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type BoberController struct {
	BaseController
	Engine *xorm.Engine
}

func NewBoberController(engine *xorm.Engine) *BoberController {
	return &BoberController{
		Engine: engine,
	}
}

func (uc *BoberController) GetAllBobers(w http.ResponseWriter, r *http.Request) {
	var bober []models.Bober
	err := uc.Engine.Find(&bober)
	if err != nil {
		uc.SendJSONResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	uc.SendJSONResponse(w, http.StatusOK, bober, nil)
}

func (uc *BoberController) GetBoberByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	bober := &models.Bober{}
	has, err := uc.Engine.ID(id).Get(bober)
	if err != nil {
		uc.SendJSONResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	if !has {
		uc.SendJSONResponse(w, http.StatusNotFound, nil, nil)
		return
	}

	uc.SendJSONResponse(w, http.StatusOK, bober, nil)
}

func (uc *BoberController) CreateTestBober(w http.ResponseWriter, r *http.Request) {
	bober := &models.Bober{
		ID:   uuid.New().String(),
		Name: "Boberino",
		Age:  3,
	}
	log.Printf("Attempting to create bober: %+v", bober)
	_, err := uc.Engine.Insert(bober)
	if err != nil {
		log.Printf("Error creating bober: %v", err)
		uc.SendJSONResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	uc.SendJSONResponse(w, http.StatusCreated, bober, nil)
}
