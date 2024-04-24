/*
Create a Go Program that reads the entire contents of a file called info.txt into a string.
You can use ioutil.ReadAll() or any other function you want.

The file exists in the current working directory.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	// defer closing the file
	defer file.Close()

	// ioutil.ReadAll() reads from the file until an error or EOF and returns the data it read
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
}
