/////////////////////////////////
// The Typical Structure of a Go Application
// Go Playground: https://play.golang.org/p/vY_IeYBb1GN
/////////////////////////////////

// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main

// import declaration declares packages used in this file
import "fmt"

// package scoped variables and constants
var x int = 100

const y = 0

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main() {

	// Local Scoped Variables and Constants Declarations, calling functions etc
	var a int = 7
	var b float64 = 3.5
	const c int = 10

	// Println() function prints out a line to stdout
	// It belongs to package fmt
	fmt.Println("Hello Go world!") // => Hello Go world!
	fmt.Println(a, b, c)           // => 7 3.5 10

}
