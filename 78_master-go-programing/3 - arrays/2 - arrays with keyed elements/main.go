/////////////////////////////////
// Arrays with Keyed Elements
// Go Playground: https://play.golang.org/p/OXW0QyD5lDY
/////////////////////////////////

package main

import "fmt"

func main() {

	// each key corresponds to an index of the array
	grades := [3]int{ //the keyed elements can be in any order
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) // -> [5 10 7]

	accounts := [3]int{
		2: 50,
	}
	fmt.Println(accounts) //[0 0 50]

	names := [...]string{
		4: "Dan",
	}
	fmt.Println(len(names)) // -> 5

	// unkeyed element gets its index from the last keyed element
	cities := [...]string{
		5:        "Paris",
		"London", // this is at index 6
		1:        "NYC",
	}
	fmt.Printf("%#v\n", cities) // -> [7]string{"", "NYC", "", "", "", "Paris", "London"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // => [false false false false false true true]
}
