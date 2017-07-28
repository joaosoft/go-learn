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

	// MANAGER
	manager := pm.NewManager()
	config := &nsq.Config{
		Topic:   "topic_1",
		Channel: "channel_2",
		Lookupd: []string{"http://localhost:4151"},
	}

	// PROCESSES
	nsqConsumer, _ := manager.NewNSQConsumer(config, &DummyHandler{})
	nsqProducer, _ := manager.NewNSQProducer(config)
	manager.AddProcess("teste_1", nsqConsumer)
	manager.AddProcess("teste_2", nsqProducer)

	// CONFIGURATION
	simpleConfig, _ := manager.NewSimpleConfig("/Users/joaoribeiro/workspace/go/src/golang-learn/examples/64_manager/run/", "config", "json")
	manager.AddConfig("teste_3", simpleConfig)
	fmt.Println("a: ", manager.GetConfig("teste_3").Get("a"))
	fmt.Println("caa: ", manager.GetConfig("teste_3").Get("c.ca.caa"))

	manager.Start()
}
