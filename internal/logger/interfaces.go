package logger

import "log/slog"

// Logger defines the interface for logging operations
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// LoggerImpl implements Logger interface
type LoggerImpl struct {
	logger *slog.Logger
}

func NewLogger(logger *slog.Logger) Logger {
	return &LoggerImpl{logger: logger}
}

func (l *LoggerImpl) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}

func (l *LoggerImpl) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

func (l *LoggerImpl) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

func (l *LoggerImpl) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}
