package gateway

// Config ...
type Config struct {
	Host string `json:"host"`
}

// NewConfig ... created a new gateway config
func NewConfig(host string) *Config {
	config := &Config{
		Host: host,
	}

	return config
}
