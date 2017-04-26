// testes

package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_flag()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

func _flag() {
	var period = flag.Duration("period", 1*time.Second, "sleep period")

	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
