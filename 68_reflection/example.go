package main

import (
	"fmt"
	"reflect"
	"strings"
)

type IEvent interface{}

// ErrorTest represents an event that the aggregate doesn't know how to handle.
type ErrorTest struct {
	eventType string
}

// NewErrorTest returns a new ErrorTest.
func NewErrorTest(event IEvent) *ErrorTest {
	e := new(ErrorTest)
	fmt.Println("TYPEOF:", reflect.TypeOf(event).String())
	s := strings.Split(reflect.TypeOf(event).String(), ".")
	e.eventType = s[len(s)-1]
	return e
}

// Error returns a string representation of the error.
func (err *ErrorTest) Error() string {
	return fmt.Sprintf("Event type '%s' not recognized by the system", err.eventType)
}

func main() {
	print("->\n")
	type OlaEstaEumaEstruturaFixe struct{}
	e := NewErrorTest(&OlaEstaEumaEstruturaFixe{})
	fmt.Println("ERROR: ", e.eventType, "\n", e.Error())

	fmt.Println("TYPE:", OlaEstaEumaEstruturaFixe{})
}
