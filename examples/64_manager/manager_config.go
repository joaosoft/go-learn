package pm

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/config"
)

// -------------- CONFIGURATION CLIENTS --------------
// NewJSONFile ... creates a new nsq producer
func (instance *manager) NewSimpleConfig(path string, file string, extension string) (IConfig, error) {
	return config.NewSimpleConfig(path, file, extension)
}

// -------------- CONFIGURATION --------------
// GetConfig ... get a config with key
func (instance *manager) GetConfig(key string) IConfig {
	return instance.configController[key]
}

// AddProcess ... add a config with key
func (instance *manager) AddConfig(key string, config IConfig) error {
	if instance.started {
		panic("manager, can not add config after start")
	}

	instance.configController[key] = &configController{
		path:   "",
		config: config}

	log.Debug(fmt.Sprintf("manager, config '%s' added", key))

	return nil
}

// RemConfig ... remove the config by bey
func (instance *manager) RemConfig(key string) (IConfig, error) {
	// get config
	controller := instance.configController[key]

	// delete config
	delete(instance.configController, key)
	log.Debug(fmt.Sprintf("manager, config '%s' removed", key))

	return controller, nil
}
