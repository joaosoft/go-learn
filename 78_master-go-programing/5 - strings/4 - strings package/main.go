/////////////////////////////////
// Strings Package
// Go Playground: https://play.golang.org/p/uBy1U0Zve87
/////////////////////////////////

package main

import (
	"fmt"
	"strings"
)

func main() {

	// declaring a variable of type func to call the Println function easier.
	p := fmt.Println

	// it returns true whether a substr is within a string
	result := strings.Contains("I love Go Programming!", "love")
	p(result) // -> True

	// it returns true whether any Unicode code points are within our string, and false otherwise.
	result = strings.ContainsAny("success", "xy")
	p(result) // => false

	// it reports whether a rune is within a string.
	result = strings.ContainsRune("golang", 'g')
	p(result) // => true

	// it returns the number of instances of a substring in a string
	n := strings.Count("cheese", "e")
	p(n) // => 3

	// if the substr is an empty string Count() returns 1 + the number of runes in the string
	n = strings.Count("five", "")
	p(n) // => 5 (1 + 4 runes)

	// ToUpper() and ToLower() return a new string with all the letters
	// of the original string converted to uppercase or lowercase.
	p(strings.ToLower("Go Python Java")) // -> go python java
	p(strings.ToUpper("Go Python Java")) // -> GO PYTHON JAVA

	// comparing strings (case matters)
	p("go" == "go") // -> true
	p("Go" == "go") // -> false

	// comparing strings (case doesn't matter) -> it is not efficient
	p(strings.ToLower("Go") == strings.ToLower("go")) // -> true

	// EqualFold() compares strings (case doesn't matter) -> it's efficient
	p(strings.EqualFold("Go", "gO")) // -> true

	// it returns a new string consisting of n copies of the string that is passed as the first argument
	myStr := strings.Repeat("ab", 10)
	p(myStr) // => abababababababababab

	// it returns a copy of a string by replacing a substring (old) by another substring (new)
	myStr = strings.Replace("192.168.0.1", ".", ":", 2) // it replaces the first 2 occurrences
	p(myStr)                                            // -> 192:168:0.1

	// if the last argument is -1 it replaces all occurrences of old by new
	myStr = strings.Replace("192.168.0.1", ".", ":", -1)
	p(myStr) // -> 192:168:0:1

	// ReplaceAll() returns a copy of the string s with all non-overlapping instances of old replaced by new.
	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr) // -> 192:168:0:1

	// it slices a string into all substrings separated by separator and returns a slice of the substrings between those separators.
	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)                  // -> []string
	fmt.Printf("strings.Split():%#v\n", s) // => strings.Split():[]string{"a", "b", "c"}

	// If separator is empty Split function splits after each UTF-8 rune literal.
	s = strings.Split("Go for Go!", "")
	fmt.Printf("strings.Split():%#v\n", s) // -> []string{"G", "o", " ", "f", "o", "r", " ", "G", "o", "!"}

	// Join() concatenates the elements of a slice of strings to create a single string.
	// The separator string is placed between elements in the resulting string.
	s = []string{"I", "learn", "Golang"}
	j := strings.Join(s, "-")
	fmt.Printf("%T\n", j) // -> string
	p(j)                  // -> I-learn-Golang

	// splitting a string by whitespaces and newlines.
	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr) // it returns a slice of strings
	fmt.Printf("%T\n", fields)      // -> []string
	fmt.Printf("%#v\n", fields)     // -> []string{"Orange", "Green", "Blue", "Yellow"}

	// TrimSpace() removes leading and trailing whitespaces and tabs.
	s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux!\n ")
	fmt.Printf("%q\n", s1) // "Goodbye Windows, Welcome Linux!"

	// To remove other leading and trailing characters, use Trim()
	s2 := strings.Trim("...Hello, Gophers!!!?", ".!?")
	fmt.Printf("%q\n", s2) // "Hello, Gophers"

}
