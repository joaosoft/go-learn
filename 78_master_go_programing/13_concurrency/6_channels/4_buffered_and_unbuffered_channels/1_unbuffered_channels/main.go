/////////////////////////////////
// Unbuffered Channels
// Go Playground: https://play.golang.org/p/_44csjQDJvM
/////////////////////////////////

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) //unbuffered channel

	// Launching a goroutine
	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1) // calling the anonymous func and passing c1 as argument

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	// we sleep for a second to give time to the goroutine to finish
	time.Sleep(time.Second)

	// After running the program we notice that the sender (the func goroutine) blocks on the channel
	// until the receiver (the main goroutine) receives the data from the channel.
}

//** EXPECTED OUTPUT: **//
// main goroutine sleeps for 2 seconds
// func goroutine starts sending data into the channel
// main goroutine starts receiving data
// main goroutine received data: 10
// func goroutine after sending data into the channel