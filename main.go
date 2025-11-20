package main

import (
	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/routes"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// 导入生成的docs包
	_ "github.com/JiquanZhong/realworld-go/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           用户管理系统API
// @version         1.0
// @description     这是一个基于Gin+GORM+SQLite的用户管理系统
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api/v1

// @schemes http https

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// 初始化日志系统
	utils.InitLogger()
	defer utils.Logger.Sync()

	utils.Logger.Info("Application starting...")

	db.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	utils.Logger.Info("Starting server on :8081")
	if err := r.Run(":8081"); err != nil {
		utils.Logger.Fatal("Failed to run server", zap.Error(err))
	}
}
