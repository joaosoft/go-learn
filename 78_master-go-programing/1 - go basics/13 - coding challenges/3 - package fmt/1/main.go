/*
Consider the following variable declarations:

x, y, z := 10, 15.5, "Gophers"
score := []int{10, 20, 30}
Using fmt.Printf():

Print each variable using a specific verb for its type

Print the string value enclosed by double quotes ("Gophers")

Print each variable using the same verb

Print the type of y and score
*/

package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// 1.
	fmt.Printf("x is %d, y is %f, z is %s\n", x, y, z)
	fmt.Printf("score is %#v\n", score)

	// 2.
	fmt.Printf("z is %q\n", z)

	// 3.
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score)

	// 4.
	fmt.Printf("y type: %T, score type: %T\n", y, score)

}
