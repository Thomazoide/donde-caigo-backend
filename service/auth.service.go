package service

import (
	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/middleware"
	"gorm.io/gorm"
)

type AuthService struct {
	instance  *gorm.DB
	encrypter *middleware.Encrypter
}

func NewAuthService() *AuthService {
	return &AuthService{
		instance:  config.GetInstance(),
		encrypter: middleware.NewEncrypter(),
	}
}

func (s *AuthService) Login(email string, password string) {
}
