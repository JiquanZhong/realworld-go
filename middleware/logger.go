package middleware

import (
	"time"

	"github.com/JiquanZhong/realworld-go/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ZapLogger Gin 日志中间件
func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 计算耗时
		duration := time.Since(startTime)

		// 记录日志
		utils.Logger.Debug(
			"HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.Duration("duration", duration),
			zap.String("user-agent", c.Request.UserAgent()),
		)
	}
}
