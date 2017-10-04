package workers

import "fmt"

var WorkerQueue chan chan WorkRequest
var workers []*Worker

func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the worker' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	// Now, create all of our worker.
	for i := 0; i<nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
		workers = append(workers, &worker)
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}