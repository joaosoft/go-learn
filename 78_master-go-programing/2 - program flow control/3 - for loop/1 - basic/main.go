/////////////////////////////////
// For Loops
// Go Playground: https://play.golang.org/p/RiErMJCR3Z_c
/////////////////////////////////

package main

import "fmt"

func main() {

	// printing numbers from 0 to 9
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// has the same effect as a while loop in other languages
	// there is no while loop in Go
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}


	// infinite loop
	// sum := 0
	// for {
	//  sum++
	// }
	// fmt.Println(sum) //this line is never reached
}