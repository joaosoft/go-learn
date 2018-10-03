package example

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
				Expect(person.Name).To(Equal("joao"))
			})
		})

		Context("validating age", func() {
			It("the name should be 30", func() {
				Expect(person.Age).To(Equal(30))
			})
		})
	})
})
