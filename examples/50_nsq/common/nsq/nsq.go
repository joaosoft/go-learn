package nsq

import (
	"github.com/labstack/gommon/log"
	nsqlib "github.com/nsqio/go-nsq"
)

type NSQConsumer struct {
	Connection       *nsqlib.Consumer
	Started          bool
	Handler          INSQHandler
	NSQConfiguration NSQ
}

// Creates and returns a worker_add NSQConsumer based on configurations sent by param
func NewNSQConsumer(topic string, channel string, nsqConfig NSQ) *NSQConsumer {
	log.Infof("creating consumer [topic:%s][channel:%s]", topic, channel)

	cfg := newNSQConfig(nsqConfig)

	nsqConsumer, err := nsqlib.NewConsumer(topic, channel, cfg)
	if err != nil {
		panic(err)
	}

	consumer := NSQConsumer{
		Connection:       nsqConsumer,
		Started:          false,
		NSQConfiguration: nsqConfig,
	}

	nsqConsumer.SetLogger(&consumer, nsqlib.LogLevelWarning)

	log.Infof("consumer [topic:%s][channel:%s] created", topic, channel)

	return &consumer
}

// Set the message handler
func (consumer *NSQConsumer) SetHandler(handler INSQHandler, concurrency uint) error {
	consumer.Handler = handler
	consumer.Connection.AddConcurrentHandlers(consumer, int(concurrency))

	return nil
}

// NSQ Handler implementation
// Handle the message and send to a custom handler
func (consumer *NSQConsumer) HandleMessage(nsqMsg *nsqlib.Message) error {
	nsqMsg.DisableAutoResponse()

	if err := consumer.Handler.HandleMessage(nsqMsg.Body); err != nil {
		return err
	}

	return nil
}

// INSQConsumer implementation
// Connect to NSQD and Lookup addresses and start listening the channel
func (consumer *NSQConsumer) Start(control chan int) error {
	consumer.Connection.ChangeMaxInFlight(consumer.NSQConfiguration.GetInboundBufferSize())

	if consumer.NSQConfiguration.HasLookupdAddresses() {
		consumer.Started = true
		for _, addr := range consumer.NSQConfiguration.GetLookupdAddresses() {
			log.Infof("NSQ consumer connecting to %s", addr)
		}
		if err := consumer.Connection.ConnectToNSQLookupds(consumer.NSQConfiguration.GetLookupdAddresses()); err != nil {
			return err
		}
	}
	if consumer.NSQConfiguration.HasNodeAddresses() {
		consumer.Started = true
		for _, addr := range consumer.NSQConfiguration.GetNodeAddresses() {
			log.Infof("NSQ consumer connecting to %s", addr)
		}
		if err := consumer.Connection.ConnectToNSQDs(consumer.NSQConfiguration.GetNodeAddresses()); err != nil {
			return err
		}
	}

	if !consumer.Started {
		panic("failed to start consumer")
	}

	go func() {
		control <- <-consumer.Connection.StopChan
	}()

	return nil
}

// INSQConsumer implementation
// Stop the connection between the consumer and channel
func (consumer *NSQConsumer) Stop() error {
	consumer.Connection.Stop()
	consumer.Started = false

	return nil
}

func newNSQConfig(config NSQ) *nsqlib.Config {
	cfg := nsqlib.NewConfig()

	if config.GetDefaultRequeueDelay() > 0 {
		cfg.DefaultRequeueDelay = config.GetDefaultRequeueDelay()
	}
	if config.GetInboundBufferSize() > 0 {
		cfg.MaxInFlight = config.GetInboundBufferSize()
	}
	if config.GetProcessTimeout() > 0 {
		cfg.ReadTimeout = config.GetProcessTimeout()
	}

	return cfg
}

// Output implements the common.Logger interface Output method
func (consumer *NSQConsumer) Output(level int, msg string) error {
	switch level {
	case int(nsqlib.LogLevelDebug):
		log.Debug(msg)
	case int(nsqlib.LogLevelInfo):
		log.Debug(msg)
	case int(nsqlib.LogLevelWarning):
		log.Warn(msg)
	case int(nsqlib.LogLevelError):
		log.Error(msg)
	}
	return nil
}
