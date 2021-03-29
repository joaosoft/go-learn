/////////////////////////////////
// The select Statement
// Go Playground: https://play.golang.org/p/6qRtwfSPzef
/////////////////////////////////

package main

import (
	"fmt"
	"time"
)

func main() {

	// retrieving the current timestamp in milliseconds
	s := time.Now().UnixNano() / 1000000

	// The `select` statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case.
	// Select is only used with channels.

	// declaring 2 channels
	c1 := make(chan string)
	c2 := make(chan string)

	// starting the first goroutine using an anonymous function
	go func() {
		time.Sleep(2 * time.Second)

		// sending a message into the channel
		c1 <- "Hello!"
	}()

	// starting the second goroutine using an anonymous function
	go func() {
		time.Sleep(1 * time.Second)

		// sending a message into the channel
		c2 <- "Salut!"
	}()

	// using select to wait on both goroutines
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)

		}
	}

	e := time.Now().UnixNano() / 1_000_000
	// total execution time is only ~2 seconds since the goroutines executed concurrently.
	fmt.Println(e - s)

	// Basic sends and receives on channels are blocking.
	// However, we can use `select` with a `default` clause to implement non-blocking channels.
}
