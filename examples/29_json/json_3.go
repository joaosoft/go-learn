package main

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

// Teste ...
type Teste struct {
	A uuid.UUID `json:"a"`
	B uuid.UUID `json:"b"`
	C int64     `json:"c"`
	D uuid.UUID `json:"d"`
	E string    `json:"e"`
	F string    `json:"f"`
	G uuid.UUID `json:"g"`
	H string    `json:"h"`
	I time.Time `json:"i"`
}

func main() {

	// TESTE 1
	a, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	b, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	c := int64(1)
	d, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")
	e := "1000"
	f := "2000"
	g, _ := uuid.FromString("00000000-0000-0000-0000-000000000004")
	h := "disabled"
	var bytes []byte
	var err error

	bytes = []byte(`{
				"A": "` + a.String() + `",
                "b": "` + b.String() + `",
				"C": ` + strconv.Itoa(int(c)) + `,
				"D": "` + d.String() + `",
				"E": "` + e + `",
				"F": "` + f + `",
                "G": "` + g.String() + `",
                "H": "` + h + `"
            }`)
	fmt.Println("MY:", string(bytes))

	var teste = &Teste{
		A: a,
		B: b,
	}
	if err = json.Unmarshal(bytes, teste); err != nil {
		fmt.Println("ERROR:", err)
	}

	// TESTE 2
	var teste2 = &Teste{
		A: a,
		B: b,
		C: int64(c),
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
	}
	if bytes, err = json.Marshal(teste2); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("SUCCESS")
	}
	fmt.Println("STRUCTURE:", string(bytes))
	var teste3 = &Teste{}
	if err = json.Unmarshal(bytes, teste3); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("SUCCESS")
	}
}
