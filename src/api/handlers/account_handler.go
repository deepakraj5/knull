package handlers

import (
	"knull/internal/dtos"
	"knull/internal/utils"
	"knull/necrosword"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	// user := entities.User{}

	// db.DB().Create(&user)

	necrosword.Execute()

	payload := dtos.ResponseDto{
		ResponseCode: 200,
		Message:      "Signup successful",
		Data:         "user data",
	}

	utils.JsonResponse(w, payload)
}
