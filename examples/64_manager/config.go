package pm

// Config ... config interface
type IConfig interface {
	Get(key string) interface{}
}

// configController ... config structure
type configController struct {
	path   string
	config IConfig
	obj    interface{}
}

// NewConfig ... create a new configController
func NewConfig(path string, config IConfig, obj interface{}) IConfig {

	return &configController{
		path:   path,
		config: config,
		obj:    obj,
	}
}

// Get ... get a configuration by key
func (instance *configController) Get(key string) interface{} {
	return instance.config.Get(key)
}

// Reload ... reload the configuration file
func (instance *configController) Reload(key string) error {
	return nil
}
