// testes

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_files()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _files() {
	fmt.Println("_files()")

	// exemplo 1
	file := "teste.txt"
	fmt.Println("opening file: ", file)
	//f, err := os.Open(file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File does not exists. Creating new file!")
		fmt.Fprintf(os.Stderr, "dump: %v\n", err)

		str := []byte("hello\n")
		err := ioutil.WriteFile(file, str, 0644)

		if err != nil {
		}
	} else {
		// Print file lines
		fmt.Println("File exists. Printing file!")

		for _, line := range strings.Split(string(data), "\n") {
			fmt.Println(line)
		}
	}
	//f.Close()

	// exemplo 2
	// Open the file and scan it.
	f, _ := os.Open(".teste.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		// Split the line on commas.
		parts := strings.Split(line, ",")

		// Loop over the parts from the string.
		for i := range parts {
			fmt.Println(parts[i])
		}
		// Write a newline.
		fmt.Println()
	}
}
