package workers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)

func HttpCollector(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var work WorkRequest
	err = json.Unmarshal(body, &work)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Push the work onto the queue.
	WorkQueue <- work
	fmt.Println("Work request queued")

	// And let the user know their work request was created.
	w.WriteHeader(http.StatusCreated)
	return
}

func SimpleCollector(work WorkRequest) error {
	// Push the work onto the queue.
	fmt.Println("Adding request to queue")
	WorkQueue <- work
	fmt.Println("Work request queued")

	return nil
}