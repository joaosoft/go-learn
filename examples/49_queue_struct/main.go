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
	controller := controllers.NewController()
	queue := common.NewQueue(shutdownChannelIn, workChannelBufferSize, controller)

	bytes := []byte(`a, b, c`)
	work := controllers.NewWork(bytes)
	queue.AddWork(work)

	bytes = []byte(`d, e, f`)
	work = controllers.NewWork(bytes)
	queue.AddWork(work)

	<- shutdownChannelIn

	log.Infof("JOB END")
}
