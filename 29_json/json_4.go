package main

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v3"
)

// People ...
type People struct {
	Nome   string      `json:"nome"`
	Idade  int         `json:"idade"`
	TESTE1 null.Int    `json:"teste1"`
	TESTE2 null.String `json:"teste2"`
}

func main() {
	var person People
	bytes := []byte(`{
		"nome": "joao",
		"idade": 29
	}`)

	if err := json.Unmarshal(bytes, &person); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("RESULT:", person)
	}

	// INDENT
	var jsonTestIndented bytes2.Buffer
	json.Indent(&jsonTestIndented, bytes, "", "")
	fmt.Println("INDENTED:", string(jsonTestIndented.Bytes()))

	var person2 = People{
		Nome:   "joao",
		Idade:  12,
		TESTE1: null.IntFrom(12),
		TESTE2: null.StringFrom("teste"),
	}

	bytes2, _ := json.Marshal(person2)
	fmt.Println("-->\n", string(bytes2))

}
