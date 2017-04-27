// testes

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// functions
func _files2() {
	fmt.Println("_files2()")

	filename := "./extras/rabbits.txt"

	// read in file, one command loads file into content variable
	// if an error occurs returns it as the second return value
	// if no error, err will be nil
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// log is a nifty little utility which can also output
		// in this case, a fatal log will halt the program
		log.Fatalln("Error reading file", filename)
	}

	// content returned as []byte not string
	// so must cast to string first and then can display
	fmt.Println(string(content))

	// write back to new file
	// see documentation for which methods take what type
	outfile := "output.txt"
	err = ioutil.WriteFile(outfile, content, 0644)
	if err != nil {
		log.Fatalln("Error writing file: ", err)
	}
}
