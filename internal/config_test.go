package internal

import (
	"MenuConsulter/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	config.LoadConfig("config/config.yaml")

	assert.NotNil(t, config.K)
	assert.Equal(t, "MenuConsulter", config.K.String("app.name"))
	assert.Equal(t, "development", config.K.String("app.env"))
}
