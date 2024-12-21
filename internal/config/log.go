package config

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger(logFile string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo de log: %v", err)
	}

	// Configurar salida múltiple: archivo y terminal
	multiWriter := io.MultiWriter(file, os.Stdout)

	Logger = log.New(multiWriter, "MenuConsulter: ", log.Ldate|log.Ltime|log.Lshortfile)
}
