// testes

package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_loops()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _loops() {
	fmt.Println("_loops()")
	// while
	for true {
		fmt.Println("for true")
		break
	}

	// infinite for
	for {
		fmt.Println("infinite for")
		break
	}

	//O range do laço for itera sobre uma slice ou map.
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	// creating an array of strings
	alphas := []string{"abc", "def", "ghi"}

	// standard for loop
	for i := 0; i < len(alphas); i += 1 {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	// if iterating over a array, this would be how you would actually do it
	// the standard loop would be used if uncommon step function
	// range returns two values, i -> index, val -> value
	for i, val := range alphas {
		fmt.Printf("%d: %s \n", i, val)
	}

	// if you did not care about the value and only wanted the index
	for i := range alphas {
		fmt.Println(i)
	}

	// if you did not care about the index and only the value, you use the _
	// which means don't save this value
	for _, val := range alphas {
		fmt.Println(val)
	}

	// how to use for like a while
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// infinite loop
	// for {
	//   fmt.Print(".")
	// }

	// Go is picky, "else" must be on the same line as closing if bracket
	num := 1
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Many")
	}

}
