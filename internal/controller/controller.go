package controller

import (
	"encoding/json"
	"net/http"
)

// BaseController is the base for all controllers
type BaseController struct{}

// SendJSONResponse sends a JSON response to the client
func (bc *BaseController) SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	} else {
		json.NewEncoder(w).Encode(data)
	}
}
