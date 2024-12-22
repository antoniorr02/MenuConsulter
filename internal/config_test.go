package internal

import (
	"MenuConsulter/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	config.InitConfig()

	assert.NotNil(t, config.Config)
	assert.Equal(t, "MenuConsulter", config.Config.GetString("app.name"))
	assert.Equal(t, "development", config.Config.GetString("app.env"))
}
