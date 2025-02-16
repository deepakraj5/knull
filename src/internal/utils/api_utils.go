package utils

import (
	"encoding/json"
	"knull/internal/dtos"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, payload dtos.ResponseDto) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)

}
