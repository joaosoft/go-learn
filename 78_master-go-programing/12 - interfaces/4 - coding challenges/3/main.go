/*
There is an error in the following Go program.
Try to identify the error, change the code and run the program without errors.



package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{}
	x = cube{edge: 5}
	v := volume(x)
	fmt.Printf("Cube Volume: %v\n", v)
}

*/

package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {

	var x interface{}
	x = cube{edge: 5}

	// Type Assertion
	v := volume(x.(cube))

	fmt.Printf("Cube Volume: %v\n", v)

}
