/*
Consider the following constant declaration: const x float64 = 1.422349587101

Write a Go program that prints the value of x with 4 decimal points.
*/
package main

import "fmt"

func main() {
	const x float64 = 1.422349587101
	fmt.Printf("x is %.4f\n", x) // => formatting with 4 decimal points

}
