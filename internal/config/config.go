package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var K = koanf.New(".")

func LoadConfig(configPath string) {
	// Cargar configuración desde un archivo YAML
	if err := K.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		log.Fatalf("Error al cargar el archivo de configuración: %v", err)
	}

	// Sobrescribir con variables de entorno (prefijo APP_)
	if err := K.Load(env.Provider("APP_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, "APP_")), "_", ".")
	}), nil); err != nil {
		log.Fatalf("Error al cargar las variables de entorno: %v", err)
	}

	log.Println("Configuración cargada exitosamente")
}
