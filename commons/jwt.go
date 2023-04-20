package commons

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

var (
	JwtKey = []byte("some_secret_key")
)
