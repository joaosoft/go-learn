
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"golang-learn/examples/46_workers/common/workers"
	"golang-learn/examples/46_workers/worker"
	"golang-learn/examples/46_workers/config"
	"golang-learn/examples/46_workers/common/config"
	"golang-learn/examples/46_workers/controllers"
	"github.com/labstack/gommon/log"
)

var (
	NWorkers = flag.Int("n", 4, "The number of worker to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
)

var configuration config.Configuration

func init() {
	if err := common_config.LoadConfigFromFile("config.json", &configuration); err != nil {
		log.Error("Error loading config.json: ", err)
		os.Exit(0)
	}
}

func main() {
	stop := make(chan bool)
	// Parse the command-line flags.
	flag.Parse()

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	workers.StartDispatcher(*NWorkers)

	controller := controllers.SomeController{}

	// DO TEST
	request := workers.WorkRequest{
		Controller: &worker.WorkerController {
			SomeController: controller,
		},
	}

	workers.SimpleCollector(request)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", workers.HttpCollector)

	// Start the HTTP server!
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}

	<- stop
}