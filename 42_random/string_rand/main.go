package main

import (
	"fmt"
	"golang-learn/42_random/string_rand/api"
)

func main() {

	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_rand()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions

type MathFunc func(int, int) int

func _rand() {
	fmt.Println("_rand()")
	fmt.Println(fmt.Sprintf("GENERATED STRING 10A, %s", api.String(10)))
}
