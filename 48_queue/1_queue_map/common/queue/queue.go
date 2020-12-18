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
	addWorkTopChannel     chan map[string]interface{}
	workListChannel       chan []map[string]interface{}
	workChannelBufferSize int
	timeoutNotifyChannel  chan bool
}

func NewQueue(shutdownChannelIn chan bool, workChannelBufferSize int) *Queue {
	log.Infof("->NewQueue()")

	queue := Queue{
		shutdownChannel:       shutdownChannelIn,
		addWorkTopChannel:     make(chan map[string]interface{}),
		workChannelBufferSize: workChannelBufferSize,
		workListChannel:       make(chan []map[string]interface{}, workChannelBufferSize),
		timeoutNotifyChannel:  make(chan bool),
	}

	go bulkBufferHandler(queue)
	go bufferLoopTime(queue)

	return &queue
}

func (queue *Queue) AddWork(work map[string]interface{}) error {
	log.Infof("AddWork()")

	queue.addWorkTopChannel <- work

	return nil
}

// Buffer Handlers
func bulkBufferHandler(queue Queue) {
	log.Infof("bulkBufferHandler()")

	bulkBufferSize := 0
	var bulkBuffer []map[string]interface{}

	flush := func(buffer []map[string]interface{}) {
		tempBuffer := make([]map[string]interface{}, len(buffer))
		copy(tempBuffer, buffer)

		queue.workListChannel <- tempBuffer

		bulkBuffer = bulkBuffer[:0]
		bulkBufferSize = 0
	}

	for {
		select {
		case data := <-queue.addWorkTopChannel:
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

func flushBulkCall(buffer []map[string]interface{}, queue Queue) {
	//var obj []domain.BulkCall
	var err error

	log.Infof("flushBulkCall(buffer []map[string]interface{})")
	for _, v := range buffer {
		fmt.Println(v)
	}

	if err != nil {
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
