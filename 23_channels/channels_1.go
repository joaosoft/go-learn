package main

import (
	"fmt"
)

func main() {
	// 	ch <- x  // a send statement
	// x = <-ch // a receive expression in an assignment statement
	// <-ch     // a receive statement; result is discarded

	// ch = make(chan int)    // unbuffered channel
	// ch = make(chan int, 0) // unbuffered channel
	// ch = make(chan int, 3) // buffered channel with capacity 3

	// Unbuffered Channels
	// A send operation on an unbuffered channel blocks the sending goroutine until another goroutine
	// executes a corresponding receive on the same channel, at which point the value is transmitted and
	// both goroutines may continue. Conversely, if the receive operation was attempted first, the receiving
	// goroutine is blocked until another goroutine performs a send on the same channel.

	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_channels()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func _channels() {
	fmt.Println("_channels()")
	// CANAIS
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // ch has type 'chan int'
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// Canais Bufferizados
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Range e Close
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}
}
