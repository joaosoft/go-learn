/////////////////////////////////
// Types and Zero Values
// Go Playground: https://play.golang.org/p/zItROROXi64
/////////////////////////////////

package main

import "fmt"

func main() {

	// you must provide a type for each variable you declare or Go should be able to infer it
	var a int = 10
	var b = 15.5      // type inference (deduction)
	c := "Gopher"     // short declaration, type inference
	_, _, _ = a, b, c // Blank Identifier (_) to get rid of unused variable error

	// Go is a Statically and Strong Typed Programming Language
	// a = 3.14 -> error. A variable cannot change it's type
	// a = b    -> error. It's not allowed to assign a type to another type

	//** ZERO VALUES **//

	// An uninitialized variable or empty variable  will get the so called ZERO VALUE
	// The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
	var value int                         // initialized with 0
	var price float64                     // initialized with 0.0
	var name string                       // initialized with empty string -> ""
	var done bool                         // initialized with false
	fmt.Println(value, price, name, done) // -> 0 0.0 ""  false
}