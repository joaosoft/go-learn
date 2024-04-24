/*
1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.

2. Declare and initialize two values of type person, one called me and another called you.

3. Print out the struct values.
*/

package main

import "fmt"

// 1.
type person struct {
	name   string
	age    int
	colors []string
}

func main() {
	// 2.
	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}

	// 3.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

}
