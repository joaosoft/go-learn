package nsq

import (
	"github.com/labstack/gommon/log"
	nsqlib "github.com/nsqio/go-nsq"
	"time"
)

// IProducer poducer interface
type IProducer interface {
	Start() error
	Stop() error
	Publish(topic string, body []byte, maxRetries int) error
	Ping() error
}

// Producer ... producer structure
type Producer struct {
	Client *nsqlib.Producer
	config *Config
}

// NewProducer ... creates a new producer
func NewProducer(config *Config) (IProducer, error) {
	var addr string

	// Creating nsq configuration
	nsqConfig := nsqlib.NewConfig()
	nsqConfig.MaxAttempts = config.MaxAttempts
	nsqConfig.DefaultRequeueDelay = time.Duration(config.RequeueDelay) * time.Second
	nsqConfig.MaxInFlight = config.MaxInFlight
	nsqConfig.ReadTimeout = 120 * time.Second

	if config.Lookupd != nil && len(config.Lookupd) > 0 {
		addr = config.Lookupd[0]
	} else {
		panic("nsq producer, no NSQ address to connect to")
	}

	log.Infof("nsq producer, connecting to %s", addr)

	nsqProducer, err := nsqlib.NewProducer(addr, nsqConfig)
	if err != nil {
		panic(err)
	}

	producer := &Producer{
		Client: nsqProducer,
		config: config,
	}

	return producer, nil
}

// Publish ... published a message on to a topic.
func (instance *Producer) Publish(topic string, body []byte, maxRetries int) error {
	var err error

	for count := 0; count < maxRetries; count++ {
		if err = instance.Client.Publish(topic, body); err == nil {
			return nil
		}
	}

	return err
}

// Start ... start the producer
func (instance *Producer) Start() error {
	return nil
}

// Stop ... stop the producer
func (instance *Producer) Stop() error {
	log.Infof("nsq producer, stopping")
	instance.Client.Stop()
	log.Infof("nsq producer, stopped")
	return nil
}

// Ping ... pings the producer
func (instance *Producer) Ping() error {
	return instance.Client.Ping()
}
