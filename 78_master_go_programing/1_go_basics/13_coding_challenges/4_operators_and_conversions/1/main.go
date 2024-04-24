/*
Consider the following declarations:

var i int = 3
var f float64 = 3.2
Write a Go program that converts i to float64 and f to int.

Print out the type and the value of the variables created after conversion.
*/

package main

import "fmt"

func main() {
	var i = 3
	var f = 3.2

	// int to float64
	f1 := float64(i)
	fmt.Printf("f1's type: %T, f1's value: %f\n", f1, f1)

	// float64 to int
	i1 := int(f)
	fmt.Printf("i1's type: %T, i1's value: %d\n", i1, i1)

}
