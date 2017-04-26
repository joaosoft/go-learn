// testes

package main

import (
	"fmt"
	"os"
)

// Main1 ...
func Main1() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_panic()

	SoleTitle("teste")

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _panic() {
	fmt.Println("_panic()")
	if os.Args[1] != "arg1" {
		panic("arg1")
	}
}

// _soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
// SoleTitle ...
func SoleTitle(titleIn string) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()

	// Bail out of recursion if we find more than one non-empty title.
	if titleIn != "" {
		panic(bailout{}) // multiple title elements
	}

	if titleIn == "" {
		return "", fmt.Errorf("no title element")
	}
	return titleIn, nil
}
