/////////////////////////////////
// Variables and Declarations
// Go Playground: https://play.golang.org/p/PKdAxUp8mNT
/////////////////////////////////

package main

import "fmt"

func main() {

	//** DECLARING VARIABLES **///

	// Syntax: var var_name type
	var s1 string
	s1 = "Learning Go!"
	fmt.Println(s1) // printing string s1

	//** TYPE INFERENCE **//

	// Go deduces automatically the type of the variable by looking at the initial value (bool, int, string etc)

	var k int = 6 // not necessary to say the type (int). It is inferred from the literal on the right side of =
	var i = 5     // type int
	var j = 5.6   // type float64

	// printing i, j and k
	fmt.Println("i:", i, "j:", j, "k:", k)

	// ii == jj  // -> error: cannot assign float to int (Go is a strong typed language)

	// declaring and initializing a new variable of type string (type inference)
	var s2 = "Go!"
	_ = s2 //in Go each variable must be used or there is a compile-time error
	// _ is the Blank Identifier and mutes the error of unused variables
	// _ can be only on the left hand side of the = operator

	// multiple assignments
	var ii, jj int
	ii, jj = 5, 8 // -> tuple assignment. It allows several variables to be assigned at once

	// swapping two variables using multiple assignments
	ii, jj = jj, ii

	fmt.Println(ii, jj)

	//** Short Declaration (works only in Block Scope!) **//

	// := (colon equals syntax) used only when declaring a new variable (or at least a new variable)
	// := tells go we are going to create a new variable and go figures out what type it will be
	s3 := "Learning golang!"
	_ = s3

	// can't use short declaration at Package Scope (outside main() or other function)
	// all statements at package scope must start with a Go keyword (package, var, import, func etc)

	// multiple short declaration
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	// redeclaration with short declaration syntax
	// at least one variable must be NEW on the left side of :=
	var deleted = false
	deleted, file := true, "a.txt"
	_, _ = deleted, file

	// expressions in short declarations are allowed
	sum := 5 + 2.5
	fmt.Println(sum)

	// multiple declaration is good for readability
	var (
		age       float64
		firstName string
		gender    bool
	)
	_, _, _ = age, firstName, gender

	// a concise way to declare multiple variables that have the same type
	var a, b int
	_, _ = a, b

}
