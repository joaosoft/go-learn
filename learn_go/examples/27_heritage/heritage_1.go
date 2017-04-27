package main

import (
	"fmt"
)

type Veiculo struct {
	Nome  string
	Rodas int
}

type Carro struct {
	Veiculo
	Portas int
}
type Moto struct {
	Veiculo
	Capacete bool
}

func main() {
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	moto := new(Moto)

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}
