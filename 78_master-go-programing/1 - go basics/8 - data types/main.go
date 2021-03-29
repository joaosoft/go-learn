/////////////////////////////////
// Go Data Types
// Go Playground: https://play.golang.org/p/P0cJCe-cR51
/////////////////////////////////

package main

import "fmt"

func main() {

	//** NUMERIC TYPES **//

	// uint8      the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// uint     either 32 or 64 bits
	// int      same size as uint

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers
	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	//int type
	var i1 int8 = -128     //min value
	fmt.Printf("%T\n", i1) // => int8

	var i2 uint16 = 65535  //max value
	fmt.Printf("%T\n", i2) // => int16

	var i3 int64 = -324_567_345  // underscores are used to write large numbers for a better readability
	fmt.Printf("%T\n", i3)       // => int64
	fmt.Printf("i3 is %d\n", i3) // => i3 is -324567345 (underscores are ignored)

	//float64 type
	var f1, f2, f3 float64 = 1.1, -.2, 5. // trailing and leading zeros can be ignored
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	//rune type
	var r rune = 'f'
	fmt.Printf("%T\n", r) // => int32 (rune is an alias to int32)
	fmt.Printf("%x\n", r) // => 66,  the hexadecimal ascii code for 'f'
	fmt.Printf("%c\n", r) // => f

	//bool type
	var b bool = true
	fmt.Printf("%T\n", b) // => bool

	//string type
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s) // => string

	//array type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers) // =>  [4]int

	//slice type
	var cities = []string{"London", "Bucharest", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities) // => []string

	//map type
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
	}
	fmt.Printf("%T\n", balances) // => map[string]float64

	//struct type
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you) // => main.Person

	//pointer type
	var x int = 2
	ptr := &x                                                 // pointer to int
	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr) // => ptr is of type *int with value 0xc000016168

	//function type
	fmt.Printf("%T\n", f) // => func()

}

func f() {
}
