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
	"github.com/golang-jwt/jwt/v5"
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

func (s *AuthService) Login(email string, password string) (string, *models.UserSchema, error) {
	var user *models.User
	if err := s.instance.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return "", nil, err
	}
	if !s.encrypter.VerifyPassword(password, user.Password) {
		fmt.Println(user)
		return "", nil, fmt.Errorf("error al verificar contraseña")
	}
	token, tokenErr := GenerateTokenV2(user)
	if tokenErr != nil {
		return "", user.ToSchema(), tokenErr
	}
	return token, user.ToSchema(), nil
}

func GenerateToken(user *models.User) string {
	var token string = user.Password + ":" + fmt.Sprintf("%d", rand.Int()) + ":" + fmt.Sprint(user.ID) + ":" + os.Getenv("SECRET")
	return token
}
func GenerateTokenV2(user *models.User) (string, error) {
	secret := []byte(os.Getenv("SECRET"))
	claims := jwt.MapClaims{
		"ID":     user.ID,
		"email":  user.Email,
		"nombre": user.Nombre,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
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
