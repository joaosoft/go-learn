/*
Write the correct logical operator (&&, || , !) inside the expression so that result1 will be false and result2 will be true.

package main

import "fmt"

func main() {
	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := x < z 		 int(x) != int(z)
	fmt.Println(result1)

	result2 := y == 1 * 5 		 int(z) == 0
	fmt.Println(result2)

}
*/
package main

import "fmt"

func main() {
	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	// false
	result1 := x < z || int(x) != int(z)
	fmt.Println(result1)

	// true
	result2 := y == 1*5 && int(z) == 0
	fmt.Println(result2)

}
