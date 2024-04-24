/*
Create a function that takes in an int value and prints out that value.

Assign the function to a variable, print out the type of the variable
and then call that function through the variable name.
*/

package main

import "fmt"

func f1(a int) {
	fmt.Println(a)
}

func main() {
	x := f1
	fmt.Printf("%T\n", x) // => func(int)
	x(8)
}
