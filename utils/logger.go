package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

// InitLogger 初始化日志系统
func InitLogger() {
	// 定义日志输出格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建日志文件轮转器
	fileWriter := &lumberjack.Logger{
		Filename:   "logs/app.log", // 日志文件路径
		MaxSize:    100,            // 单个日志文件最大大小 MB
		MaxBackups: 10,             // 保留最多备份文件数
		MaxAge:     30,             // 保留最多天数
		Compress:   true,           // 是否压缩旧日志文件
	}

	// 定义多个输出目标：文件 + 控制台
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 文件输出
	fileCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(fileWriter),
		zapcore.InfoLevel,
	)

	// 控制台输出（开发环境）
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	// 合并多个输出
	core := zapcore.NewTee(fileCore, consoleCore)

	// 创建 logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer Logger.Sync()
}

// GetLogger 获取全局 logger 实例
func GetLogger() *zap.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}
