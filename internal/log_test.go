package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLoggerInitialization(t *testing.T) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	logger, _ := config.Build()
	defer logger.Sync()

	logger.Info("Esto es un mensaje de prueba")

	assert.NotNil(t, logger, "El logger no está inicializado")
}
