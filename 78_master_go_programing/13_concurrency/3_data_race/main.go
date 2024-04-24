/////////////////////////////////
// Data Races in Go
// Go Playground:https://play.golang.org/p/AiSjfvn4O3T
/////////////////////////////////

//** IMPORTANT **//
// Run this program on your local machine (not in Go Playground)
// Execute: go run main.go

// Use Go Race Detector to detect the Data Race
// Execute: go run -race main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100

	// declaring a WaitGroup to synchronize the goroutines with the main function.
	var wg sync.WaitGroup

	// adding 200 goroutines to the WaitGroup
	wg.Add(gr * 2)

	// declaring a shared value
	var n int = 0

	// starting 200 goroutines
	for i := 0; i < gr; i++ {

		// each goroutine is an anonymous function
		go func() {

			// introducing some artificial time
			time.Sleep(time.Second / 10)

			// increment the shared value
			n++

			// notifying the WaitGroup that the goroutine is done
			wg.Done()
		}()

		// goroutine that decrements the shared value
		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()

	}
	// waiting for the goroutines to terminate.
	wg.Wait()

	//  printing the final value of n
	fmt.Println(n) // it will have another value for each program execution -> DATA RACE
}