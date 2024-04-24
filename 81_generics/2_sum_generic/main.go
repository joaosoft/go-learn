package main

import "fmt"

type Number interface {
	int64 | float64 | string
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// Initialize a map for the string values
	strings := map[string]string{
		"first":  "a",
		"second": "b",
	}

	// non generic way
	fmt.Printf("Non-Generic Sums: %v and %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
		SumStrings(strings))

	// generics, defining the type
	fmt.Printf("Generic Sums: %v and %v and %+v\n",
		SumIntsOrFloatsOrStrings[string, int64](ints),
		SumIntsOrFloatsOrStrings[string, float64](floats),
		SumIntsOrFloatsOrStrings[string, string](strings))

	// generics, without defining the type
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v and %v\n",
		SumIntsOrFloatsOrStrings(ints),
		SumIntsOrFloatsOrStrings(floats),
		SumIntsOrFloatsOrStrings(strings))

	// generics with a defined multi type
	fmt.Printf("Generic Sums with Constraint: %v and %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
		SumNumbers(strings))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumStrings adds together the values of m.
func SumStrings(m map[string]string) string {
	var s string
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloatsOrStrings sums the values of map m. It supports floats and integers and strings
// as map values.
func SumIntsOrFloatsOrStrings[K comparable, V int64 | float64 | string](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. Its supports integers
// and floats and strings as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
