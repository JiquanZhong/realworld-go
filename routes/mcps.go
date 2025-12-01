package routes

import (
	"github.com/JiquanZhong/realworld-go/handlers"
	"github.com/JiquanZhong/realworld-go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMcpRoutes(v1 *gin.RouterGroup) {
	mcps := v1.Group("/mcps")
	mcps.GET("", handlers.GetMcpServices)
	mcps.GET("/:id", middleware.OptionalJWTAuth(), handlers.GetMcpService)

	protected := v1.Group("/mcps")
	protected.Use(middleware.JWTAuth())
	protected.POST("", handlers.RegisterMcpService)
	protected.POST("/:id/favorite", handlers.AddMcpServiceFavorite)
	protected.DELETE("/:id/favorite", handlers.RemoveMcpServiceFavorite)

	admin := v1.Group("mcps")
	admin.Use(middleware.JWTAuth(), middleware.RequireAdmin())
	admin.DELETE("/:id", handlers.DeleteMcpService)
}
