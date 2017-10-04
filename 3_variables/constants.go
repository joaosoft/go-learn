package main

import (
	"fmt"
	"math"
)

const s string = "constant"


func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_constants()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _constants() {
	fmt.Println("_constants()")

	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
