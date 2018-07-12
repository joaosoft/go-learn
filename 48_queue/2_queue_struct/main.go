package main

import (
	common "go-learn/48_queue/2_queue_struct/common/queue"
	"go-learn/48_queue/2_queue_struct/controllers"

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
