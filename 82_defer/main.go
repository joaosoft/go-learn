// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	deferExample()
	fmt.Println("finished")
}

func deferExample() {
	fmt.Println("this is the defer function")

	fmt.Println("running the for")
	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Println("for at index", i)
		}(i)
	}
	defer func() {
		fmt.Println("running the defer 1")
	}()
	defer func() {
		fmt.Println("running the defer 2")
	}()
	defer func() {
		fmt.Println("running the defer 3")
	}()
}
