/*
Using Iota declare the following months of the year: Jun, Jul and Aug

Jun, Jul and Aug are constant and their value is 6, 7 and 8.
*/

package main

import "fmt"

func main() {

	const (
		//iota starts from zero
		Jun = iota + 6
		Jul //automatically incremented by one
		Aug
	)

	fmt.Println(Jun, Jul, Aug)
}
