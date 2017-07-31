package web

// Config ...
type Config struct {
	Address string `json:"address"`
}

// NewConfig ... created a new nsq config
func NewConfig(address string) *Config {
	config := &Config{
		Address: address,
	}

	return config
}
