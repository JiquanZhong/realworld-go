package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireAdmin 中间件用于验证用户是否为管理员
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户角色（由 JWTAuth 中间件设置）
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role information not found"})
			c.Abort()
			return
		}

		// 验证角色是否为管理员
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}
