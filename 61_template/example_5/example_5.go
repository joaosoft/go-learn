package example_5

import (
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const tmpl = `
({{range .}}
	ID:{{.IdMigration}};
	NAME:{{.Name}}
{{end}})
`

	// Prepare some data to insert into the template.
	type User struct {
		Id   int
		Name string
	}

	type UserList []User
	var myuserlist UserList = UserList{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("tmpl").Parse(tmpl))

	t.Execute(os.Stdout, myuserlist)
}
