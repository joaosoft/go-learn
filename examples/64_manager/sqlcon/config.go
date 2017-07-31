package sqlcon

// Config ...
type Config struct {
	Endpoint           string `json:"endpoint"`
	Driver             string `json:"driver"`
	MaxIdleConnections int    `json:"max_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections"`
}

// NewConfig ... created a new sqlpool config
func NewConfig(endpoint string, driver string, maxIdleConnections int, maxOpenConnections int) *Config {
	config := &Config{
		Endpoint:           endpoint,
		Driver:             driver,
		MaxIdleConnections: maxIdleConnections,
		MaxOpenConnections: maxOpenConnections,
	}

	return config
}
