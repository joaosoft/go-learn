/////////////////////////////////
// Label and Goto Statements
// Go Playground: https://play.golang.org/p/YCHNqu1IqaN
/////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	//** LABEL STATEMENT **//

	// declaring a variable
	// there is no conflict name between variable and label
	outer := 19
	_ = outer

	// declaring two arrays
	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Marry"}

	// searching for a single friend in a list of people.

outer: //label, it doesn't conflict with other names
	// iterating over the array.
	for index, name := range people { // range returns both the index and the elements of the array one by one
		for _, friend := range friends { //iterating over the second array
			if name == friend {
				fmt.Printf("FOUND A FRIEND: %q at index %d\n", friend, index)
				break outer //breaking outside the outer loop which terminates
			}
			fmt.Println("trying to find ...")
		}
	}

	fmt.Println("Next instruction after the break.")
}
