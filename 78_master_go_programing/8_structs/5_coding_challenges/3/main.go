/*
Consider the code from Exercise #1.

Iterate and print out the favorite colors of the struct value called me.



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
*/

package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
}

func main() {

	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}

	for index, color := range me.colors {
		fmt.Printf("%d -> %q\n", index, color)
	}

	_ = you

}
