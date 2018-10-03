package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
	Age  int
}

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	person := Person{
		Name: "joao",
		Age:  30,
	}

	if assert.NotNil(person) {

		assert.Equal("joao", person.Name, "the name is invalid")

		assert.NotEqual(31, person.Age, "the age is invalid")

		assert.Equal(30, person.Age, "the age is invalid")
	}
}
