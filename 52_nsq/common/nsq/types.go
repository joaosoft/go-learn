package nsq

import "time"

type NSQ struct {
	Nsqd   []string `json:"nsqds"`
	Topics []struct {
		Topic   string `json:"topic"`
		Channel string `json:"channel"`
	} `json:"topics"`
	RequeueDelay time.Duration `json:"requeueDelay"`
	MaxBuffer    int           `json:"maxBuffer"`
	Lookupd      []string      `json:"lookupds"`
}

// SetRequeueDelay ...
func (config *NSQ) SetRequeueDelay(delay time.Duration) {
	config.RequeueDelay = delay
}

// SetMaxBuffer ...
func (config *NSQ) SetMaxBuffer(size int) {
	config.MaxBuffer = size
}

// HasLookupdAddresses returns true if there are lookupd addresses configured.
func (config *NSQ) HasLookupdAddresses() bool {
	return config.Lookupd != nil && len(config.Lookupd) > 0
}

// GetLookupdAddresses ...
func (config *NSQ) GetLookupdAddresses() []string {
	return config.Lookupd
}

// HasNodeAddresses returns true if there are NSQDaemon addresses configured.
func (config *NSQ) HasNodeAddresses() bool {
	return config.Nsqd != nil && len(config.Nsqd) > 0
}

// GetNodeAddresses ...
func (config *NSQ) GetNodeAddresses() []string {
	return config.Nsqd
}

// GetDefaultRequeueDelay ...
func (config *NSQ) GetDefaultRequeueDelay() time.Duration {
	return config.RequeueDelay
}

// GetInboundBufferSize ...
func (config *NSQ) GetInboundBufferSize() int {
	return config.MaxBuffer
}

// GetProcessTimeout sets the timeout for the application to process a message.
func (config *NSQ) GetProcessTimeout() time.Duration {
	return 120 * time.Second
}
