/*
1. Declare a variable called empty of type empty interface. Print out its type.

2. Assign an int value to the variable called empty. Print out its type.

3. Assign a float64 value to empty. Print out its type.

4. Assign an int slice value to empty. Print out its type.

5. Add a new int value to the slice (empty variable).

6. Print out the slice (empty variable).
*/

package main

import "fmt"

func main() {
	// 1.
	var empty interface{}
	fmt.Printf("%T\n", empty)

	// 2.
	empty = 5
	fmt.Printf("%T\n", empty)

	// 3.
	empty = 10.10
	fmt.Printf("%T\n", empty)

	// 4.
	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)

	// 5. Type Assertion
	empty = append(empty.([]int), 10)

	// 6.
	fmt.Printf("%v\n", empty)

}
