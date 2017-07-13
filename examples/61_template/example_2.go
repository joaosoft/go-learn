package main

import (
	"os"
	"text/template"
	"fmt"
)

type Person1 struct {
	Name string
	nonExportedAgeField string //because it doesn't start with a capital letter
}

func main() {
	p:= Person1{Name: "Mary", nonExportedAgeField: "31"}

	t := template.New("nonexported template demo")
	t, _ = t.Parse("hello {{.Name}}! Age is {{.nonExportedAgeField}}.")
	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
}