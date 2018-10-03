package example

import (
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestSomething(t *testing.T) {
	person := Person{
		Name: "joao",
		Age:  30,
	}

	if person.Name != "joao" {
		t.Error("error validating person name")
	}

	if person.Age != 30 {
		t.Error("error validating person age")
	}
}
