/*
Consider the following variable declaration x := 10.10

1. Print out the address of x

2. Declare a pointer called ptr that stores the address of x.

3. Print out the type and the value of ptr.

4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
*/

package main

import "fmt"

func main() {
	x := 10.10
	// 1.
	fmt.Printf("The address of x is %p\n", &x)

	// 2.
	ptr := &x

	// 3.
	fmt.Printf("ptr type: %T, ptr value: %v\n", ptr, ptr)

	// 4.
	fmt.Printf("The address of ptr: %p\n", &ptr)
	fmt.Printf("The value of x through the pointer:%f\n", *ptr)

}
