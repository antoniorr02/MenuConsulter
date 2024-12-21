package internal

import (
	"bytes"
	"log/slog"
	"testing"
)

func TestLoggerInitialization(t *testing.T) {
	var buffer bytes.Buffer

	// Crear un handler con opciones para el logger
	handler := slog.NewTextHandler(&buffer, &slog.HandlerOptions{
		AddSource: false, // En el test, no necesitamos la fuente
	})

	logger := slog.New(handler)

	logger.Info("Esto es un mensaje de prueba")
	if buffer.Len() == 0 {
		t.Errorf("El logger no registró ningún mensaje")
	}
}
