// There are some errors in the following Go program.
// Try to identify the errors, change the code and run the program without errors.
package main

import "fmt"

// ERROR -> := is not allowed in package scope (only local scope)
// version := "3.1"

var version = "3.1"

func main() {
	// ERROR -> A string is initialized only using double quotes ("")
	// name = 'Golang'

	name := "Golang"
	fmt.Println(name)
}
