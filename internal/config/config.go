package config

import (
	"os"

	"github.com/alecthomas/kong"
)

type Config struct {
	Env     string `help:"Entorno de la aplicación (development, production)." env:"APP_ENV" default:"development"`
	Logging struct {
		Level string `help:"Nivel de logging (debug, info, warn, error)." env:"APP_LOGGING_LEVEL" default:"info"`
	} `embed:"" prefix:"logging."`
}

var Cfg Config

func LoadConfig() {
	if os.Getenv("TEST_ENV") == "" {
		os.Setenv("TEST_ENV", "true")
	}
	if os.Getenv("TEST_ENV") == "true" {
		if Cfg.Env == "" {
			Cfg.Env = "development"
		}
		if Cfg.Logging.Level == "" {
			Cfg.Logging.Level = "info"
		}
		return
	}

	kong.Parse(&Cfg, kong.Name("MenuConsulter"), kong.Description("Una aplicación para gestionar menús."))
}
