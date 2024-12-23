package internal

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerInitialization(t *testing.T) {
	logFile, err := os.CreateTemp("", "logfile-temporal.log")
	assert.NoError(t, err, "No se pudo crear el archivo de log temporal")
	defer os.Remove(logFile.Name())

	handler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
		AddSource: false,
	})

	logger := slog.New(handler)

	logger.Info("Esto es un mensaje de prueba")

	logFile.Close()

	content, err := os.ReadFile(logFile.Name())
	assert.NoError(t, err, "No se pudo leer el archivo de log")

	assert.Contains(t, string(content), "Esto es un mensaje de prueba", "El log debería contener el mensaje 'Esto es un mensaje de prueba'")
}
