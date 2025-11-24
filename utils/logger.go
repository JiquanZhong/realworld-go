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

	// 创建不同级别的日志文件轮转器
	debugWriter := &lumberjack.Logger{
		Filename:   "logs/debug.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}

	infoWriter := &lumberjack.Logger{
		Filename:   "logs/info.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}

	warnWriter := &lumberjack.Logger{
		Filename:   "logs/warn.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}

	errorWriter := &lumberjack.Logger{
		Filename:   "logs/error.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// DEBUG 级别日志文件
	debugCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(debugWriter),
		zapcore.DebugLevel,
	)

	// INFO 级别日志文件（只记录 INFO 级别）
	infoCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(infoWriter),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)

	// WARN 级别日志文件（只记录 WARN 级别）
	warnCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(warnWriter),
		zap.NewAtomicLevelAt(zapcore.WarnLevel),
	)

	// ERROR 级别日志文件（记录 ERROR 和 FATAL）
	errorCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(errorWriter),
		zap.NewAtomicLevelAt(zapcore.ErrorLevel),
	)

	// 控制台输出（所有级别）
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	// 合并所有输出
	core := zapcore.NewTee(debugCore, infoCore, warnCore, errorCore, consoleCore)

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
