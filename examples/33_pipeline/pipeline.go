package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_pipeline1()
	_pipeline2()
	_pipeline3()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions

func _pipeline1() {
	//As you might expect, the program prints the infinite series of squares 0, 1, 4, 9, and so on.
	// Pipelines like this may be found in long-running server programs where channels are used
	// for lifelong communication between goroutines containing infinite loops. But what if we want
	// to send only a finite number of values through the pipeline?

	//If the sender knows that no further values will ever be sent on a channel, it is useful to
	// communicate this fact to the receiver goroutines so that they can stop waiting.
	// This is accomplished by closing the channel using the built-in close function: close(naturals)
	fmt.Println("_pipeline()")
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

func _pipeline2() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

func _pipeline3() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
