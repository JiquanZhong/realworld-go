package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	UserID uint   `json:"user_id"`
	name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(userId uint, name, email string) (string, error) {
	claims := Claims{
		UserID: userId,
		name:   name,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(24 * time.Hour).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err

}
