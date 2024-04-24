/////////////////////////////////
// Package fmt
// Go Playground: https://play.golang.org/p/JGb4akovl8W
/////////////////////////////////

package main

// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
// It's used mainly to print out to stdout
import "fmt"

func main() {

	// fmt.Println() writes to standard output.
	// spaces are always added between operands and a newline is appended.
	fmt.Println("Hello Go World!") // => Hello Go World!

	var name, age = "Andrei", 35
	fmt.Println(name, "is", age, "years old.") // => Andrei is 35 years old.

	//** fmt.Printf() **//

	// fmt.Printf() prints out to stdout according to a format specifier called verb.
	// It doesn't add a newline (\n)

	// VERBS:
	// %d -> decimal
	// %f -> float
	// %s -> string
	// %q -> double-quoted string
	// %v -> value (any)
	// %#v -> a Go-syntax representation of the value
	// %T -> value Type
	// %t -> bool (true or false)
	// %p -> pointer (address in base 16, with leading 0x)
	// %c -> char (rune) represented by the corresponding Unicode code point

	a, b, c := 10, 15.5, "Gophers"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s, grades are %+v \n", a, b, c, grades) // => a is 10, b is 15.500000, c is Gophers, grades are [10 20 30]
	fmt.Printf("%q\n", c)                                                       // => "Gophers"
	fmt.Printf("%v\n", grades)                                                  // => [10 20 30]
	fmt.Printf("%#v\n", grades)                                                 // => b is of type float64 and grades is of type []int
	fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)
	// => b is of type float64 and grades is of type []int
	fmt.Printf("The address of a: %p\n", &a) // => The address of a: 0xc000016128
	fmt.Printf("%c and %c\n", 100, 51011)    // =>  d and 읃  (runes for code points 101 and 51011)

	const pi float64 = 3.14159265359
	fmt.Printf("pi is %.4f\n", pi) // => formatting with 4 decimal points

	// %b -> base 2
	// %x -> base 16
	fmt.Printf("255 in base 2 is %b\n", 255)  //  => 255 in base 2 is 11111111
	fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65

	// fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
	s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)
	fmt.Println(s) // => a is 10, b is 15.500000, c is Gophers
	m := fmt.Sprintln("a:", a, "b:", b, "c:", c)
	fmt.Println(m) // => a is 10, b is 15.500000, c is Gophers
}
