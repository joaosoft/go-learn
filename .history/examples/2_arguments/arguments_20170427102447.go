package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_arguments()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _arguments() {
	fmt.Println("_arguments()")
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Arguments: ", s)
}
