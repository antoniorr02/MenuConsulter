package config

import (
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func InitLogger(logFile string) {
	var err error

	// Crear un logger con zap
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{logFile, "stderr"}

	Logger, err = config.Build()
	if err != nil {
		panic("No se pudo inicializar el logger: " + err.Error())
	}

	// Sync es necesario si necesitas manejar la seguridad de los logs
	defer Logger.Sync()
}
