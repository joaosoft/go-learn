/*
Consider the money type declared at exercise #1.
Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.


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

*/

package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func main() {

	var salary money = 1.5 * 5.503
	fmt.Println(salary) // => 8.2545

	salary.print() // => 8.25

	s := salary.printStr()
	fmt.Println(s)
}
