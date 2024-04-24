/*
1. Create a new struct type called book with 2 fields: price and title of type float64 and string.

2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.

3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.
*/

package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() { //use a pointer receiver to change the struct value
	(*b).price = b.price * 0.9
	// equivalent to:
	// b.price = b.price * 0.9

}

func main() {

	// book value
	bestBook := book{title: "The Trial", price: 9.9}

	// calling the methods
	vat := bestBook.vat()
	fmt.Printf("Vat:%v\n", vat)

	bestBook.discount()
	fmt.Printf("%#v\n", bestBook)

}
