package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Thomazoide/donde-caigo-backend/service"
	"github.com/Thomazoide/donde-caigo-backend/structs"
)

func (s *APIServer) handleAuth(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		return s.handleLogin(w, r)
	}
	WriteJSON(w, http.StatusMethodNotAllowed, &structs.ApiResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "ONLY POST METHOD ALLOWED",
		Error:      fmt.Errorf("method not allowed: %s", r.Method),
	})
	return nil
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	var authPayload *structs.LoginPayload
	decodeErr := json.NewDecoder(r.Body).Decode(&authPayload)
	if decodeErr != nil {
		return decodeErr
	}
	authService := service.NewAuthService()
	loginErr := authService.Login(authPayload.Email, authPayload.Password)
	if loginErr != nil {
		return loginErr
	}
	return nil
}