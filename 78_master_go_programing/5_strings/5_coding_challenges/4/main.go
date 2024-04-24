/*
There are some errors in the following Go program.
Try to identify the errors, change the code and run the program without errors.

package main

import "fmt"

func main() {
	s1 := "Go is cool!"

	x := s1[0]
	fmt.Println(x)

	s1[0] = 'x'

	// printing the number of runes in the string
	fmt.Println(len(s1))

}
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Go is cool!"

	x := s1[0]
	fmt.Println(x)

	// ERROR -> cannot assign to s1[0]
	// s1[0] = 'x'

	// printing the number of runes in the string
	// fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))

}
