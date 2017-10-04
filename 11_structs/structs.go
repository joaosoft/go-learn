package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_structs()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _structs() {
	fmt.Println("_structs()")
	/*
		Structs(
		A struct é uma coleção de campos.
		(E uma declaração type faz o que você espera.)
	*/

	// exemplo 1
	type Dummy struct {
		Z int
	}
	type Vertex struct {
		X int
		Y int
		Dummy
	}
	fmt.Println(Vertex{1, 2, Dummy{3}})

	// exemplo 2
	v := Vertex{1, 2, Dummy{3}}
	v.X = 4
	v.Z = 5
	fmt.Println(v.X, v.Z)
}
