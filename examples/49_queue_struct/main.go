package main

import (
	"github.com/labstack/gommon/log"
	"golang-learn/examples/49_queue_struct/controllers"
	common "golang-learn/examples/49_queue_struct/common/queue"
)

func main() {
	log.Infof("JOB START")

	shutdownChannelIn := make(chan bool)
	workChannelBufferSize := 5
	queue := common.NewQueue(shutdownChannelIn, workChannelBufferSize)

	bytes := []byte(`a, b, c`)
	controller := controllers.NewController(bytes)
	queue.AddWork(controller)

	bytes = []byte(`d, e, f`)
	controller = controllers.NewController(bytes)
	queue.AddWork(controller)

	<- shutdownChannelIn

	log.Infof("JOB END")
}
