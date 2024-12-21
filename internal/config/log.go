package config

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(logFile string) {
	var err error

	// Crear un logger con zap
	Logger, err = zap.NewProduction()
	if err != nil {
		panic("No se pudo inicializar el logger: " + err.Error())
	}

	// Abrir el archivo de log
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Logger.Fatal("No se pudo abrir el archivo de log", zap.Error(err))
	}

	// Configurar salida múltiple: archivo y terminal
	multiWriter := zap.CombineWriteSyncers(zapcore.AddSync(file), zapcore.AddSync(os.Stderr))
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), multiWriter, zap.InfoLevel)
	Logger = zap.New(core)
}
