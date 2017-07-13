package example_6

import (
	"os"
	"text/template"
)

type Query struct {
	Name string
	Value string
}

func main() {

	const tmpl = `
"query": {
	"match_all": { "{{.Name}}" : "{{.Value}}" }
}
`

	query := Query {
		Name: "nome",
		Value: "valor",
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("tmpl").Parse(tmpl))

	t.Execute(os.Stdout, query)
}