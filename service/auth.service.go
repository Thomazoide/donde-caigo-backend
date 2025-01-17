package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/middleware"
	"github.com/Thomazoide/donde-caigo-backend/models"
	"github.com/joho/godotenv"
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

func (s *AuthService) Login(email string, password string) (*string, error) {
	var user *models.User
	if err := s.instance.Where(&models.User{Email: email}).Select("email", "password", "ID").First(&user).Error; err != nil {

		return nil, err
	}
	if !s.encrypter.VerifyPassword(password, user.Password) {
		fmt.Println(user)
		return nil, fmt.Errorf("error al verificar contraseña")
	}
	var token string = GenerateToken(user)
	return &token, nil
}

func GenerateToken(user *models.User) string {
	godotenv.Load()
	var token string = user.Password + ":" + fmt.Sprintf("%d", rand.Int()) + ":" + fmt.Sprint(user.ID) + ":" + os.Getenv("SECRET")
	return token
}

func (s *AuthService) SignCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})
}
