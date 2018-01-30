package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/guregu/null.v3"
)

type Teste struct {
	String null.String `json:"string"`
	Int    null.Int    `json:"int"`
}

func main() {
	fmt.Println("\n\n\n\n\n######### TESTE 1 #########")

	xstr := null.StringFrom("teste 1")
	xint := null.IntFrom(123)

	teste := Teste{
		String: xstr,
		Int:    xint,
	}
	fmt.Println("TESTE: ", teste)

	bytes, err := json.Marshal(teste)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println("BYTES: ", bytes)

	var teste1 Teste
	if err := json.Unmarshal([]byte(`{"string":"", "int": 3}`), &teste1); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println("TESTE 1: ", teste1)

	if !teste1.String.Valid {
		fmt.Printf("is null!!!\n")
	}
	fmt.Printf("\n\nmy_string: %s\n", teste1.String.String)
	fmt.Printf("my_int: %d\n", teste1.Int.Int64)

	fmt.Println("\n\n\n\n\n######### TESTE 2 #########")
	var teste2 Teste
	if err := json.Unmarshal([]byte(`{"string":null, "int": 3}`), &teste2); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println("TESTE 2: ", teste2)

	if !teste2.String.Valid {
		fmt.Printf("is null!!!\n")
	}
	fmt.Printf("\n\nmy_string: %s\n", teste2.String.String)
	fmt.Printf("my_int: %d\n", teste2.Int.Int64)

	fmt.Println("\n\n\n\n\n######### TESTE 3 #########")
	var teste3 Teste
	if err := json.Unmarshal([]byte(`{"int": 3}`), &teste3); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println("TESTE 3: ", teste3)

	if !teste3.String.Valid {
		fmt.Printf("is null!!!\n")
	}
	fmt.Printf("\n\nmy_string: %s\n", teste3.String.String)
	fmt.Printf("my_int: %d\n", teste3.Int.Int64)
}
