/*
Change the program from Exercise #4 and calculate the square of all values between 1 and 50 included
using an anonymous function.



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
*/

package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 1; i <= 100; i++ {
		go func(x int) {
			ch <- x * x
		}(i)
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(<-ch)
	}
}
