package main

import (
	common "golang-learn/49_queue_struct/common/queue"
	"golang-learn/49_queue_struct/controllers"

	"github.com/labstack/gommon/log"
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

	<-shutdownChannelIn

	log.Infof("JOB END")
}
