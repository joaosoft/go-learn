/////////////////////////////////
// Command Line Arguments
// Go Playground: https://play.golang.org/p/C62c_H3UxN2
/////////////////////////////////

// You run the program at the console like this:
// go run main.go I learn Go Programming!

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args:", os.Args) // os.Args is slice of strings ([]string)

	// accessing command line arguments using indexes
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st Argument:", os.Args[1])
	fmt.Println("2nd Argument:", os.Args[2])
	fmt.Println("No. of items inside os.Args:", len(os.Args))
}