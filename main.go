package main

import (
	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/routes"
	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志系统
	utils.InitLogger()
	defer utils.Logger.Sync()

	utils.Logger.Info("Application starting...")

	db.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	utils.Logger.Info("Starting server on :8081")
	if err := r.Run(":8081"); err != nil {
		utils.Logger.Fatal("Failed to run server", zap.Error(err))
	}
}
