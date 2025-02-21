package internal

import (
	"testing"

	"MenuConsulter/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestConfiguracion(t *testing.T) {
	config.InitConfig()

	assert.Equal(t, "production", config.Config.GetString("app.env"), "El entorno por defecto debería ser 'production'")
}
