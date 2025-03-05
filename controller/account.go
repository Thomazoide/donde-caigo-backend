package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Thomazoide/donde-caigo-backend/models"
	"github.com/Thomazoide/donde-caigo-backend/service"
	"github.com/Thomazoide/donde-caigo-backend/structs"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetAccounts(w, r)
	}
	if r.Method == http.MethodPost {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}
	if r.Method == http.MethodPut {
		return s.handleUpdateUser(w, r)
	}
	WriteJSON(w, http.StatusMethodNotAllowed, &structs.ApiResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "METHOD NOT ALLOWED",
		Error:      fmt.Errorf("method not allowed: %s", r.Method).Error(),
	})
	return nil
}

func (s *APIServer) handleAccountWithParams(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetAccountByID(w, r)
	}
	if r.Method == http.MethodPost {
		return s.handleUpdatePassword(w, r)
	}
	WriteJSON(w, http.StatusMethodNotAllowed, &structs.ApiResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "METHOD NOT ALLOWED",
		Error:      fmt.Errorf("method not allowed: %s", r.Method).Error(),
	})
	return fmt.Errorf("method not allowed: %s", r.Method)
}

// HandleGetAccounts se encarga de entregar todas las cuentas de usuario
// @Summary Entrega todas las cuentas de usuario
// @Tags Usuarios
// @Produce json
// @Success 200 {object} structs.ApiResponse
// @Router /cuenta [get]
func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	userService := service.NewUserService()
	userList, err := userService.GetAllUsers()
	if err != nil {
		return err
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "lista de usuarios",
		Data:       userList,
	}
	fmt.Println(r.Method)
	WriteJSON(w, http.StatusOK, response)
	return nil
}

// HandleGetAccountByID se encarga de entregar una cuenta de usuario por su id
// @Summary Entrega una cuenta de usuario por su id
// @Tags Usuarios
// @Produce json
// @Param id path int true "ID del usuario"
// @Success 200 {object} structs.ApiResponse
// @Router /cuenta/{id} [get]
func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	userService := service.NewUserService()
	vars := mux.Vars(r)
	id := vars["id"]
	uid, parseErr := strconv.ParseUint(id, 10, 32)
	if parseErr != nil {
		return parseErr
	}
	usr, getUsrErr := userService.GetUserByID(uint(uid))
	if getUsrErr != nil {
		return getUsrErr
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "Usuario encontrado",
		Data:       usr,
	}
	WriteJSON(w, http.StatusOK, response)
	return nil
}

// HandleCreateAccount se encarga de crear una cuenta de usuario
// @Summary Crea una cuenta de usuario
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param user body models.UserSchema true "Usuario a crear"
// @Success 201 {object} structs.ApiResponse
// @Router /cuenta [post]
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	userService := service.NewUserService()
	var newUser models.User
	decodeErr := json.NewDecoder(r.Body).Decode(&newUser)
	if decodeErr != nil {
		return decodeErr
	}
	result, insertErr := userService.CreateUser(newUser.Nombre, newUser.Password, newUser.Rut, newUser.Email, newUser.ProfilePicture, newUser.ProfileDescription, newUser.Age)
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

// HandleDeleteAccount se encarga de eliminar una cuenta de usuario
// @Summary Elimina una cuenta de usuario
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param user body models.UserSchema true "Usuario a eliminar"
// @Success 200 {object} structs.ApiResponse
// @Router /cuenta [delete]
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	userService := service.NewUserService()
	var userToDelete models.User
	decodeErr := json.NewDecoder(r.Body).Decode(&userToDelete)
	if decodeErr != nil {
		return decodeErr
	}
	deleteErr := userService.DeleteUser(&userToDelete)
	if deleteErr != nil {
		return deleteErr
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "usuario eliminado",
		Data:       userToDelete,
	}
	WriteJSON(w, http.StatusOK, response)
	return nil
}

// HandleUpdateUser se encarga de actualizar una cuenta de usuario
// @Summary Actualiza una cuenta de usuario
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param user body models.UserSchema true "Usuario a actualizar"
// @Success 202 {object} structs.ApiResponse
// @Router /cuenta [put]
func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	userService := service.NewUserService()
	var user *models.User
	decodeErr := json.NewDecoder(r.Body).Decode(&user)
	if decodeErr != nil {
		return decodeErr
	}
	updatedUser, err := userService.UpdateUser(user)
	if err != nil {
		return err
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusAccepted,
		Message:    "usuario actualizado",
		Data:       updatedUser,
	}
	WriteJSON(w, http.StatusAccepted, response)
	return nil
}

// HandleUpdatePassword se encarga de actualizar la contraseña de un usuario
// @Summary Actualiza la contraseña de un usuario
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param id path int true "ID del usuario"
// @Param user body structs.ChangePasswordPayload true "Payload para cambiar la contraseña"
// @Success 202 {object} structs.ApiResponse
// @Router /cuenta/{id} [post]
func (s *APIServer) handleUpdatePassword(w http.ResponseWriter, r *http.Request) error {
	userService := service.NewUserService()
	id, parseErr := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if parseErr != nil {
		return parseErr
	}
	var payload *structs.ChangePasswordPayload
	if decodeErr := json.NewDecoder(r.Body).Decode(&payload); decodeErr != nil {
		return decodeErr
	}
	if changePassErr := userService.UpdatePassword(uint(id), payload.ActualPassword, payload.NewPassword); changePassErr != nil {
		return changePassErr
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusAccepted,
		Message:    "Clave actualizada",
	}
	WriteJSON(w, http.StatusAccepted, response)
	return nil
}
