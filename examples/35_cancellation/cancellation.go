// testes

package main

import (
	"fmt"
	"os"
)

func main() {
	//Sometimes we need to instruct a goroutine to stop what it is doing, for example,
	// 	in a web server performing a computation on behalf of a client that has disconnected.
	// There is no way for one goroutine to terminate another directly, since that
	// would leave all its shared variables in undefined states.
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_cancellation()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
var done = make(chan struct{})

func _cancellation() {
	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()
}
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
