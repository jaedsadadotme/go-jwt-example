package main

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			responseErrorWithJSON(w, http.StatusUnauthorized, "Malformed Token")
			return
		} else {
			jwtToken := authHeader[1]
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			if err != nil {
				responseErrorWithJSON(w, http.StatusUnauthorized, err.Error())
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
