package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_arrays()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _arrays() {
	fmt.Println("_arrays()")

	// exemplo 1
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// exemplo 2 (slices of matrix)
	var s []int = primes[1:4]
	fmt.Println(s)

	// exemplo 3 (changing matrix values)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)

	// exemplo 4
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)

	// exemplo 5 (capacidade e tamanho)
	s5 := []int{2, 3, 5, 7, 11, 13}
	printSlice("s5", s5)

	// Slice the slice to give it zero length.
	s5 = s5[:0]
	printSlice("s5", s5)

	// Extend its length.
	s5 = s5[:4]
	printSlice("s5", s5)

	// Drop its first two values.
	s5 = s5[2:]
	printSlice("s5", s5)

	// exemplo 6 (valor default de slice é 'nil')
	var s6 []int
	fmt.Println(s, len(s6), cap(s6))
	if s6 == nil {
		fmt.Println("nil!")
	}

	// exemplo 7 (criar slices com 'make' com 'len' e 'cap') make([]int, len, cap)
	sa := make([]int, 5)
	printSlice("sa", sa)

	sb := make([]int, 0, 5)
	printSlice("sb", sb)

	sc := sb[:2]
	printSlice("sc", sc)

	sd := sc[2:5]
	printSlice("sd", sd)

	// exemplo 8 (slices de slices)
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// exemplo 9 (append values in slices)
	var s9 []int
	printSlice("s9", s9)

	// append works on nil slices.
	s9 = append(s9, 0)
	printSlice("s9", s9)

	// The slice grows as needed.
	s9 = append(s9, 1)
	printSlice("s9", s9)

	// We can add more than one element at a time.
	s9 = append(s, 2, 3, 4)
	printSlice("s9", s9)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
