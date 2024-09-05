package logger

import (
	"github.com/hanjm/zaplog"
	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.Logger
}

func New(level string) Logger {
	logger := zaplog.NewNoCallerLogger(level == "DEBUG")
	log := logger
	return Logger{Logger: log}
}

func (l Logger) Info(msg string, attrs ...zap.Field) {
	l.Logger.Info(msg, attrs...)
}

func (l Logger) Error(msg string, attrs ...zap.Field) {
	l.Logger.Error(msg, attrs...)
}

func (l Logger) Debug(msg string, attrs ...zap.Field) {
	l.Logger.Debug(msg, attrs...)
}

func (l Logger) Warn(msg string, attrs ...zap.Field) {
	l.Logger.Warn(msg, attrs...)
}
