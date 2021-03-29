/*
1. Create a new type called money. Its underlying type is float64.

2. Create a method called print that prints out the money value with exact 2 decimal points.
*/

package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func main() {

	var salary money = 1.5 * 5.503
	fmt.Println(salary) // => 8.2545

	salary.print() // => 8.25
}
