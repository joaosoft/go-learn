/////////////////////////////////
// Reading From Standard Input (console)
// Go Playground: https://play.golang.org/p/n8JuneN40_p
/////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// creating a scanner
	scanner := bufio.NewScanner(os.Stdin) //os.Stdin reads the output from the command line
	fmt.Printf("%T\n", scanner)           //pointer to bufio.scanner

	scanner.Scan() //it waits for the input and buffers the input untill a new line

	// gettting the scanned data
	text := scanner.Text()   // string type
	bytes := scanner.Bytes() // uint8[] slice type

	fmt.Println("Input text:", text)
	fmt.Println("Input bytes:", bytes)

	// reading the input continously until a specific string is scanned
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered:", text)
		if text == "exit" {
			fmt.Println("Exiting the scanning ...")
			break
		}
	}

	// error handling
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}