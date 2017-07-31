package pm

// Config ... config interface
type IConfig interface {
	Get(key string) interface{}
}

// ConfigController ... config structure
type ConfigController struct {
	path   string
	config IConfig
	obj    interface{}
}

// NewConfig ... create a new ConfigController
func NewConfig(path string, config IConfig, obj interface{}) IConfig {

	return &ConfigController{
		path:   path,
		config: config,
		obj:    obj,
	}
}

// Get ... get a configuration by key
func (instance *ConfigController) Get(key string) interface{} {
	return instance.config.Get(key)
}

// Reload ... reload the configuration file
func (instance *ConfigController) Reload(key string) error {
	return nil
}
