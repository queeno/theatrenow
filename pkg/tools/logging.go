package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func init() {
	logger, _ = zapConfig().Build()
}

func Logger() *zap.Logger {
	return logger
}

var logLevels = map[string]zapcore.Level {
	"info": zapcore.InfoLevel,
	"warn": zapcore.WarnLevel,
	"debug": zapcore.DebugLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func zapConfig() *zap.Config {
	logLevelEnv := os.Getenv("LOG_LEVEL")

	if _, ok := logLevels[logLevelEnv]; !ok {
		logLevelEnv = "info"
	}

	hostname, _ := os.Hostname()

	return &zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(logLevels[logLevelEnv]),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "@timestamp",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		InitialFields: map[string]interface{}{
			"app_owner":   "simon",
			"app_name":    "theatrenow",
			"logger_name": "theatrenow",
			"host_name":   hostname,
		},
	}
}