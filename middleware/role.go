package middleware

import (
	"net/http"

	"github.com/JiquanZhong/realworld-go/cst"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
)

// RequireAdmin 中间件用于验证用户是否为管理员
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户角色（由 JWTAuth 中间件设置）
		role, exists := c.Get("role")
		if !exists {
			utils.Error(c, http.StatusUnauthorized, "请先登录")
			return
		}

		// 验证角色是否为管理员
		if role != cst.RoleAdmin {
			utils.Error(c, http.StatusForbidden, "无权限")
			return
		}

		c.Next()
	}
}
