package config

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger(logFile string) {
	// Crear un logger con un formato JSON
	Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Abrir el archivo de log
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		zerolog.Fatal.Err(err).Msg("No se pudo abrir el archivo de log")
	}

	// Configurar salida múltiple: archivo y terminal
	multiWriter := io.MultiWriter(file, os.Stderr)
	Logger = Logger.Output(multiWriter)
}
