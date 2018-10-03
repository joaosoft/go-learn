package example

import (
	. "github.com/onsi/ginkgo"
)

type Person struct {
	Name string
	Age  int
}

var _ = Describe("Book", func() {
	var (
		person Person
	)

	BeforeEach(func() {
		person = Person{
			Name: "joao",
			Age:  30,
		}
	})

	Describe("Validating person", func() {
		Context("validating name", func() {
			It("the name should be joao", func() {
				if person.Name != "joao" {
					Fail("error comparing person name")
				}
			})
		})

		Context("validating age", func() {
			It("the name should be 30", func() {
				if person.Age != 30 {
					Fail("error comparing person age")
				}
			})
		})
	})
})
