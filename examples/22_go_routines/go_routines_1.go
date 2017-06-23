package main

import (
	"fmt"
	"time"
)

func main() {
	//f()    // call f(); wait for it to return
	//go f() // create a new goroutine that calls f(); don't wait

	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_goRoutines()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func _goRoutines() {
	fmt.Println("_goRoutines()")

	// GO ROUTINES
	go say("world")
	say("hello")
}
