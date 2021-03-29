/*
Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.

Print those elements and their sum.
*/

// run this program in a terminal with arguments
// ex: go run main.go 5 7.1 9.9 10

package main

import "fmt"

func main() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for _, v := range nums[2 : len(nums)-2] {
		fmt.Println(v)
		sum += v
	}
	fmt.Println("Sum:", sum)

}
