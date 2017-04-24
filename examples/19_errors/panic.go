// testes

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_panic()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _panic() {
	fmt.Println("_panic()")
	if os.Args[1] != "arg1" {
		panic("arg1")
	}
}
