package logger

import (
	"log/slog"
	"os"
)

var (
	Info  *slog.Logger
	Warn  *slog.Logger
	Error *slog.Logger
)

func InitLogger() {
	err := os.MkdirAll("logs", os.ModePerm)

	if err != nil {
		panic(err)
	}

	warnfile, err := os.OpenFile("logs/warn.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	infofile, err := os.OpenFile("logs/info.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	errorfile, err := os.OpenFile("logs/error.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	Warn = slog.New(slog.NewJSONHandler(warnfile, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))

	Info = slog.New(slog.NewJSONHandler(infofile, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	Error = slog.New(slog.NewJSONHandler(errorfile, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))
}
