/////////////////////////////////
// Goroutines and Channels
// Go Playground: https://play.golang.org/p/cNsq9YcFvW7
/////////////////////////////////

package main

import (
	"fmt"
	"strings"
)

// declaring a function that computes the factorial of n
func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	// sending the factorial value into the channel
	c <- f
}

func main() {

	// declaring and initializing a channel of type `chan int`
	ch := make(chan int)

	// defer closing the channel
	defer close(ch)

	// launching a goroutine
	go factorial(5, ch)

	// main() is waiting for a value to come from the channel
	// this is called a `blocking call`

	f := <-ch // receiving the value from the channel in a new variable
	fmt.Println("5 factorial =", f)

	// Spawning 20 goroutines that calculate the factorial
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Printf("Factorial of %d: %d\n", i, f)
	}

	fmt.Println(strings.Repeat("#", 10))

	// Spawning another 10 goroutines this time as anonymous functions
	for i := 5; i < 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}

			// sending the value f into the channel
			c <- f
		}(i, ch)
		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}
}