/////////////////////////////////
// Spawning Goroutines and Synchronize them using Waitgroups
// Go Playground: https://play.golang.org/p/zj_7v820Ipe
/////////////////////////////////

// The pattern to use sync.WaitGroup is:
// 1. Create a new variables of a `sync.WaitGroup` (wg)
// 2. Call `wg.Add(n)` where `n` is the number of goroutines to wait for
// 3. Execute `defer wg.Done()` in each goroutine to indicate to the WaitGroup that the goroutine has finished executing
// 4. Call `wg.Wait()` in main() where we want to block.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Declaring two functions: f1 and f2
func f1(wg *sync.WaitGroup) { // wg is passed as a pointer
	fmt.Println("f1(goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		// sleep for a second to simulate an expensive task.
		time.Sleep(time.Second)

	}
	fmt.Println("f1 execution finished")

	//3.
	// Before exiting, call wg.Done() in each goroutine
	// to indicate to the WaitGroup that the goroutine has finished executing.
	wg.Done()
	//or:
	// (*wg).Done()
}

func f2() {
	fmt.Println("f2 execution started")
	time.Sleep(time.Second)

	for i := 5; i < 8; i++ {
		fmt.Println("f2(), i=", i)

	}
	fmt.Println("f2 execution finished")

}

func main() {
	fmt.Println("main execution started")

	// 1.
	// Create a new instance of sync.WaitGroup (we’ll call it symply wg)
	// This WaitGroup is used to wait for all the goroutines that have been launched to finish.
	var wg sync.WaitGroup

	// 2.
	// Call wg.Add(n) method before attempting to
	// launch the go routine.
	wg.Add(1) //  n which is 1 is the number of goroutines to wait for

	// Launching a goroutine
	go f1(&wg) // it takes in a pointer to sync.WaitGroup

	// No. of running goroutines
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => 2

	// calling other functions:
	f2()

	// 4.
	// Finally, we call wg.Wait()to block the execution of main() until the goroutines
	// in the WaitGroup have successfully completed.
	wg.Wait()

	fmt.Println("main execution stopped")
}

// Run the program: go run main.go

//** EXPECTED OUTPUT: **//
// main execution started
// No. of Goroutines: 2
// f2 execution started
// f1(goroutine) execution started
// f1, i= 0
// f2(), i= 5
// f2(), i= 6
// f2(), i= 7
// f2 execution finished
// f1, i= 1
// f1, i= 2
// f1 execution finished
// main execution stopped