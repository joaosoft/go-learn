// testes

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_scanner()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _scanner() {
	fmt.Println("_scanner()")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert text (\"stop\" to exit): ")
	for input.Scan() {
		fmt.Println("Insert text (\"stop\" to exit): ")
		if input.Text() == "stop" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
