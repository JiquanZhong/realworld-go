package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userId uint, name, email, role string) (string, error) {
	claims := Claims{
		UserID: userId,
		Name:   name,
		Email:  email,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(24 * time.Hour).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		GetLogger().Error("Failed to generate token", zap.Error(err))
		return "", err
	}
	return tokenString, err

}
