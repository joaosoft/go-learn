package main

import (
	common "golang-learn/50_queue_controller/common/queue"
	"golang-learn/50_queue_controller/controllers"
	"golang-learn/50_queue_controller/repositories"

	"github.com/labstack/gommon/log"
)

func main() {
	log.Infof("JOB START")

	shutdownChannelIn := make(chan bool)
	workChannelBufferSize := 5
	repository := repositories.Repository{}
	controller := controllers.NewController(repository)
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
