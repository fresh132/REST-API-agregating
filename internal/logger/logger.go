package logger

import (
	"log/slog"
	"os"
	"strings"
)

var (
	Debug *slog.Logger
	Info  *slog.Logger
	Warn  *slog.Logger
	Error *slog.Logger
)

func InitLogger() {
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Get log level from environment or default to info
	logLevelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevelStr == "" {
		logLevelStr = "info"
	}

	var level slog.Level
	switch logLevelStr {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Debug logger
	debugfile, err := os.OpenFile("logs/debug.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Debug = slog.New(slog.NewJSONHandler(debugfile, &slog.HandlerOptions{
		Level: level,
	}))

	// Info logger
	infofile, err := os.OpenFile("logs/info.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Info = slog.New(slog.NewJSONHandler(infofile, &slog.HandlerOptions{
		Level: level,
	}))

	// Warn logger
	warnfile, err := os.OpenFile("logs/warn.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Warn = slog.New(slog.NewJSONHandler(warnfile, &slog.HandlerOptions{
		Level: level,
	}))

	// Error logger
	errorfile, err := os.OpenFile("logs/error.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Error = slog.New(slog.NewJSONHandler(errorfile, &slog.HandlerOptions{
		Level: level,
	}))
}
