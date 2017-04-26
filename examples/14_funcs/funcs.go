// testes

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_funcs()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func funcs_compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func _funcs() {
	fmt.Println("_funcs()")
	hypot := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(1, 1))

	fmt.Println(funcs_compute(hypot))
	fmt.Println(funcs_compute(math.Pow))

	//Anonymous Functions
	_anonymousFunctions()

	// Variadic Functions
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
}

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func _anonymousFunctions() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

// Variadic Functions
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
