package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_switch()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _switch() {
	fmt.Println("_switch()")

	// exemplo 1
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	// exemplo 2
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// exemplo 3
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
		break
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
		fallthrough
	default:
		fmt.Println("Good evening.")
	}
}
