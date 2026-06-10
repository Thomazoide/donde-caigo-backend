package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/models"
	"github.com/golang-jwt/jwt/v5"
)

func MiddleWareCookieConsumer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		method := r.Method
		if strings.HasPrefix(uri, "/docs/") {
			next.ServeHTTP(w, r)
			return
		}
		if method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}
		if (uri == "/auth" || uri == "/cuenta") && (method == http.MethodPost) {
			next.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header not found", http.StatusUnauthorized)
			return
		}
		token := authHeader
		lower := strings.ToLower(authHeader)
		if strings.HasPrefix(lower, "bearer ") {
			token = strings.TrimSpace(authHeader[7:])
		}
		if !validateAccessTokenV2(token) {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func parseID(str string) uint {
	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	i := uint(id)
	return i
}

// deprecated
func validateAccessToken(token string) bool {
	args := strings.Split(token, ":")
	if len(args) < 3 {
		return false
	}
	db := config.GetInstance()
	id := parseID(args[2])
	pass := args[0]
	var user *models.User
	err := db.Where("id = ?", id).Select("Password").First(&user).Error
	if err != nil {
		return false
	}
	return user.Password == pass
}

func validateAccessTokenV2(tokenString string) bool {
	secret := []byte(os.Getenv("SECRET"))
	token, err := jwt.ParseWithClaims(
		tokenString,
		jwt.MapClaims{},
		func(tk *jwt.Token) (any, error) {
			if tk.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Metodo de encriptacion incorrecto...")
			}
			return secret, nil
		},
	)
	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	return true
}
