package main

import "github.com/golang-jwt/jwt"

var jwtKey = []byte("secret")

type (
	AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	Claims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}
)

type ErrorHandle struct {
	ErrorStatus  int         `json:"error_status"`
	ErrorMessage string      `json:"error_message"`
	ErrorInput   interface{} `json:"error_input"`
	ErrorFrom    string      `json:"error_from"`
}
