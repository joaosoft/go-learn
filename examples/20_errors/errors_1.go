package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_errors()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
// ERROS - ver que tem o metodo 'Error' que imprime corretamente o erro!
// MyError ...
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func _errors() {
	fmt.Println("_errors()")

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
