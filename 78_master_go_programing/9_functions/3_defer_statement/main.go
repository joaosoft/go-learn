/////////////////////////////////
// Defer Function Calls
// Go Playground: https://play.golang.org/p/KGulBq828NT
/////////////////////////////////

package main

import (
	"fmt"
)

func foo() {
	fmt.Println("This is foo()!")
}

func bar() {
	fmt.Println("This is bar()!")
}

func foobar() {
	fmt.Println("This is foobar()!")
}

func main() {

	// a defer statement defers or postpones the execution of a function until the surrounding function returns.

	// by deferring foo() it will execute it just before exiting the surrounding function which is main()
	defer foo()
	bar()

	fmt.Println("Just a string after deferring foo() and calling bar()")

	// if there are more functions deferred, Go will execute them in the reverse order they were deferred
	defer foobar()

	/*
	   When executing the program the fallowing output is printed out:

	   This is bar()!
	   Just a string after deferring foo() and calling bar()
	   This is foobar()!
	   This is foo()!
	*/

}