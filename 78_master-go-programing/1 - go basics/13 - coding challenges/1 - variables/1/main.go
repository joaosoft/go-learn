///////////////////////////////
// Using var keyword, declare 3 variables called a, b and c of type int, float64, bool and string.
// Using short declaration syntax declare and assign these values to variables x, y and z:
// - 20
// - 15.5
// - "Gopher!"
// Using fmt.Println() print out the values of a, b, c, x, y and z.
// Try to run the program without error.
//////////////////////////////

package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string

	x := 20
	y := 15.5
	z := "Gopher!" // https://blog.golang.org/gopher

	fmt.Println(a, b, c, d)
	fmt.Println("x is", x, "y is", y, "z is", z)
}
