package main

import (
	"go-learn/48_queue/1_queue_map/common/queue"

	"github.com/labstack/gommon/log"
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

	<-shutdownChannelIn

	log.Infof("JOB END")
}
