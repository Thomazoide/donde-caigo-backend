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
		Error:      fmt.Errorf("method '%s' not allowed\nstatus: %d", r.Method, http.StatusMethodNotAllowed),
	}
	WriteJSON(w, http.StatusMethodNotAllowed, response)
	return nil
}

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

func (s *APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	postService := service.NewPostService()
	var tmpPost *models.Post
	decodeErr := json.NewDecoder(r.Body).Decode(&tmpPost)
	if decodeErr != nil {
		return decodeErr
	}
	newPost, err := postService.CreatePost(tmpPost.Title, tmpPost.Description, tmpPost.Pics, tmpPost.AuthorID)
	if err != nil {
		return err
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusCreated,
		Message:    "post creado",
		Data:       newPost,
	}
	WriteJSON(w, http.StatusCreated, response)
	return nil
}

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

func (s *APIServer) handleDeletePost(w http.ResponseWriter, r *http.Request) error {
	postService := service.NewPostService()
	var postToDelete *models.Post
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
