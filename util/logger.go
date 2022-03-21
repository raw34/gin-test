package util

import (
    "go.uber.org/zap"
)

type LoggerInterface interface {
    Debug(args ...interface{})
    Info(args ...interface{})
    Warn(args ...interface{})
    Error(args ...interface{})
    Fatal(args ...interface{})
}

func NewLogger() *Logger {
    logger, _ := zap.NewProduction()
    return &Logger{logger: logger, sugar: logger.Sugar()}
}

type Logger struct {
    logger *zap.Logger
    sugar  *zap.SugaredLogger
}

func (l *Logger) Debug(args ...interface{}) {
    l.sugar.Debug("DEBUG", args)
}

func (l *Logger) Info(args ...interface{}) {
    l.sugar.Info("INFO", args)
}

func (l *Logger) Warn(args ...interface{}) {
    l.sugar.Warn("WARN", args)
}

func (l *Logger) Error(args ...interface{}) {
    l.sugar.Error("Error", args)
}

func (l *Logger) Fatal(args ...interface{}) {
    l.sugar.Fatal("FATAL", args)
}
