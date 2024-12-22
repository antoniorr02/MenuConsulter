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
		slog.Warn("No se puede abrir el fichero de log", "error", err)
		file = os.Stdout
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	handler := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		AddSource: true,
	})

	Logger = slog.New(handler)
}
