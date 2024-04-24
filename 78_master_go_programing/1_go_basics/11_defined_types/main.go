/////////////////////////////////
// Named (Defined) Types in Go
// Go Playground: https://play.golang.org/p/v2-QZsESmC-
/////////////////////////////////

package main

import "fmt"

type age int        //new type, int is the underlying type
type oldAge age     //new type, int (not age) is the underlying type
type veryOldAge age //new type, int (not age) is the underlying type

func main() {

	// new type speed (underlying type uint)
	type speed uint

	// s1, s2 of type speed
	var s1 speed = 10
	var s2 speed = 20

	// performing operations with the new types
	fmt.Println(s2 - s1) // -> 10

	// uint and speed are different types (they have different names)
	var x uint

	// x = s1  //error different types

	// correct
	x = uint(s1)
	_ = x

	// correct
	var s3 speed = speed(x)
	_ = s3
}
