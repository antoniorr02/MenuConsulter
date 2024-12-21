package internal

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLoggerInitialization(t *testing.T) {
	var buffer bytes.Buffer

	// Crear un logger de prueba
	logger := zerolog.New(&buffer).With().Timestamp().Logger()

	logger.Info().Msg("Esto es un mensaje de prueba")

	assert.NotEmpty(t, buffer.String(), "El logger no registró ningún mensaje")
	assert.Contains(t, buffer.String(), "Esto es un mensaje de prueba", "El mensaje no es el esperado")
}
