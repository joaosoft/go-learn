/////////////////////////////////
// Anonymous and Embedded Structs
// Go Playground: https://play.golang.org/p/NtH6I30gtxb
/////////////////////////////////

package main

import (
	"fmt"
	"strings"
)

func main() {
	//** ANONYMOUS FIELDS **//

	// fields type becomes fields name.
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1) // => main.Book{string:"1984 by George Orwell", float64:10.2, bool:false}

	fmt.Println(b1.string) // => 1984 by George Orwell

	// mixing anonymous with named fields:
	type Employee1 struct {
		name   string
		salary int
		bool
	}

	e := Employee1{"John", 40000, false}
	fmt.Printf("%#v\n", e) // => main.Employee1{name:"John", salary:40000, bool:false}
	e.bool = true          // changing a field

	fmt.Println(strings.Repeat("#", 10))
}