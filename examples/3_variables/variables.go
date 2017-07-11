package main

import (
	"fmt"
	"github.com/shopspring/decimal"

	"strconv"
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

	// decimal value 0
	regular_price, _ := decimal.NewFromString("")
	fmt.Println("PRECO: ", regular_price)

	// zero value 0
	zero, _ := decimal.NewFromString("0")
	fmt.Println("ZERO: ", zero)

	// -4 less then 0
	price, _ := decimal.NewFromString("-4")
	fmt.Println("??", price.Cmp(zero) == -1)

	fmt.Println(fmt.Sprintf("-----> %.2d", 1))
	fmt.Println(fmt.Sprintf("-----> %.2d", 10))
	fmt.Println(fmt.Sprintf("-----> %.2d", 100))

	val, _ := strconv.Atoi("01")
	fmt.Println(fmt.Sprintf("-----> %.5d.%.2d.%.2d", val, 2, 3))

	fmt.Println(fmt.Sprintf("-----> %5s.%2s.%2s", "1", "2", "3"))
}
