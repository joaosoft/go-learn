/////////////////////////////////
// Anonymous and Embedded Structs
// Go Playground: https://play.golang.org/p/NtH6I30gtxb
/////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	// an anonymous struct is a struct with no explicitly defined struct type alias.
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	// =>struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:30
}
