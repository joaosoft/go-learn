/*
Consider the code from the previous exercise and:

1. Change the name or the struct value called me to "Andrei"

2. Take in a new variable the favorites colors of struct value called you.
Print out the type and the value of the new variable.

3. Add a new favorite color to the second struct value called you.

4. Print out the struct values.
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

	// 1.
	me.name = "Andrei"

	// 2.
	var colors []string = you.colors
	fmt.Printf("Type: %T, Value: %v\n", colors, colors)

	// 3.
	colors = append(colors, "green")
	you.colors = colors

	// 4.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

}
