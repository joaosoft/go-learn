///////////////////////////////
// 1. Declare a, b, c, d on a single line for better readability -> multiple declarations
// 2. Declare x, y and z on a single line -> multiple short declarations
// 3. Remove the statement that prints out the variables. See the error!
// 4. Change the program to run without error using the blank identifier (_)
//////////////////////////////

package main

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher!"

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z //using blank identifier to mute the unused variable error
	// _ stays always on the left side of =

}
