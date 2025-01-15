package structs

import "github.com/Thomazoide/donde-caigo-backend/models"

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LikePayload struct {
	Post models.Post
	Id   uint
}
