package controller

import (
    "net/http"
)

type TestController struct {
    BaseController
}

func NewTestController() *TestController {
    return &TestController{}
}

func (test *TestController) GetTest(w http.ResponseWriter, r *http.Request) {
    userData := map[string]string{
		"success": "Hello World!",
		"message": "This is a test message.",
		"authors": "LYTEX MEDIA",
	} 
    test.SendJSONResponse(w, http.StatusOK, userData, nil)
}
