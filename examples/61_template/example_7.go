package main

import (
	"bytes"
	"fmt"
	"html/template"
)

type Query1 struct {
	Match string
	Name string
	Value string
}

func main() {

	const tmpl = `
{
  "query": {
    "bool": {
      "should": [
        {{range .}}
        { "{{.Match}}": { "{{.Name}}": "{{.Value}}" }},
        {{end}}
      ]
    }
  }
}
`

	type Queries []Query1
	var queries Queries = Queries{
		Query1 {
			Match: "match",
			Name: "nome",
			Value: "valor",
		},
		Query1 {
			Match: "match_all",
			Name: "nome_1",
			Value: "valor_1",
		},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("tmpl").Parse(tmpl))

	var tpl bytes.Buffer
	t.Execute(&tpl, queries)

	fmt.Println("--->" + tpl.String())
}
