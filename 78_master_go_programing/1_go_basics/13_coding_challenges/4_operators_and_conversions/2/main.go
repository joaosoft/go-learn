/*
Consider the following declarations:

var i = 3
var f = 3.2
var s1, s2 = "3.14", "5"
Write a Go program that converts:

1. i to string (int to string)

2. s2 to int (string to int)

3. f to string (float64 to string)

4. s1 to float64  (string to float64)

5. Print the value and the type for each variable created after conversion.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// 1. int to string
	s := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// 2. string to int
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}

}