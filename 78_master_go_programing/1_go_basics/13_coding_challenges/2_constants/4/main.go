// There are some errors in the following Go program.
// Try to identify the errors, change the code and run the program without errors.
package main

func main() {
	const x int = 10

	//** There are ONLY boolean constants, rune constants, integer constants, **//
	//** floating-point constants, complex constants, and string constants. **//

	// declaring a constant of type slice int ([]int)
	// ERROR -> const initializer []int literal is not a constant
	// const m = []int{1: 3, 4: 5, 6: 8}	-> You cannot declare a slice constant.
	// _ = m
}
