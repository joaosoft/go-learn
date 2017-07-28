package nsq

import (
	"fmt"
	"github.com/labstack/gommon/log"
	nsqlib "github.com/nsqio/go-nsq"
	"time"
)

// IConsumer costumer interface
type IConsumer interface {
	Start() error
	Stop() error
}

// Consumer ... costumer structure
type Consumer struct {
	client  *nsqlib.Consumer
	started bool
	handler IHandler
	config  *Config
}

// NewConsumer ... creates a new costumer
func NewConsumer(config *Config, handler IHandler) (IConsumer, error) {
	log.Infof("nsq consumer, creating instance [topic:%s][channel:%s]", config.Topic, config.Channel)

	// Creating nsq configuration
	nsqConfig := nsqlib.NewConfig()
	nsqConfig.MaxAttempts = config.MaxAttempts
	nsqConfig.DefaultRequeueDelay = time.Duration(config.RequeueDelay) * time.Second
	nsqConfig.MaxInFlight = config.MaxInFlight
	nsqConfig.ReadTimeout = 120 * time.Second

	nsqConsumer, err := nsqlib.NewConsumer(config.Topic, config.Channel, nsqConfig)
	nsqConsumer.AddHandler(handler)
	if err != nil {
		panic(err)
	}

	instance := &Consumer{
		client:  nsqConsumer,
		config:  config,
		started: false,
		handler: handler,
	}

	log.Infof("nsq consumer, instance [topic:%s][channel:%s] created", config.Topic, config.Channel)

	return instance, nil
}

// HandleMessage ... handle message queue
func (instance *Consumer) HandleMessage(nsqMsg *nsqlib.Message) error {
	nsqMsg.DisableAutoResponse()

	if err := instance.handler.HandleMessage(nsqMsg); err != nil {
		return err
	}

	return nil
}

// Start ... start's nsq costumer
func (instance *Consumer) Start() error {
	if instance.handler == nil {
		return fmt.Errorf("nsq consumer, no handler configured")
	}

	if instance.config.Lookupd != nil && len(instance.config.Lookupd) > 0 {
		instance.started = true
		for _, addr := range instance.config.Lookupd {
			log.Infof("nsq consumer, instance connecting to %s", addr)
		}
		if err := instance.client.ConnectToNSQLookupds(instance.config.Lookupd); err != nil {
			log.Error(err)
			return err
		}
	}
	if instance.config.Nsqd != nil && len(instance.config.Nsqd) > 0 {
		instance.started = true
		for _, addr := range instance.config.Nsqd {
			log.Infof("nsq consumer, instance connecting to %s", addr)
		}
		if err := instance.client.ConnectToNSQDs(instance.config.Nsqd); err != nil {
			return err
		}
	}

	if !instance.started {
		panic("nsq consumer, failed to start instance")
	}

	<-instance.client.StopChan

	instance.started = false

	return nil
}

// Stop ... stop's nsq costumer
func (instance *Consumer) Stop() error {
	log.Info("nsq consumer, stopping ")
	instance.client.Stop()
	instance.started = false
	log.Info("nsq consumer, stopped")

	return nil
}
