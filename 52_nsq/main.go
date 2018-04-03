package main

import (
	common "go-learn/52_nsq/common/config"
	"go-learn/52_nsq/common/nsq"
	"go-learn/52_nsq/config"
	"go-learn/52_nsq/controllers"
	"go-learn/52_nsq/interactors"

	"github.com/labstack/gommon/log"

	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var _configuration config.Configuration
var _controllers []controllers.Controller

func init() {
	if err := common.LoadConfigFromFile("config", &_configuration); err != nil {
		log.Error("Error loading config: ", err)
		os.Exit(0)
	}
	fmt.Println("Configuration: ", _configuration)
}

func main() {
	controlChannel := make(chan int)

	for _, topic := range _configuration.NSQ.Topics {
		controller := controllers.Controller{
			Consumer:   nsq.NewNSQConsumer(topic.Topic, topic.Channel, _configuration.NSQ),
			Interactor: interactors.Interactor{},
		}

		if err := controller.Start(controlChannel); err != nil {
			log.Errorf("Service not started for [topic:%s][channel:%s]", topic.Topic, topic.Channel)
			continue
		}

		_controllers = append(_controllers, controller)
	}

	log.Infof("Service started: %d", len(_controllers))

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	for i := 0; i < len(_controllers); i++ {
		_controllers[i].Stop()
	}

	for i := 0; i < len(_controllers); i++ {
		<-controlChannel
	}

	log.Info("Service terminated!")
}
