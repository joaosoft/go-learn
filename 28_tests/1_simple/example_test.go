package example

import (
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestSomething(t *testing.T) {


	if Get != "joao" {
		t.Error("error validating person name")
	}

	if person.Age != 30 {
		t.Error("error validating person age")
	}
}
