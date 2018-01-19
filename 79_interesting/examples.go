package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type A struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type B struct {
		Name    string `json:"name2"`
		Address string `json:"address"`
	}

	type Mix struct {
		A
		B
		Name string `json:"name"`
	}

	a := A{
		Name: "AAA",
		Age:  111,
	}

	b := B{
		Name:    "BBB",
		Address: "ADDRESS",
	}

	mix := Mix{
		A:    a,
		B:    b,
		Name: "ola",
	}

	fmt.Println(mix)

	json, _ := json.Marshal(mix)
	fmt.Println("PRINTING: " + string(json))

}
