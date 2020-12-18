package nsq

import (
	nsqlib "github.com/nsqio/go-nsq"
	"github.com/stretchr/testify/mock"
)

type NSQConsumerMock struct {
	mock.Mock
}

func (consumer *NSQConsumerMock) SetHandler(handler INSQHandler, concurrency uint) error {
	args := consumer.Called(handler, concurrency)

	return args.Error(0)
}

// NSQ Handlers implementation
func (consumer *NSQConsumerMock) HandleMessage(nsqMsg *nsqlib.Message) error {
	args := consumer.Called(nsqMsg)

	return args.Error(0)
}

// INSQConsumer implementation
func (consumer *NSQConsumerMock) Start(control chan int) error {
	args := consumer.Called(control)

	return args.Error(0)
}

func (consumer *NSQConsumerMock) Stop() error {
	args := consumer.Called()

	return args.Error(0)
}
