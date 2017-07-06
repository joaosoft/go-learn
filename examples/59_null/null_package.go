package main

import (
	"encoding/json"
	"gopkg.in/guregu/null.v3"
	"fmt"
)

type Example struct {
	My_string_ok            null.String     `json:"My_string_ok"`
	My_string_ko            null.String     `json:"My_string_ko"`
	My_string_null          null.String     `json:"My_string_null"`
	My_int_ok               null.Int        `json:"My_int_ok"`
	My_int_ko               null.Int        `json:"My_int_ko"`
	My_int_null             null.Int        `json:"My_int_null"`
}

type Example3 struct {
	My_string_ok            null.String     `json:"My_string_ok"`
}

func main() {
	fmt.Println("\n\n\n\n\n######### TESTE 1 #########")

	teste := Example {
		My_string_ok:      null.StringFrom("teste"),
		My_string_ko:      null.NewString("", false),
		My_string_null:    null.String{},
		My_int_ok:         null.NewInt(123, true),
		My_int_ko:         null.NewInt(0, false),
		My_int_null:       null.Int{},
	}
	fmt.Println("TESTE: ", teste)


	bytes, err := json.Marshal(teste)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println("BYTES: ", bytes)


	var teste1 Example
	if err := json.Unmarshal(bytes, &teste1); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println("TESTE 1: ", teste1)


	fmt.Println("\n\nMy_string_ok:", teste.My_string_ok.Valid)
	fmt.Println("My_string_ko:", teste.My_string_ko.Valid)
	fmt.Println("My_string_null:", teste.My_string_null.Valid)
	fmt.Println("My_int_ok:", teste.My_int_ok.Valid)
	fmt.Println("My_int_ko:", teste.My_int_ko.Valid)
	fmt.Println("My_int_null:", teste.My_int_null.Valid)

	fmt.Println("\n\nMy_string_ok:", teste.My_string_ok)
	fmt.Println("My_string_ko:", teste.My_string_ko)
	fmt.Println("My_string_null:", teste.My_string_null)
	fmt.Println("My_int_ok:", teste.My_int_ok)
	fmt.Println("My_int_ko:", teste.My_int_ko)
	fmt.Println("My_int_null:", teste.My_int_null)



	fmt.Println("\n\n\n\nMy_string_ok:", teste1.My_string_ok.Valid)
	fmt.Println("My_string_ko:", teste1.My_string_ko.Valid)
	fmt.Println("My_string_null:", teste1.My_string_null.Valid)
	fmt.Println("My_int_ok:", teste1.My_int_ok.Valid)
	fmt.Println("My_int_ko:", teste1.My_int_ko.Valid)
	fmt.Println("My_int_null:", teste1.My_int_null.Valid)

	fmt.Println("\n\nMy_string_ok:", teste1.My_string_ok)
	fmt.Println("My_string_ko:", teste1.My_string_ko)
	fmt.Println("My_string_null:", teste1.My_string_null)
	fmt.Println("My_int_ok:", teste1.My_int_ok)
	fmt.Println("My_int_ko:", teste1.My_int_ko)
	fmt.Println("My_int_null:", teste1.My_int_null)




	fmt.Println("\n\n\n\n\n######### TESTE 2 #########")
	var My_string_ok = null.StringFrom("teste")
	My_string_ok_bytes, err := json.Marshal(My_string_ok)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println("BYTES: ", My_string_ok_bytes)


	var My_string_ok_2 null.String
	if err := json.Unmarshal(My_string_ok_bytes, &My_string_ok_2); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println("TESTE 2: ", My_string_ok_2)




	fmt.Println("\n\n\n\n\n######### TESTE 3 #########")

	teste3 := Example3 {
		My_string_ok:      null.StringFrom("teste"),
	}
	fmt.Println("TESTE: ", teste3)


	bytes3, err := json.Marshal(teste3)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println("BYTES: ", bytes3)


	var teste3_1 Example3
	if err := json.Unmarshal(bytes3, &teste3_1); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println("TESTE 3: ", teste3_1)


	fmt.Println("\n\nMy_string_ok:", teste3.My_string_ok.Valid)
	fmt.Println("My_string_ok:", teste3.My_string_ok)

	fmt.Println("\n\n\nMy_string_ok:", teste3_1.My_string_ok.Valid)
	fmt.Println("My_string_ok:", teste3_1.My_string_ok)
}
