/*
Change the code from the previous exercise and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.
*/
package main

import "fmt"

func main() {

	count := 0
	for i := 1; i <= 50; i++ {
		if i%7 != 0 { // if i is not divisible by 7
			continue
		}
		fmt.Printf("%d ", i)
		count++

		if count == 3 { // if i've already found 3 numbers, then break
			break
		}

	}
	fmt.Println("")
}
