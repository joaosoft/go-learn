package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	// Variadic functions can be called in the usual way with individual arguments.
	sumx(1, 2)
	sumx(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sumx(nums...)
	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

func sumx(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
