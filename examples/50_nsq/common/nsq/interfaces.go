package nsq

type INSQConsumer interface {
	Start(control chan int) error
	Stop() error
	SetHandler(handler INSQHandler, concurrency uint) error
}

type INSQHandler interface {
	HandleMessage(data []byte) error
}
