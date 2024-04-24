// There are some errors in the following Go program.
// Try to identify the errors, change the code and run the program without errors.
package main

func main() {
	var a float64 = 7.1

	x, y := true, 3.7

	// ERROR -> no new variables on left side of :=
	// a, x := 5.5, false

	a, x = 5.5, false

	_, _, _ = a, x, y

}
