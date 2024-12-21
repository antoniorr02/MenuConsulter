package config

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger(logFile string) {
	// Crear un nuevo logger
	Logger = logrus.New()

	// Configurar formato
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Abrir el archivo de log
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("No se pudo abrir el archivo de log: %v", err)
	}

	// Configurar salida múltiple: archivo y terminal
	multiWriter := io.MultiWriter(file, os.Stdout)
	Logger.SetOutput(multiWriter)

	// Establecer nivel de log
	Logger.SetLevel(logrus.InfoLevel)
}
