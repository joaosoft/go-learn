package example_1

import (
	"os"
	"text/template"
)

type Person struct {
	Name string //exported field since it begins with a capital letter
}

func main() {
	t := template.New("hello template") //create a new template with some name
	t, _ = t.Parse("hello {{.Name}}!")  //parse some content and generate a template, which is an internal representation

	p := Person{Name: "Mary"} //define an instance with required field

	t.Execute(os.Stdout, p) //merge template ‘t’ with content of ‘p’
}
