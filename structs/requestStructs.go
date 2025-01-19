package structs

import "github.com/Thomazoide/donde-caigo-backend/models"

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LikePayload struct {
	Post models.Post `json:"post"`
	Id   uint        `json:"id"`
}

type ChangePasswordPayload struct {
	ActualPassword string `json:"actualPassword"`
	NewPassword    string `json:"newPassword"`
}
