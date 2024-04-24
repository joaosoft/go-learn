/////////////////////////////////
// Mutexes
// Go Playground: https://play.golang.org/p/xVBcUn-CS_4
/////////////////////////////////

//** IMPORTANT **//
// Run this program on your local machine (not in Go Playground)
// Execute: go run main.go

// Use Go Race Detector to check that there is no Data Race
// Execute: go run -race main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	// declaring a shared value
	var n int = 0

	// 1.
	// Declaring a mutex. It's available in sync package
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)

			// 2.
			// Lock the access to the shared value
			m.Lock()
			n++

			// 3.
			// Unlock the variable after it's incremented
			m.Unlock()

			wg.Done()
		}()

		// Doing the same for the 2nd goroutine
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()
		}()

	}

	wg.Wait()

	// printing the final value of n
	fmt.Println(n) // the final final of n will be always 0
}