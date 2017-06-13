package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_channels()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func doSomething(channel, quit chan int) {
	x, y := 10, 1
	for {
		select {
		case channel <- x:
			x, y = x-1, y+1
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func _channels() {
	channel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()
	doSomething(channel, quit)
}
