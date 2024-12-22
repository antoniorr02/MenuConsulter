package config

import (
	"io"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger(logFile string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		slog.Warn("Could not open log file, falling back to stdout", "error", err)
		file = os.Stdout
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	// Configurar el handler de slog con opciones
	handler := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		AddSource: true,
	})

	Logger = slog.New(handler)
}
