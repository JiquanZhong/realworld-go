package middleware

import (
	"net/http"
	"strings"

	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var jwtSecret = []byte("your_secret_key")

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement JWT authentication logic here
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.GetLogger().Error("Authorization header is missing")
			utils.Error(c, http.StatusUnauthorized, "Authorization header is missing")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.GetLogger().Error("Invalid Authorization header format")
			utils.Error(c, http.StatusUnauthorized, "Invalid Authorization header format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			utils.GetLogger().Error("Invalid token", zap.Error(err))
			utils.Error(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if err != nil || !token.Valid {
			utils.GetLogger().Error("Invalid token", zap.Error(err))
			utils.Error(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("userID", claims["user_id"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}
