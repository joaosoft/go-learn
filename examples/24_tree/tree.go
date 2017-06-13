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
	var tree_head, tree_second, tree_tail Tree
	tree_head = Tree{ Value: 1, Right: &tree_second}
	tree_second = Tree{ Value: 2, Left: &tree_head, Right: &tree_tail}
	tree_tail = Tree{ Value: 3, Left: &tree_second }

	fmt.Println("HEAD:", tree_head.Value, "\nSECOND:", tree_head.Right.Value, "\nTHIRD:", tree_head.Right.Right.Value)
}
