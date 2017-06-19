package __webserver

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_webServer()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
// EXEMPLO WEBSERVER 1
// handler echoes the Path component of the requested URL.
func handler_server1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// EXEMPLO WEBSERVER 2
var mu sync.Mutex
var count int

// handler echoes the Path component of the requested URL.
func handler_server2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func _webServer() {
	// WEB SERVER 1
	http.HandleFunc("/a", handler_server1) // each request calls handler
	// WEB SERVER 2
	http.HandleFunc("/b", handler_server2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8123", nil))
}
