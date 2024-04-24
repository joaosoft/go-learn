/*
A junior Go Programmer has just developed the following Go Program.   You want the change the book's price by calling changePrice method. However, you notice that after calling the method the price is not changed.

You task is to change the code in order for the changePrice method to work as expected.



package main

import "fmt"

type book struct {
	title string
	price float64
}

// method for book type
func (b book) changePrice(p float64) {
	b.price = p
}

func main() {
	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
}
*/

package main

import "fmt"

type book struct {
	title string
	price float64
}

// method for book type
// the method should receive a pointer to book, not a value
// in Go everyhing is passed by value so functions work on copies!
func (b *book) changePrice(p float64) {
	b.price = p
}

func main() {
	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // the price is changed
}
