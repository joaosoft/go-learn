package main

import (
	_ "encoding/json" // forced import
	"fmt"             // normal import, you need to use it!
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_package()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions

func _package() {

}
