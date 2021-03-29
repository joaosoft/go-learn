/////////////////////////////////
// Appending to a Slice
// Go Playground: https://play.golang.org/p/WYrW0Y_FeEF
/////////////////////////////////

package main

import "fmt"

func main() {
	numbers := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers = append(numbers, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n := []int{100, 200, 300}
	numbers = append(numbers, n...) // ... is the ellipsis operator
	fmt.Println(numbers)            // -> [2 3 10 20 30 40 100 200 300]
}