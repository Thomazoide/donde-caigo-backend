package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/models"
)

func MiddleWareCookieConsumer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		method := r.Method
		if strings.HasPrefix(uri, "/docs/") {
			next.ServeHTTP(w, r)
			return
		}
		if (uri == "/auth" || uri == "/cuenta") && (method == http.MethodPost) {
			next.ServeHTTP(w, r)
			return
		}
		accessCookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "access token not fount", http.StatusUnauthorized)
			return
		}
		if !validateAccessToken(accessCookie.Value) {
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

func validateAccessToken(token string) bool {
	args := strings.Split(token, ":")
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
