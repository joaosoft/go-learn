/////////////////////////////////
// Variadic Functions
// Go Playground: https://play.golang.org/p/ANNpW2SgpKw
/////////////////////////////////

// Variadic functions are functions that take a variable number of arguments.
// Ellipsis prefix (three-dots) in front of the parameter type makes a function variadic.
// The function may be called with zero or more arguments for that parameter.
// If the function takes parameters of different types, only the last parameter of a function can be variadic.

package main

import (
	"fmt"
	"strings"
)

// creating a variadic function
func f1(a ...int) {
	fmt.Printf("%T\n", a) // => []int, slice of int
	fmt.Printf("%#v\n", a)
}

// variadicfunction that modifies one of the arguments passed.
func f2(a ...int) {
	a[0] = 50
}

// creating a variadic function that calculates and returns the sum and product of its arguments
func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

// mixing variadic and non-variadic parameters is allowed
// non-variadic parameters are always before the variadic parameter
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
	return returnString
}

func main() {

	// calling variadic functions
	// a variadic function can be invoked with zero or more arguments
	f1(1, 2, 3, 4)

	f1() // a is: []int(nil)

	// an example of a well-known variadic function is append() built-in function.
	// it appends items to an existing slice and returns back the same slice.
	nums := []int{1, 2}
	nums = append(nums, 3, 4)

	s, p := sumAndProduct(2., 5., 10.)
	fmt.Println(s, p) // -> 17 100

	info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // => Age: 35, Full Name:Wolfgang Amadeus Mozart
}