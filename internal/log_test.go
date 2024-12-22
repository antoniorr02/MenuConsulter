package internal

import (
	"bytes"
	"log/slog"
	"testing"
)

func TestLoggerInitialization(t *testing.T) {
	var buffer bytes.Buffer

	handler := slog.NewTextHandler(&buffer, &slog.HandlerOptions{
		AddSource: false,
	})

	logger := slog.New(handler)

	logger.Info("Esto es un mensaje de prueba")
	if buffer.Len() == 0 {
		t.Errorf("El logger no registró ningún mensaje")
	}
}
