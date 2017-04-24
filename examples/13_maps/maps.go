// testes

package main

import (
	"fmt"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_maps()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _maps() {
	fmt.Println("_maps()")

	// exemplo 1
	type Vertex struct {
		Latitude, Longitude float64
	}
	var myVertex map[string]Vertex

	myVertex = make(map[string]Vertex)
	myVertex["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(myVertex["Bell Labs"])

	// exemplo 2
	//Maps literais são como structs literais, mas as chaves são obrigatórias.
	var m1 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m1)

	// sem indicar o nome Vertex na instancia
	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)

	// operações
	m3 := make(map[string]int)

	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	v1, ok1 := m3["Answer"]
	fmt.Println("The value:", v1, "Present?", ok1)

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	v2, ok2 := m3["Answer"]
	fmt.Println("The value:", v2, "Present?", ok2)

	m3["Answer"] = 40
	fmt.Println("The value:", m3["Answer"])

	v3, ok3 := m3["Answer"]
	fmt.Println("The value:", v3, "Present?", ok3)

}
