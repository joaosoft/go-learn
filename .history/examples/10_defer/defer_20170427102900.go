package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_defer()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _defer() {
	fmt.Println("_defer()")
	// exemplo 1
	defer fmt.Println("world")
	fmt.Println("hello")

	// exemplo 2
	/*
		Chamadas de funções adiadas são empurradas por uma pilha.
		Quando a função retorna, as chamadas de adiamento são executadas
		na ordem em que a última a entrar é a primeira a sair.
	*/
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
