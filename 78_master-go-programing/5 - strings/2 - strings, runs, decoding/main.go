/////////////////////////////////
// Strings, Runes, Bytes and Unicode Code Points
// Go Playground: https://play.golang.org/p/pttCqLAAvKA
/////////////////////////////////

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// characters or rune literals are expressed in Go by enclosing them in single quotes
	// declaring a variable of type rune (alias to int32)
	var1 := 'a'
	fmt.Printf("Type: %T, Value:%d \n", var1, var1) // => Type: int32, Value:97
	// value is 97 (the code point for a)

	// declaring a string value that contains non-ascii characters
	str := "ţară" // ţară means country in Romanian
	// 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.

	//The len() built-in function returns the no. of bytes not runes or chars.
	fmt.Println(len(str)) // -> 6,  4 runes in the string but the length is 6

	// returning the number of runes in the string
	m := utf8.RuneCountInString(str)
	fmt.Println(m) // => 4

	// by using indexes we get the byte at that position, not rune.
	fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163

	// decoding a string byte by byte
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // -> Å£arÄ
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune manually:
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:]) // it returns the rune in string in variable r
		//and its size in bytes in variable size

		// printing out each rune
		fmt.Printf("%c", r) // -> ţară
		i += size           // incrementing i by the size of the rune in bytes
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune automatically:
	for i, r := range str { //the first value returned by range is the index of the byte in string where rune starts
		fmt.Printf("%d -> %c", i, r) // => ţară
	}

}
