package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
)

type WorkRequest struct {
	Name  string
	Delay time.Duration
}

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)

func Collector(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the delay.
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check to make sure the delay is anywhere from 1 to 10 seconds.
	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	// Now, we retrieve the person's name from the request.
	name := r.FormValue("name")

	// Just do a quick bit of sanity checking to make sure the client actually provided us with a name.
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	// Now, we take the delay, and the person's name, and make a WorkRequest out of them.
	work := WorkRequest{Name: name, Delay: delay}

	// Push the work onto the queue.
	WorkQueue <- work
	fmt.Println("Work request queued")

	// And let the user know their work request was created.
	w.WriteHeader(http.StatusCreated)
	return
}

var WorkerQueue chan chan WorkRequest

// StartDispatcher ...
func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
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

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				// Receive a work request.
				fmt.Printf("worker%d: Received work request, delaying for %f seconds\n", w.ID, work.Delay.Seconds())

				time.Sleep(work.Delay)
				fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)

			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func main() {
	// Parse the command-line flags.
	flag.Parse()

	// Start the dispatcher.
	fmt.Printf("Starting the dispatcher (workers: %d)", *NWorkers)
	StartDispatcher(*NWorkers)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", Collector)

	// Start the HTTP server!
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
