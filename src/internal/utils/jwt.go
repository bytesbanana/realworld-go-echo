package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(username, email string) (string, error) {
	claims := jwtCustomClaims{
		username,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * 24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
