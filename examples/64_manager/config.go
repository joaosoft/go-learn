package pm

// Config ... config interface
type IConfig interface {
	GetConfig(key string) (interface{}, error)
	AddConfig(key string, config interface{}) (interface{}, error)
	RemConfig(key string) (interface{}, error)
}

// Config ... config structure
type Config struct {
	configs map[string]IConfig
}

// NewConfig ... create a new config
func NewConfig() *Config {
	return &Config{
		configs: make(map[string]IConfig),
	}
}

// GetConfig ... get a config with key
func (instance *Config) GetConfig(key string) (IConfig, error) {
	return instance.configs[key], nil
}

// AddConfig ... add a config with key
func (instance *Config) AddConfig(key string, config IConfig) error {
	instance.configs[key] = config

	return nil
}

// RemConfig ... remove the config by bey
func (instance *Config) RemConfig(key string) (IConfig, error) {
	// get config
	config := instance.configs[key]

	// delete config
	delete(instance.configs, key)

	return config, nil
}
