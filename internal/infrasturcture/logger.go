package infrasturcture

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	logLevelStr := os.Getenv("LOG_LEVEL")
	var logLevel zap.AtomicLevel

	switch logLevelStr {
	case "debug":
		logLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		logLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	default:
		logLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	config := zap.Config{
		Encoding:         "json",
		Level:            logLevel,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	logger, _ = config.Build()

	LogInfo("Logger Zap inited!")
}

func LogErr(err error, fields ...zap.Field) {
	logger.Error(fmt.Sprintf("%v", err), fields...)
}

func LogWarn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func LogInfo(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func SyncLogger() {
	logger.Sync()
}
