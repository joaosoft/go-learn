package main

import (
	"html/template"
	"log"
	"os"
)

var tmplString = `    // content of index.html
    {{define "index"}}
    {{.var1}} is equal to {{.var2}}
    {{end}}
`

func main() {
	tmpl, err := template.New("test").Parse(tmplString)
	if err != nil {
		log.Fatal(err)
	}
	varmap := map[string]interface{}{
		"var1": "value",
		"var2": 100,
	}
	tmpl.ExecuteTemplate(os.Stdout, "index", varmap)

}
