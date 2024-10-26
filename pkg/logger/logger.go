package logger

import (
	"errors"
	"log/slog"
	"os"
)

var (
	Log *slog.Logger
)

func New(
	mode string,
	filePath string,
) (*os.File, error) {
	logFile, errOpenFile := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if errOpenFile != nil {
		return nil, errOpenFile
	}

	switch mode {
	case "local":
		Log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		Log = slog.New(
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		Log = slog.New(
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		logFile.Close()
		return nil, errors.New("undefined log mode")
	}

	return logFile, nil
}
