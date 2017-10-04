package main

import (
	"github.com/labstack/gommon/log"
	"golang-learn/examples/48_queue_map/common/queue"
)

func main() {
	log.Infof("JOB START")

	shutdownChannelIn := make(chan bool)
	workChannelBufferSize := 5
	queue := queue.NewQueue(shutdownChannelIn, workChannelBufferSize)

	work := make(map[string]interface{})
	work["1"] = 1
	work["2"] = 2
	queue.AddWork(work)

	work = make(map[string]interface{})
	work["1"] = 3
	work["2"] = 4
	queue.AddWork(work)

	<- shutdownChannelIn

	log.Infof("JOB END")
}
