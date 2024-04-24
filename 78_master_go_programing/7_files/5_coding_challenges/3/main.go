/*
Remove the file created at exercise #1.
Take care that the file is now called data.txt (it has been renamed at exercise #2).

Perform error handling.
*/

package main

import (
	"log"
	"os"
)

func main() {
	// removing the file
	err := os.Remove("data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
