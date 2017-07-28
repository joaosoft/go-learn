package nsq

// Config ...
type Config struct {
	Lookupd      []string `mapstructure:"lookupd"`
	Nsqd         []string `mapstructure:"nsqd"`
	Topic        string   `mapstructure:"topic"`
	Channel      string   `mapstructure:"channel"`
	RequeueDelay int64    `mapstructure:"requeue_delay"`
	MaxInFlight  int      `mapstructure:"max_in_flight"`
	MaxAttempts  uint16   `mapstructure:"max_attempts"`
	AutoRespond  bool     `mapstructure:"auto_respond"`
}

// NewConfig ... created a new nsq config
func NewConfig(topic, channel string, lookupd, nsqd []string) *Config {
	config := &Config{
		Topic:        topic,
		Channel:      channel,
		Lookupd:      lookupd,
		Nsqd:         nsqd,
		RequeueDelay: 30,
		MaxInFlight:  5,
	}

	return config
}
