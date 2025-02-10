package handlers

import (
	"knull/utils"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	payload := utils.ResponseDto{
		ResponseCode: 200,
		Message:      "Signup successful",
		Data:         "user data",
	}

	utils.JsonResponse(w, payload)
}
