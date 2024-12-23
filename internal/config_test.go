package internal

import (
	"bytes"
	"testing"

	"MenuConsulter/internal/config"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	var logOutput bytes.Buffer

	config.InitConfig()

	logger := logrus.New()
	logger.Out = &logOutput
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{})

	logger.Info("Prueba de log")

	logContent := logOutput.String()
	assert.Contains(t, logContent, "Prueba de log", "El contenido del log debería contener el mensaje esperado")
}
