package structs

import "github.com/Thomazoide/donde-caigo-backend/models"

type ApiResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Error      string `json:"error,omitempty"`
}

type LoginResponse struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Token      string            `json:"token"`
	UserData   models.UserSchema `json:"userData"`
}
