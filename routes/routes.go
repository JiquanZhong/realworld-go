package routes

import (
	"github.com/JiquanZhong/realworld-go/handlers"
	"github.com/JiquanZhong/realworld-go/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 添加日志中间件
	r.Use(middleware.ZapLogger())

	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	v1 := r.Group("/api/v1")
	{

		users := v1.Group("/users")
		{
			users.POST("/login", handlers.Login)
			users.POST("", handlers.CreateUser)
		}

		mcps := v1.Group("/mcps")
		{
			mcps.GET("", handlers.GetMcpServices)
			mcps.GET("/:id", handlers.GetMcpService)
		}

		mcpTags := v1.Group("/mcp-tags")
		{
			mcpTags.GET("", handlers.GetMcpTags)
		}

		// 需要认证的路由
		protected := v1.Group("/")
		protected.Use(middleware.JWTAuth())
		{
			users := protected.Group("/users")

			users.GET("/:id", handlers.GetUser)
			users.PUT("/:id", handlers.UpdateUser)

			mcps := protected.Group("/mcps")
			mcps.POST("", handlers.RegisterMcpService)

		}

		// 需要管理员权限的路由
		admin := v1.Group("/")
		admin.Use(middleware.JWTAuth(), middleware.RequireAdmin())
		{
			adminUsers := admin.Group("/users")

			adminUsers.GET("", handlers.GetUsers)
			adminUsers.DELETE("/:id", handlers.DeleteUser)
			adminUsers.PUT("/:id/role", handlers.UpdateUserRole)

			mcps := admin.Group("/mcps")
			mcps.DELETE("/:id", handlers.DeleteMcpService)
			


			mcpTags := admin.Group("/mcp-tags")
			mcpTags.POST("",handlers.CreateMcpTag)
			mcpTags.DELETE("/:id", handlers.DeleteMcpTag)
			mcpTags.PUT("/:id", handlers.UpdateMcpTag)
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
