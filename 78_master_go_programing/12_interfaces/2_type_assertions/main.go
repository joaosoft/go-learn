/////////////////////////////////
// Interfaces - Type Assertions and Type Switches
// Go Playground: https://play.golang.org/p/0A07wpGtXYa
/////////////////////////////////

// A type assertion provides access to an interface’s concrete value.

package main

import (
	"fmt"
	"math"
)

// declaring an interface type called shape
type shape interface {
	area() float64
	perimeter() float64
}

// declaring a struct type
type rectangle struct {
	width, height float64
}

// declaring a struct type
type circle struct {
	radius float64
}

// declaring a method for circle type
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// declaring a method for circle type
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// declaring a method for circle type
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// declaring a method for rectangle type
func (r rectangle) area() float64 {
	return r.height * r.width
}

// declaring a method for rectangle type
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// declaring a function that takes an interface value
func print(s shape) {

	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {

	// declaring an interface value that holds a circle type value
	var s shape = circle{radius: 2.5}

	fmt.Printf("%T\n", s) //interface dynamic type is circle

	// no direct access to interface's dynamic values
	// s.volume() -> error

	// there is access only to the methods that are defined inside the interface
	fmt.Printf("Circle Area:%v\n", s.area())

	// an interface value hides its dynamic value.
	// use type assertion to extract and return the dynamic value of the interface value.
	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	// checking if the assertion succeded or not
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//

	// it permits several type assertions in series
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}
}