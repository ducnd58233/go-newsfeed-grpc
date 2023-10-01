package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initializes the global logger instance.
func NewLogger() (*zap.Logger, error) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, err
	}

	logFileName := filepath.Join(logDir, fmt.Sprintf("app-%s.log", time.Now().Format("2006-01-02")))
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Add color to level
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(logFile)),
			zap.DebugLevel,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig), // Colored console output
			zapcore.Lock(os.Stdout),
			zap.DebugLevel,
		),
	)

	logger := zap.New(core)

	return logger, nil
}

// SyncLogger syncs the logger.
// SyncLogger syncs the logger.
func SyncLogger(logger *zap.Logger) error {
	if logger == nil {
		return nil // Logger not initialized
	}
	if err := logger.Sync(); err != nil {
		return err
	}
	return nil
}
