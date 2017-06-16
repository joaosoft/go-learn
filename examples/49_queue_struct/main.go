package main

import (
	"github.com/labstack/gommon/log"
	common "golang-learn/examples/49_queue_struct/common/queue"
)

func main() {
	log.Infof("JOB START")

	shutdownChannelIn := make(chan bool)
	workChannelBufferSize := 5
	queue := common.NewQueue(shutdownChannelIn, workChannelBufferSize)

	bytes := []byte(`a, b, c`)
	controller := common.NewController(bytes)
	queue.AddWork(*controller)

	bytes = []byte(`d, e, f`)
	controller = common.NewController(bytes)
	queue.AddWork(*controller)

	<- shutdownChannelIn

	log.Infof("JOB END")
}
