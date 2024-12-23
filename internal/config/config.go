package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetDefault("app.name", "MenuConsulter")
	v.SetDefault("app.env", "development")
	v.SetDefault("log.level", "info")

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		log.Printf("No se encontró el archivo de configuración: %v", err)
	} else {
		log.Printf("Archivo de configuración cargado: %s", v.ConfigFileUsed())
	}

	Config = v
}
