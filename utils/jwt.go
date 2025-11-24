package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func GetUserID(c *gin.Context) (uint, bool) {
	v, ok := c.Get("userID") // middleware/JWTAuth 中用的是 userID
	if !ok {
		return 0, false
	}

	switch id := v.(type) {
	case float64: // JWT MapClaims 解码后常是 float64
		return uint(id), true
	case uint:
		return id, true
	case int:
		return uint(id), true
	default:
		return 0, false
	}
}

func GetUserRole(c *gin.Context) (string, bool) {
	v, ok := c.Get("role")
	if !ok {
		return "", false
	}
	role, ok := v.(string)
	return role, ok
}
