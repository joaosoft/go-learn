package pm

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/config"
)

// -------------- CONFIGURATION CLIENTS --------------
// NewJSONFile ... creates a new nsq producer
func (instance *manager) NewSimpleConfig(path string, file string, extension string) (config.IConfig, error) {
	return config.NewSimpleConfig(path, file, extension)
}

// -------------- METHODS --------------
// GetConfig ... get a config with key
func (instance *manager) GetConfig(key string) config.IConfig {
	return instance.ConfigController[key]
}

// AddProcess ... add a config with key
func (instance *manager) AddConfig(key string, cfg config.IConfig) error {
	if instance.Started {
		panic("manager, can not add config after start")
	}

	instance.ConfigController[key] = &config.ConfigController{
		Path:   "",
		Config: cfg}

	log.Infof(fmt.Sprintf("manager, config '%s' added", key))

	return nil
}

// RemConfig ... remove the config by bey
func (instance *manager) RemConfig(key string) (config.IConfig, error) {
	// get config
	controller := instance.ConfigController[key]

	// delete config
	delete(instance.ConfigController, key)
	log.Infof(fmt.Sprintf("manager, config '%s' removed", key))

	return controller, nil
}
