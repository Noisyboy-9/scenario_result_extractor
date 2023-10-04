package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadViper() {
	configPath := "configs/config.yaml"
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file in path: %s can't be found", configPath))
	}
}

func Init() {
	HttpServer = new(httpServer)
	err := viper.UnmarshalKey("httpServer", HttpServer)
	if err != nil {
		// fatal log
	}
}
