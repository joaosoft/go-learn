/*
Create a goroutine named power() that has one parameter of type int,
calculates the square value of its parameter and then sends the result into a channel.

In the main function launch 50 goroutines that calculate the square values
of all numbers between 1 and 50 included.

Print out the square values.

A square(or raising to power 2) is the result of multiplying a number
by itself. e.g., 25 is the square of 5.
*/

package main

import (
	"fmt"
)

func power(x int, c chan int) {
	c <- x * x
}

func main() {
	ch := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-ch)
	}

}
