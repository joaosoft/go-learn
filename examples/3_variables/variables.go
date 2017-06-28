package main

import (
	"fmt"
	"github.com/shopspring/decimal"

)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_variables()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _variables() {
	fmt.Println("_variables()")
	/*
		%d	decimal integer
		%x, %o, %b	integer in hexadecimal, octal, binary
		%f, %g, %e	floating-point number: 3.141593 3.141592653589793 3.141593e+00
		%t	boolean: true or false
		%c	rune (Unicode code point)
		%s	string
		%q	quoted string "abc" or rune 'c'
		%v	any value in a natural format
		%T	type of any value
		%%	literal percent sign (no operand)
	*/

	s1 := ""
	var s2 string
	var s3 = ""
	var s4 string = ""

	a := 1
	b := 2
	var is bool

	if is && s1 == "" && s2 == "" && s3 == "" && s4 == "" {

	}

	// if condition
	if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}

	regular_price, _ := decimal.NewFromString("")
	fmt.Println("PRECO: ", regular_price)
}
