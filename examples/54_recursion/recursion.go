// Go supports recursive functions. Here’s a classic factorial example.
package main
import "fmt"

//This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	val :=  n * fact(n-1)
	fmt.Println(n)
	return val
}
func main() {
	fmt.Println(fact(7))
}