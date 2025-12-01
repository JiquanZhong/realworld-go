package routes

import (
	"github.com/JiquanZhong/realworld-go/handlers"
	"github.com/JiquanZhong/realworld-go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(v1 *gin.RouterGroup) {
	users := v1.Group("/users")
	users.POST("/login", handlers.Login)
	users.POST("", handlers.CreateUser)

	protected := v1.Group("/users")
	protected.Use(middleware.JWTAuth())
	users.GET("/:id", handlers.GetUser)
	users.PUT("/:id", handlers.UpdateUser)

	admin := v1.Group("users")
	admin.Use(middleware.JWTAuth(), middleware.RequireAdmin())
	admin.GET("", handlers.GetUsers)
	admin.DELETE("/:id", handlers.DeleteUser)
	admin.PUT("/:id/role", handlers.UpdateUserRole)
}
