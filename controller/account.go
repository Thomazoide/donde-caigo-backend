package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Thomazoide/donde-caigo-backend/models"
	"github.com/Thomazoide/donde-caigo-backend/service"
	"github.com/Thomazoide/donde-caigo-backend/structs"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetAccount(w, r)
	}
	if r.Method == http.MethodPost {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}
	WriteJSON(w, http.StatusMethodNotAllowed, &structs.ApiResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "METHOD NOT ALLOWED",
		Error:      fmt.Errorf("method not allowed: %s", r.Method),
	})
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	userService := service.NewUserService()
	var newUser models.User
	decodeErr := json.NewDecoder(r.Body).Decode(&newUser)
	if decodeErr != nil {
		return decodeErr
	}
	result, insertErr := userService.CreateUser(newUser.Nombre, newUser.Rut, newUser.Email, newUser.ProfilePicture, newUser.ProfileDescription, newUser.Age)
	if insertErr != nil {
		return insertErr
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusCreated,
		Message:    "usuario creado",
		Data:       result,
	}
	WriteJSON(w, http.StatusCreated, response)
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
