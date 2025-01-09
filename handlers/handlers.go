package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Thomazoide/donde-caigo-backend/structs"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		response := &structs.ApiResponse{
			StatusCode: 200,
			Message:    "OK",
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := &structs.ApiResponse{
			StatusCode: 400,
			Message:    "ERROR",
		}
		json.NewEncoder(w).Encode(response)
	}
}
