package main

import (
	"fmt"
	"math"
)

const s string = "constant"

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 62
	Small = Big >> 61
)

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

	// SECOND EXAMPLE
	const Greeting = ""
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)

}
