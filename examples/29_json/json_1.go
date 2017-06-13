package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_json1()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

func _json1() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}

	// MARSHAL
	data1, err1 := json.Marshal(movies)
	if err1 != nil {
		log.Fatalf("JSON marshaling failed: %s", err1)
	}
	fmt.Printf("teste1: %s\n", data1)

	data2, err2 := json.MarshalIndent(movies, "", "    ")
	if err2 != nil {
		log.Fatalf("JSON marshaling failed: %s", err2)
	}
	fmt.Printf("teste2: %s\n", data2)

	// UNMARSHAL
	var titles []struct{ Title string }
	if err3 := json.Unmarshal(data1, &titles); err3 != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err3)
	}
	fmt.Println("teste3:", titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}
