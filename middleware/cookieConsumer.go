package middleware

import (
	"net/http"
	"strings"
)

func MiddlewareCookieConsumer(next http.Handler) {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			accessCookie, err := r.Cookie("access_token")
			if err != nil {
				http.Error(w, "no cookie access")
			}
			next.ServeHTTP(w, r)
		}
	)
}

func validateAccessToken(token string) bool {
	strings.Split(token, ":")
	return false
}
