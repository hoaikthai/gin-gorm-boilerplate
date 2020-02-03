package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init configuration based on env
func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

// GetConfig returns config object
func GetConfig() *viper.Viper {
	return config
}
