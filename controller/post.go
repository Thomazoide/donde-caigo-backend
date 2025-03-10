package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Thomazoide/donde-caigo-backend/models"
	"github.com/Thomazoide/donde-caigo-backend/service"
	"github.com/Thomazoide/donde-caigo-backend/structs"
)

func (s *APIServer) handlePost(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetPost(w, r)
	}
	if r.Method == http.MethodPost {
		return s.handleCreatePost(w, r)
	}
	if r.Method == http.MethodPatch {
		return s.handleAddLike(w, r)
	}
	if r.Method == http.MethodDelete {
		return s.handleDeletePost(w, r)
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "METHOD NOT ALLOWED",
		Error:      fmt.Errorf("method '%s' not allowed\nstatus: %d", r.Method, http.StatusMethodNotAllowed).Error(),
	}
	WriteJSON(w, http.StatusMethodNotAllowed, response)
	return nil
}

// HandleGetPost se encarga de entregar todos los posts
// @Summary Entrega todos los posts
// @Tags Publicaciones
// @Produce json
// @Success 200 {object} structs.ApiResponse
// @Router /publicaciones [get]
func (s *APIServer) handleGetPost(w http.ResponseWriter, r *http.Request) error {
	postService := service.NewPostService()
	posts, err := postService.GetAllPost()
	if err != nil {
		return err
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "posts",
		Data:       posts,
	}
	fmt.Println(r.Method)
	WriteJSON(w, http.StatusOK, response)
	return nil
}

// HandleCreatePost se encarga de crear un post
// @Summary Crear un post
// @Tags Publicaciones
// @Accept json
// @Produce json
// @Param body body structs.NewPostPayload true "Nuevo post"
// @Success 201 {object} structs.ApiResponse
// @Router /publicaciones [post]
func (s *APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	postService := service.NewPostService()
	var newPostPayload *structs.NewPostPayload
	decodeErr := json.NewDecoder(r.Body).Decode(&newPostPayload)
	if decodeErr != nil {
		return decodeErr
	}
	post, err := postService.CreatePost(*newPostPayload, r.Context())
	if err != nil {
		return err
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusCreated,
		Message:    "post creado",
		Data:       post,
	}
	WriteJSON(w, http.StatusCreated, response)
	return nil
}

// HandleAddLike se encarga de agregar un like a un post
// @Summary Agregar un like a un post
// @Tags Publicaciones
// @Accept json
// @Produce json
// @Param body body structs.LikePayload true "Like a agregar"
// @Success 202 {object} structs.ApiResponse
// @Router /publicaciones [patch]
func (s *APIServer) handleAddLike(w http.ResponseWriter, r *http.Request) error {
	postService := service.NewPostService()
	var likePayload *structs.LikePayload
	decodeErr := json.NewDecoder(r.Body).Decode(&likePayload)
	if decodeErr != nil {
		return decodeErr
	}
	updatedPost, updateErr := postService.AddLike(&likePayload.Post, likePayload.Id)
	if updateErr != nil {
		return updateErr
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusAccepted,
		Message:    "Like",
		Data:       updatedPost,
	}
	WriteJSON(w, http.StatusAccepted, response)
	return nil
}

// HandleDeletePost se encarga de eliminar un post
// @Summary Eliminar un post
// @Tags Publicaciones
// @Accept json
// @Produce json
// @Param body body models.PostSchema true "Post a eliminar"
// @Success 200 {object} structs.ApiResponse
// @Router /publicaciones [delete]
func (s *APIServer) handleDeletePost(w http.ResponseWriter, r *http.Request) error {
	postService := service.NewPostService()
	var postToDelete *models.PostSchema
	decodeErr := json.NewDecoder(r.Body).Decode(&postToDelete)
	if decodeErr != nil {
		return decodeErr
	}
	err := postService.DeletePost(postToDelete.ID)
	if err != nil {
		return err
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "post eliminado",
		Data:       postToDelete,
	}
	WriteJSON(w, http.StatusOK, response)
	return nil
}

func (s *APIServer) DeleteOldPosts(w http.ResponseWriter, r *http.Request) error {
	postService := service.NewPostService()
	err := postService.DeleteOldPosts()
	if err != nil {
		return err
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "posts eliminados",
	}
	WriteJSON(w, http.StatusOK, response)
	return nil
}
