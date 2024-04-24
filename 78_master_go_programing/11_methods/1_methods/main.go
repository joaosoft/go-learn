/////////////////////////////////
// Methods (OOP)
// Go Playground: https://play.golang.org/p/ybxsVE8WvUc
/////////////////////////////////

package main

import (
	"fmt"
	"time"
)

// declaring a new type
type names []string

// declaring a method (function receiver)
func (n names) print() {
	// n is called method's receiver
	// n is the actual copy of the names we're working with and is available in the function.
	// n is like this or self from OOP.
	// any variable of type names can call this function on itself like variable_name.print()

	// iterating and printing all names
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {

	// Go doesn't have classes, but you can define methods on defined types.
	// a type may have a method set associated with it which enhances the type with extra behaviour.

	const day = 24 * time.Hour
	fmt.Printf("%T\n", day) // it's type is time.Duration

	// calling a method on time.Duration type
	// Seconds() is a method aka a receiver function.
	seconds := day.Seconds()

	// Seconds() returns the duration as a floating point number of seconds.
	fmt.Printf("%T\n", seconds)               //its type is float64
	fmt.Println("Seconds in a day:", seconds) // Seconds in a day: 86400

	// declaring a value of type names
	friends := names{"Dan", "Marry"}
	// calling the print() method on friends variable
	friends.print()

	// another way to call a method on a type
	names.print(friends) // not very common
}