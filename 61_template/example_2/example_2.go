package example_2

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name                string
	nonExportedAgeField string //because it doesn't start with a capital letter
}

func main() {
	p := Person{Name: "Mary", nonExportedAgeField: "31"}

	t := template.New("nonexported template demo")
	t, _ = t.Parse("hello {{.name}}! Age is {{.nonExportedAgeField}}.")
	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
}
