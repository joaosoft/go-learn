/*
Write a function called myFunc() that takes exactly one argument
which is an int number written between double quotes (this is in fact a string).
If the argument is integer 'n', the function should return the result of n + nn + nn

Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func myFunc(n string) int {

	// convert string to int
	i, err1 := strconv.Atoi(n)

	// error handling
	if err1 != nil {
		fmt.Printf("Cannot convert %q to int.", n)
		os.Exit(1) //exit the program
	}
	ii, _ := strconv.Atoi(n + n)
	iii, _ := strconv.Atoi(n + n + n)

	return i + ii + iii

}

func main() {

	fmt.Println(myFunc("5")) // => 615
	fmt.Println(myFunc("3")) // => 369

}
