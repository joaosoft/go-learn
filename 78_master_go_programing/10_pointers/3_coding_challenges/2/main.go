/*
Consider the following variable declarations:

x, y := 10, 2
ptrx, ptry := &x, &y
Declare a new variable called z and initialize the variable by dividing x by y through the pointers.
*/

package main

import "fmt"

func main() {
	x, y := 10, 2
	ptrx, ptry := &x, &y

	z := *ptrx / *ptry
	fmt.Printf("z is %d\n", z)

}
