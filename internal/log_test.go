package internal

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLoggerInitialization(t *testing.T) {
	var buffer bytes.Buffer

	// Crear un logger de prueba
	logger := logrus.New()
	logger.SetOutput(&buffer)
	logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	logger.Info("Esto es un mensaje de prueba")

	assert.NotEmpty(t, buffer.String(), "El logger no registró ningún mensaje")
	assert.Contains(t, buffer.String(), "Esto es un mensaje de prueba", "El mensaje no es el esperado")
}
