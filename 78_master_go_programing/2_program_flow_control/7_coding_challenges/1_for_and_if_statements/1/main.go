/*
Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.
*/

package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		if i%7 == 0 { // if i is divisible by 7
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}
