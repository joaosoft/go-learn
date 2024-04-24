/////////////////////////////////
// Operators in Go
// Go Playground: https://play.golang.org/p/eqoC_bAP6Tj
/////////////////////////////////

package main

import "fmt"

func main() {
	a, b := 2, 3.5

	//** ARITHMETIC OPERATORS **//
	//  +       sum
	// -        difference
	// *        product
	// /        quotient
	// %        remainder
	// there is no power operator in Go. Use math.Pow(a, b) for raising to a power.

	fmt.Println(a + 1)      // => 3
	fmt.Println(a - 1)      // => 1
	fmt.Println(a * a)      // => 4
	fmt.Println(a / a)      // => 1
	fmt.Println(int(b) % a) // => 1

	// Go is a Strong Typed Language
	// fmt.Println(a * b)       // =>  invalid operation: a * b (mismatched types int and float64)
	fmt.Println(a * int(b))     // => 6
	fmt.Println(float64(a) * b) // => 7

	// IncDec Statements
	// The "++" and "--" statements increment or decrement their operands by the untyped constant 1.
	x := 10
	x++ // x is 11. Same as: x += 1
	x-- // x is 10. Same as: x -= 1

	//** ASSIGNMENT OPERATORS **//
	//  =   (simple assignment)
	// +=   (increment assignment)
	// -=   (decrement assignment)
	// *=   (multiplication assignment)
	// /=   (division assignment)
	// %=   (modulus assignment)

	a = 10
	a += 2 // => a is 12
	a -= 3 // => a is 9
	a *= 2 // => a is 18
	a /= 3 // => a is 6
	a %= 5 // => a is 1

	//** COMPARISON OPERATORS **//
	//  ==      equal values
	// !=       not equal
	// >        left operand is greater than right operand
	// <        left operand is less than right operand
	// >=       left operand is greater than or equal to right operand
	// <=       left operand is less than or equal to right operand

	fmt.Println(5 == 6)   // => false
	fmt.Println(5 != 6)   // => true
	fmt.Println(10 > 10)  // => false
	fmt.Println(10 >= 10) // => true
	fmt.Println(5 < 5)    // => false
	fmt.Println(5 <= 5)   // => true

	//** LOGICAL OPERATORS **//
	// &&       logical and
	// ||       logical or
	// !        logical negation

	fmt.Println(0 < 2 && 4 > 1) // => true
	fmt.Println(1 > 5 || 4 > 5) // => false
	fmt.Println(!(1 > 2))       // => true

}
