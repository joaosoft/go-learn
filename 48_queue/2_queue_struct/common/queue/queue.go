package queue

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
)

func init() {

}

type Queue struct {
	shutdownChannel       chan bool
	addWorkChannel        chan IWork
	workListChannel       chan []IWork
	workChannelBufferSize int
	timeoutNotifyChannel  chan bool
	controller            IController
}

func NewQueue(shutdownChannelIn chan bool, workChannelBufferSize int, controller IController) *Queue {
	log.Infof("->NewQueue()")

	queue := Queue{
		shutdownChannel:       shutdownChannelIn,
		addWorkChannel:        make(chan IWork),
		workChannelBufferSize: workChannelBufferSize,
		workListChannel:       make(chan []IWork, workChannelBufferSize),
		timeoutNotifyChannel:  make(chan bool),
		controller:            controller,
	}

	go bulkBufferHandler(queue)
	go bufferLoopTime(queue)

	return &queue
}

func (queue *Queue) AddWork(work IWork) error {
	log.Infof("AddWork()")

	queue.addWorkChannel <- work

	return nil
}

// Buffer Handler
func bulkBufferHandler(queue Queue) {
	log.Infof("bulkBufferHandler()")

	bulkBufferSize := 0
	var bulkBuffer []IWork

	flush := func(buffer []IWork) {
		tempBuffer := make([]IWork, len(buffer))
		copy(tempBuffer, buffer)

		queue.workListChannel <- tempBuffer

		bulkBuffer = bulkBuffer[:0]
		bulkBufferSize = 0
	}

	for {
		select {
		case data := <-queue.addWorkChannel:
			log.Infof("data := <-channel")
			if bulkBufferSize > 100 {
				fmt.Printf("[BUFFER] Buffer full: flushing")
				flush(bulkBuffer)
			}

			bulkBuffer = append(bulkBuffer, data)
			bulkBufferSize++
			fmt.Println("SIZE:", bulkBufferSize)

		case <-queue.timeoutNotifyChannel:
			log.Infof("<-timeoutNotifyChannel ->", bulkBufferSize)
			if bulkBufferSize > 0 {
				fmt.Printf("[BUFFER] Timeout: flushing")
				flush(bulkBuffer)
			}

		case <-queue.shutdownChannel:
			if bulkBufferSize > 0 {
				fmt.Printf("[BUFFER] Shutdown: flushing")
				flush(bulkBuffer)
			}
			log.Infof("Shutdown service")
			return
		}
	}
}

func flushBulkCall(buffer []IWork, queue Queue) {
	//var obj []domain.BulkCall
	var has_error bool
	var err error
	var work []byte
	var works string

	log.Infof("flushBulkCall(buffer []IWork)")
	for _, v := range buffer {
		work, err = v.GetWork()

		if err != nil {
			has_error = true
		} else {
			fmt.Println("WORK:", string(work))
		}
		works = works + string(work)
	}
	queue.controller.Do([]byte(works))

	if has_error {
		queue.workListChannel <- buffer
	}

	log.Infof("!!!!WORK DONE!!!!")

}

func bufferLoopTime(queue Queue) {
	for {
		select {
		case buffer := <-queue.workListChannel:
			log.Infof("buffer := <-interCommChannel")
			flushBulkCall(buffer, queue)

		case <-time.After(time.Second * 20):
			log.Infof("<-time.After(time.Second * 20)")
			queue.timeoutNotifyChannel <- true
		}
	}
}
