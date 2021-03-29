/*
1. Declare a rune called r that stores the non-ascii letter ă

2. Print out the type of r

3. Declare 2 strings that contains the values ma and m

4. Concatenate the strings and the rune in a new string (the new string will hold the value mamă ).

BTW: mamă means mother in Romanian.

Note: You should convert rune to string to contatenate it to another string.
*/

package main

import "fmt"

func main() {
	r := 'ă'                     // declaring a rune
	fmt.Printf("r type:%T\n", r) // rune is alias to int32

	s1, s2 := "ma", "m" // declaring 2 strings

	// concatenating strings
	s := s1 + s2 + string(r)   // converting rune to string (expliction conversion is required)
	fmt.Printf("s is %s\n", s) // => s is mamă

}
