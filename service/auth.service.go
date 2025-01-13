package service

import (
	"fmt"

	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/middleware"
	"github.com/Thomazoide/donde-caigo-backend/models"
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

func (s *AuthService) Login(email string, password string) error {
	var user *models.User
	if err := s.instance.Where("email = ?").First(&user).Error; err != nil {
		return err
	}
	if !s.encrypter.VerifyPassword(password, user.Password) {
		return fmt.Errorf("error al verificar contraseña")
	}
	return nil
}
