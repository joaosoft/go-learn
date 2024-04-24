/////////////////////////////////
// String is Go
// Go Playground: https://play.golang.org/p/-o07MQbIsDv
/////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	// Strings are defined between double quotes "..."
	// Strings in Go are UTF-8 encoded by default
	// A string is in fact a slice of bytes in Go

	// declaring a string
	s1 := "Hi there  Go!"

	// printing a string
	fmt.Printf("%s\n", s1) // => Hi there  Go!
	fmt.Printf("%q\n", s1) // => "Hi there  Go!"

	// using double-quotes inside double quotes
	fmt.Println("He say: \"Hello!\"")

	// double quotes inside backticks (backquote)
	fmt.Println(`He say: "Hello!"`)

	// a string literal inclosed in backticks is called a raw string and it is interpreted literally.
	// backslashes or \n  have no special meaning
	s2 := `Hi there Go!`
	fmt.Println(s2)

	// declaring a multiline string
	fmt.Println("Price: 100 \nBrand: Nike")

	//the same with:
	fmt.Println(`
Price: 100
Brand: Nike`)

	// using backslashes inside a string:
	fmt.Println(`C:\Users\Andrei`)
	fmt.Println("C:\\Users\\Andrei")

	// concatenating strings (+)
	// Go creates a new string because strings are immutable in Go (this is not efficient).
	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!") // -> I love Go Programming!

	// getting an element (byte) of a string:
	fmt.Println("Element at index zero:", s3[0]) // => 73 (ascii code for I)
	//  a string is in fact a slice of bytes in Go

	// strings are immutable and can't be changed
	// s3[5] = 'x' // => error: Cannon assign to s3[5].
}