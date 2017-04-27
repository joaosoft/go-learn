package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_tree()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
// TREE
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func _tree() {

}
