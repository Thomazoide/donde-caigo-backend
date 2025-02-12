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
	newPostString := r.FormValue("post")
	var newPost *models.Post
	decodeErr := json.Unmarshal([]byte(newPostString), &newPost)
	createdPost, createPostErr := postService.CreatePost(newPost.Title, newPost.Description, "", newPost.AuthorID)
	if createPostErr != nil {
		return fmt.Errorf("error al crear post\nDetalles: %v", createPostErr)
	}
	if decodeErr != nil {
		return fmt.Errorf("error al decodificar json\nDetalles: %v", decodeErr)
	}
	sizeErr := r.ParseMultipartForm(10 << 20)
	if sizeErr != nil {
		return fmt.Errorf("error al parsear formulario\nDetalles: %v", sizeErr)
	}
	var imageURLS string = ""
	for _, fileHeaders := range r.MultipartForm.File["pics"] {
		file, fileErr := fileHeaders.Open()
		if fileErr != nil {
			return fmt.Errorf("error al abrir archivo\nDetalles: %v", fileErr)
		}
		defer file.Close()
		imageURL, uploadErr := service.UploadImages(r.Context(), fileHeaders.Filename, createdPost.ID, file)
		if uploadErr != nil {
			return fmt.Errorf("error al subir imagen\nDetalles: %v", uploadErr)
		}
		if imageURLS == "" {
			imageURLS = imageURL
		} else {
			imageURLS = imageURLS + "," + imageURL
		}
	}
	createdPost.Pics = imageURLS
	if _, updateErr := postService.EditPost(createdPost); updateErr != nil {
		return fmt.Errorf("error al editar post\nDetalles: %v", updateErr)
	}
	response := &structs.ApiResponse{
		StatusCode: http.StatusCreated,
		Message:    "post creado",
		Data:       createdPost,
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
