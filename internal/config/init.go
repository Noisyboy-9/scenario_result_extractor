package config

import (
	"fmt"

	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/spf13/viper"
)

func LoadViper() {
	configPath := "configs/general.yaml"
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file in path: %s can't be found", configPath))
	}
}

func Init() {
	var err error

	Prometheus = new(prometheus)
	err = viper.UnmarshalKey("prometheus", Prometheus)
	if err != nil {
		log.App.Fatal(err)
	}
}
