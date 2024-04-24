/*
Create a Go Program that:

1. Using single function creates a file called info.txt in the current directory.
If the file already exists it will truncate it to zero size.

2. Write the string "The Go gopher is an iconic mascot!" to the file.
*/

/////////////////////////////////
// Run the program: go run main.go
/////////////////////////////////

package main

import (
	"io/ioutil"
	"log"
)

func main() {

	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.

	bs := []byte("The Go gopher is an iconic mascot!")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

