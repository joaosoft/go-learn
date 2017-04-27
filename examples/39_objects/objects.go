// testes

package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_objects()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
// define Dog object type
type Dog struct {
	Name  string
	Color string
}

// Method for object specify the object the refer to and then
// the method name and rest of normal function definition
func (d Dog) Call() {
	fmt.Printf("Here comes a %s dog, %s.\n", d.Color, d.Name)
}

func _objects() {
	fmt.Println("_channels()")
	// create instance of object and set properties
	Spot := Dog{Name: "Spot", Color: "brown"}

	// or, try it like this...
	var Rover Dog
	Rover.Name = "Rover"
	Rover.Color = "light blondish with some dark patches and very scruffy"

	// call object method
	Spot.Call()
	Rover.Call()
}
