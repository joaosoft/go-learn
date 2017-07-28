package main

import (
	"fmt"
	nsqlib "github.com/nsqio/go-nsq"
	"golang-learn/examples/64_manager"
	"golang-learn/examples/64_manager/nsq"
)

type DummyHandler struct{}

func (instance *DummyHandler) HandleMessage(msg *nsqlib.Message) error {
	fmt.Println("HELLO!")
	return nil
}

func main() {

	manager := pm.NewManager()

	config := &nsq.Config{
		Topic:   "topic_1",
		Channel: "channel_2",
		Lookupd: []string{"http://localhost:4151"},
	}
	nsqConsumer, _ := manager.NewNSQConsumer(config, &DummyHandler{})
	nsqProducer, _ := manager.NewNSQProducer(config)

	manager.AddProcess("teste_1", nsqConsumer)
	manager.AddProcess("teste_2", nsqProducer)

	manager.Start()
}
