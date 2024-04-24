/////////////////////////////////
// Arrays in Go
// Go Playground: https://play.golang.org/p/PIsTyu-TDEo
/////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	// declaring an array with four elements of type int
	var numbers [4]int

	// array zero value is zeroed value elements
	fmt.Printf("%v\n", numbers)  // -> [0 0 0 0]
	fmt.Printf("%#v\n", numbers) // -> [4]int{0, 0, 0, 0}

	// declaring an array and initialize it using an array literal
	var a1 = [4]float64{}                           //initialized with defaults (0)
	var a2 = [3]int{5, -3, 5}                       //initialized with the given values
	a3 := [4]string{"Dan", "Diana", "Paul", "John"} //short declaration syntax
	a4 := [4]string{"x", "y"}                       //initializing only the first 2 elements

	// the ellipsis operator (...) finds out automatically the length of the array
	a5 := [...]int{1, 4, 5}
	fmt.Println("The length of a5 is: ", len(a5)) // len is 3

	// declare an array on multiple lines for better readability
	a6 := [...]int{1,
		2,
		3,
	} //the ending comma is mandatory when initializing the array on multiple lines and the closing curly brace is on its own line

	_, _, _, _, _, _ = a1, a2, a3, a4, a5, a6

	// changing an array
	// we can't add or remove elements from the array (it's fixed length)
	numbers[0] = 7              //changing first element (index 0)
	fmt.Printf("%v\n", numbers) // -> [7 0 0 0]

	// compile-time error
	// numbers[5] = 8  // invalid array index 5 (out of bounds for 4-element array)

	// getting an element
	x := numbers[0]
	fmt.Println("x is ", x) // => x is  7

	// iterating over an array (2-ways)
	for i, v := range numbers { // range is a language keyword used for iteration
		fmt.Println("index:", i, "value: ", v)

	}

	// iterating over an array (C/C++, Java Style)
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value: ", numbers[i])
	}

	// declaring a multi-dimensional arrays (array of arrays or matrix)
	balances := [2][3]int{
		[3]int{5, 6, 7}, //[3]int is optional
		{8, 9, 10},
	}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println("")
	}

	//  = operator creates a copy of an array.
	// the arrays are not connected and are saved in different memory locations
	m := [3]int{1, 2, 3}
	n := m //n is a copy of m

	fmt.Println("n is equal to m: ", n == m) // => true
	m[0] = -1                                //only m is modified
	fmt.Println("n is equal to m: ", n == m) // => false

}