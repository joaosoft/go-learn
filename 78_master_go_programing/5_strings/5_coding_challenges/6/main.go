/*
Consider the following string declaration:s := "你好 Go!"

1. Convert the string to a rune slice.

2. Print out the rune slice

3. Iterate over the rune slice and print out each index and rune in the rune slice
*/

package main

import "fmt"

func main() {
	s := "你好 Go!"

	// converting string to rune slice
	r := []rune(s)

	// printing out the rune slice
	fmt.Printf("%#v\n", r)

	// iterating over the rune slice and printing out each index and rune in the rune slice
	for i, v := range r {
		fmt.Printf("%d -> %c\n", i, v)
	}

}
