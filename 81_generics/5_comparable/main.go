package main

import (
	"fmt"
)

type Pessoa struct {
	Name string
}

func Equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	// how to use it
	fmt.Println(Equal("a", "b"))
	fmt.Println(Equal(1, 1))
	fmt.Println(Equal(true, true))
	fmt.Println(Equal(Pessoa{"Tiago"}, Pessoa{"Tiago"}))
}
