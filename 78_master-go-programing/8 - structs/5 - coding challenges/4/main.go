/*
Change the code from Exercise #1 and:

1. Create a new struct type called grades with  2 fields: grade and course

2. Add another field of type grades to person struct type (embedded struct).

3. Change the initialization of the struct values called me and you to contain information about the grades.

4. Change the grade and the course of one struct value to "Golang" and 98.

5. Print out the struct values.



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

// 2
type person struct {
	name   string
	age    int
	colors []string
	gr     grades
}

// 1.
type grades struct {
	grade  int
	course string
}

func main() {
	// 3.
	me := person{
		name:   "Marius",
		age:    30,
		colors: []string{"red", "yellow"},
		gr:     grades{grade: 85, course: "Python"},
	}
	you := person{
		name:   "Maria",
		age:    22,
		colors: []string{"pink", "blue"},
		gr:     grades{grade: 100, course: "History"},
	}

	// 4.
	me.gr.grade = 98
	me.gr.course = "Golang"

	// 5.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

}
