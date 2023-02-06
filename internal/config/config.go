package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init() {
	viper.SetConfigFile(".env")

	viper.ReadInConfig()
}

// GetCriticalConfigParam @deprecated
func GetCriticalConfigParam[T any](getter func() interface{}) T {
	value, ok := getter().(T)

	if !ok {
		log.Fatal("A critical configuration parameter is not defined during configuration read!")
	}

	return value
}

func Get[T any](parameterName string) T {
	value, ok := viper.Get(parameterName).(T)

	if !ok {
		log.Fatalf("A critical configuration parameter under '%v' key is not found", parameterName)
	}

	return value
}
