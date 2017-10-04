package __webserver

import (
	"fmt"
	"net/http"
	"runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "olá, estou a utilizar %s com CPU %s", runtime.GOOS, runtime.GOARCH)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8081", nil)
}
