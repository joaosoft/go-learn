package config

import (
	"fmt"
	"github.com/labstack/gommon/log"
	viper "github.com/spf13/viper"
)

type simpleConfig struct {
	viperConfig *viper.Viper
}

// Get ... get a configuration by key
func (instance *simpleConfig) Get(key string) interface{} {
	return instance.viperConfig.Get(key)
}

func NewSimpleConfig(path string, file string, extension string) (*simpleConfig, error) {
	config, err := LoadConfig(path, file, extension)
	return &simpleConfig{
		viperConfig: config,
	}, err
}

// LoadConfig ... loads the configuration file
func LoadConfig(path string, file string, extension string) (*viper.Viper, error) {
	viperConfig := viper.New()

	viperConfig.SetConfigName(file)
	viperConfig.SetConfigType(extension)
	viperConfig.AddConfigPath(path)

	if err := viperConfig.ReadInConfig(); err != nil {
		log.Error(err)
		return nil, err
	}
	fmt.Println(viperConfig)

	return viperConfig, nil
}
