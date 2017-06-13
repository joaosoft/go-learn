package main

import (
	"fmt"
)

func main() {
	fmt.Println("_try-catch()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	SoleTitle("teste do joao")

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// _soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
// SoleTitle ...
func SoleTitle(titleIn string) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("-> nil")
			// no panic
		case bailout{}:
			// "expected" panic
			fmt.Println("-> bailout")
			err = fmt.Errorf("multiple title elements")
		default:
			fmt.Println("-> default")
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
