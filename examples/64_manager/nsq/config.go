package nsq

// Config ...
type Config struct {
	Lookupd      []string `json:"lookupd"`
	Nsqd         []string `json:"nsqd"`
	Topic        string   `json:"topic"`
	Channel      string   `json:"channel"`
	RequeueDelay int64    `json:"requeue_delay"`
	MaxInFlight  int      `json:"max_in_flight"`
	MaxAttempts  uint16   `json:"max_attempts"`
	AutoRespond  bool     `json:"auto_respond"`
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
