/*
There are some errors in the following Go program.
Try to identify the errors, change the code and run the program without errors.



package main

import (
	"fmt"
)

func main() {
	c := make(<-chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}

*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}
