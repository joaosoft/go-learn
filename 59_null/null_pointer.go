package main

import (
	"encoding/json"
	"fmt"
)

type TestePointer struct {
	String            *string     `json:"string"`
	Int               *int   	 `json:"int"`
}

func main() {
	fmt.Println("\n\n\n\n\n######### TESTE 1 #########")

	xstr := "teste 1"
	xint := 123

	teste := TestePointer {
		String:      &xstr,
		Int:         &xint,
	}
	fmt.Println("TESTE: ", teste)


	bytes, err := json.Marshal(teste)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println("BYTES: ", bytes)


	var teste1 TestePointer
	if err := json.Unmarshal(bytes, &teste1); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println("TESTE 1: ", teste1)

	fmt.Printf("\n\nmy_string: %s\n", *teste1.String)
	fmt.Printf("my_int: %d\n", *teste1.Int)





	fmt.Println("\n\n\n\n\n######### TESTE 2 #########")


	teste2 := TestePointer {
	}
	fmt.Println("TESTE: ", teste2)


	bytes2, err := json.Marshal(teste2)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println("BYTES: ", bytes2)


	var teste1_2 TestePointer
	if err := json.Unmarshal(bytes2, &teste1_2); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println("TESTE 1: ", teste1_2)

	fmt.Printf("\n\nmy_string is nil: %s\n", teste1_2.String == nil)
	fmt.Printf("my_int is nil: %d\n", teste1_2.Int == nil)
}
