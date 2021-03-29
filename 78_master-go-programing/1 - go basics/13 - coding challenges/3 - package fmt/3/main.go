// There are some errors in the following Go program.
// Try to identify the errors, change the code and run the program without errors.

/*
package main

import "fmt"

func main() {
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n")
	fmt.Printf("Circle's circumference with radius %d is %b\n", radius, circumference)
	_ = shape
}
*/
package main

import "fmt"

func main() {
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n")
	fmt.Printf("Circle's circumference with radius %d is %b\n", radius, circumference)
	_ = shape
}
