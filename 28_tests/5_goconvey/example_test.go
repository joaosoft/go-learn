package example

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Person struct {
	Name string
	Age  int
}

func TestSomething(t *testing.T) {

	Convey("Validating person", t, func() {
		person := Person{
			Name: "joao",
			Age:  30,
		}

		Convey("validating name", func() {

			Convey("the name should be joao", func() {
				So(person.Name, ShouldEqual, "joao")
			})

			Convey("the age should be 30", func() {
				So(person.Age, ShouldEqual, 30)
			})
		})
	})
}
