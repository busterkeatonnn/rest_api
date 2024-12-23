package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
	Role string `json:"role"`
}

var JwtKey = []byte("very_secret_phrase")
