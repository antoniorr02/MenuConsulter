package internal

import (
	"MenuConsulter/internal/config"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	// Ensure default values are set
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development")
	}

	// Load the config
	config.LoadConfig()

	// Check if the Env is set correctly
	if config.Cfg.Env != "development" {
		t.Errorf("Expected Env to be 'development', got '%s'", config.Cfg.Env)
	}
}
