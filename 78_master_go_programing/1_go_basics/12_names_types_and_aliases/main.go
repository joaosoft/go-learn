/////////////////////////////////
// Alias Declarations
// Go Playground: https://play.golang.org/p/bYzfoWGFWdd
/////////////////////////////////
package main

import "fmt"

func main() {

	// declaring a variable of type uint8
	var a uint8 = 10
	var b byte // byte is an alias to uit8

	// even though they have different names, byte and uit8 are the same type because they are aliases
	b = a // no error
	_ = b

	// declaring a new alias named second for uint
	// type alias_name = type_name
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint

	//no need to convert operations (same type)
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hour: 60
}