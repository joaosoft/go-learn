// testes

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

type Veiculo1 interface {
	Preco() int
}

func (a Carro) Preco() int {
	return 6
}

func (a Moto) Preco() int {
	return 9
}

// define a behavior for Car
func (carro Carro) numeroRodas() int {
	return carro.Rodas
}

// a behavior only available for the Carro
func (a Carro) dizCarro() {
	fmt.Println("Hi Carro!")
}

// a behavior only available for the Carro
func (a Moto) dizMoto() {
	fmt.Println("Hi Moto!")
}

func main1() {
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	moto := new(Moto)
	moto.Nome = "ZUNDAPP"
	moto.Rodas = 4
	moto.Capacete = true

	carro := new(Carro)
	carro.Nome = "FIAT"
	carro.Rodas = 4
	carro.Portas = 5

	f := Carro{Veiculo{"RENAULT", 4}, 5}
	fmt.Println("O carro tem", f.numeroRodas(), "rodas")
	carro.dizCarro()
	moto.dizMoto()
	fmt.Println("Carro:", carro.Preco())
	fmt.Println("Moto:", moto.Preco())

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}
