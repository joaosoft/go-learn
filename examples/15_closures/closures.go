// testes

package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_closures()
	_methods()
	_interfaces()
	_toString()
	_errors()
	_leitores()
	_image()
	_goRoutines()
	_select()
	_tree()
	_mutex()
	_webServer()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func closures_adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func _closures() {
	fmt.Println("_closures()")
	/* Funções closures
	E as funções são closures completos.
	A função closures_adder retorna um closure. Cada closure abriga sua própria variável sum.
	*/
	pos, neg := closures_adder(), closures_adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
