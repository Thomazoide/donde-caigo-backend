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
		Error:      fmt.Errorf("method not allowed: %s", r.Method).Error(),
	})
	return nil
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	var authPayload *structs.LoginPayload
	decodeErr := json.NewDecoder(r.Body).Decode(&authPayload)
	fmt.Println(authPayload)
	if decodeErr != nil {
		return decodeErr
	}
	authService := service.NewAuthService()
	token, userData, loginErr := authService.Login(authPayload.Email, authPayload.Password)
	if loginErr != nil {
		return loginErr
	}
	authService.SignCookie(w, *token)
	response := &structs.LoginResponse{
		StatusCode: http.StatusAccepted,
		Message:    "Sesion iniciada",
		Token:      *token,
		UserData:   *userData,
	}
	WriteJSON(w, http.StatusAccepted, response)
	return nil
}
