/*
Using a for loop print out all the years from your birthday to current year.
Use a variant of for loop where the post statement is moved inside the for block of code.
*/
package main

import "fmt"

func main() {
	birthYear, currentYear := 1980, 2020

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

