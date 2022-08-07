// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	recoverExample()

	fmt.Println("finished")
}

func recoverExample() {
	fmt.Println("this is the defer recover function")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()

	myPanic()

	fmt.Println("this is an unreachable code")
}

func myPanic() {
	panic("this is my panic error")
}
