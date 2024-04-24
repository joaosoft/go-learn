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

	// **GOTO STATEMENT **//

	//the following piece of code creates a loop like a for statement does
	i := 0
loop: // label
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//  goto todo //ERROR it's not permitted to jump over the declaration of x
	//  x := 5
	// todo:
	//  fmt.Println("something here")

outer2: //label, it doesn't conflict with other names
	for index, name := range people { // range returns both the index and the elements of the array one by one
		for _, friend := range friends { //iterating over the second array
			if name == friend {
				fmt.Printf("FOUND A FRIEND: %q at index %d\n", friend, index)
				continue outer2 //breaking outside the outer loop which terminates
			}
			fmt.Println("trying to find ...")
		}
	}

	fmt.Println("Next instruction after the break.")
}
