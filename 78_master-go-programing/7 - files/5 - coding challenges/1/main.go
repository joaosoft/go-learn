/*
Create a new file in the current working directory called info.txt and then close the file.
If the file cannot be created (no permissions, wrong path etc) then print out the error
and stop the program (do error handling).
*/

package main

import (
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("aa.txt")

	// error handling
	if err != nil {
		// log the error and exit the program
		log.Fatal(err) // the idiomatic way to handle errors

	}
	newFile.Close()
}
