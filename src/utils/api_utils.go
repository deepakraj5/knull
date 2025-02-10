package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, payload ResponseDto) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)

}
