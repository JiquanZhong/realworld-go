package routes

import (
	"github.com/JiquanZhong/realworld-go/handlers"
	"github.com/JiquanZhong/realworld-go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMcpTagRoutes(v1 *gin.RouterGroup) {
	mcpTags := v1.Group("/mcp-tags")
	mcpTags.GET("", handlers.GetMcpTags)

	admin := v1.Group("mcp-tags")
	admin.Use(middleware.JWTAuth(), middleware.RequireAdmin())
	admin.POST("", handlers.CreateMcpTag)
	admin.DELETE("/:id", handlers.DeleteMcpTag)
	admin.PUT("/:id", handlers.UpdateMcpTag)
}
