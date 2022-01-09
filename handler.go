package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	datas := map[string]interface{}{
		"hello": "world",
	}
	responseWithJSON(w, http.StatusOK, datas)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var request AuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responseErrorWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// check user exists
	if request.Username != "admin" || request.Password != "pingpong" {
		responseErrorWithJSON(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: request.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]interface{}{
		"access_token": tokenString,
	})
}
