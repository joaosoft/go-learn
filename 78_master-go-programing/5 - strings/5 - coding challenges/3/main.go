/*
Consider the following string declaration: s1 := "țară means country in Romanian"

1. Iterate over the string and print out byte by byte

2. Iterate  over the string and print out rune by rune and the byte index where the rune starts in the string
*/

package main

import "fmt"

func main() {
	s1 := "țară means country in Romanian"

	// iterating over the string and print out byte by byte
	fmt.Printf("Bytes in string: ")
	for i := 0; i < len(s1); i++ {

		fmt.Printf("%v ", s1[i])
	}

	fmt.Println()

	// iterating over the string and print out rune by rune
	// and the byte index where the rune starts in the string
	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}

	fmt.Println()

}
