/////////////////////////////////
// Passing Values and Pointers to Functions
// Go Playground: https://play.golang.org/p/4dAWL-iWp4I
/////////////////////////////////

// Go is a pass by value language (no exception here)

package main

import "fmt"

// declaring a function that takes an int, a float, a string and a bool value.
// the function works on copy so the changes are not seen outside (pass by value)
func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.5
	name = "Mobile Phone"
	sold = false
}

// declaring a function that takes in a pointer to int, a pointer to float, a pointer to string and a pointer to bool.
// the function makes a copy of each pointer but they point to the same address as the originals
func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	//changing the values the pointers point to is seen outside the function
	*quantity = 3
	*price = 500.5
	*name = "Mobile Phone"
	*sold = false
}

// declaring struct type
type Product struct {
	price       float64
	productName string
}

// declaring a function that takes in a struct value and modifies it
func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
	// the changes are not seen to the outside world
}

// declaring a function that takes in a pointer to struct value and modifies the value
func changeProductByPointer(p *Product) {
	(*p).price = 300
	p.productName = "Bicycle"
	// the changes are seen to the outside world

}

// declaring a function that takes in a slice
func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
	// the changes are seen to the outside world
}

// declaring a function that takes in a map
func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["x"] = 30
	// the changes are seen to the outside world
}

func main() {

	// declaring some variables
	quantity, price, name, sold := 5, 300.2, "Laptop", true
	fmt.Println("BEFORE calling changeValues():", quantity, price, name, sold)
	// => BEFORE calling changeValues(): 5 300.2 Laptop true

	// invoking the function has no effect on the variables.
	// the function works on and modifies copies, not originals.
	changeValues(quantity, price, name, sold)
	fmt.Println("AFTER calling changeValues():", quantity, price, name, sold)
	// => AFTER calling changeValues(): 5 300.2 Laptop true

	// the function modifies the values.
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("AFTER calling changeValuesByPointer():", quantity, price, name, sold)
	// => AFTER calling changeValuesByPointer(): 3 500.5 Mobile Phone false

	// declaring a struct value
	present := Product{
		price:       100,
		productName: "Watch",
	}

	// invoking the function has no effect on the struct value.
	// the function works on and modifies a copy, not the original.
	changeProduct(present)
	fmt.Println(present) // => {100 Watch}

	// the function modifies the struct value.
	changeProductByPointer(&present)
	fmt.Println("AFTER calling changeProductByPointer:", present)
	// => AFTER calling changeProductByPointer: {300 Bicycle}

	// declaring a slice
	prices := []int{10, 20, 30}

	// When a function changes a slice or a map the actual data is changed.
	changeSlice(prices)
	fmt.Println("prices slice after calling changeSlice():", prices)
	// => prices slice after calling changeSlice(): [11 21 31]

	// declaring a map
	myMap := map[string]int{"a": 1, "b": 2}
	// When a function changes a slice or a map the actual data is changed.

	changeMap(myMap)
	fmt.Println("myMap after calling changeMap():", myMap)
	// => myMap after calling changeMap(): map[a:10 b:20 x:30

	// slices and maps are not meant to be used with pointers.
}