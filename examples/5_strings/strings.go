// testes

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_strings()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _arguments() {
	fmt.Println("_arguments()")
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Arguments: ", s)
}

func _panic() {
	fmt.Println("_panic()")
	if os.Args[1] != "arg1" {
		panic("arg1")
	}
}
func _strings() {
	fmt.Println("_strings()")
	fmt.Println("string.Join: ", strings.Join(os.Args[1:], " "))

	// exemplo 1 (split string)
	data := "cat,bird,dog"
	// Split on comma.
	result := strings.Split(data, ",")
	// Display all elements.
	for i := range result {
		fmt.Println(result[i])
	}
	// Length is 3.
	fmt.Println(len(result))

	// exemplo 2
	// Create a slice and append three strings to it.
	values := []string{}
	values = append(values, "cat")
	values = append(values, "dog")
	values = append(values, "bird")
	// Join three strings into one.
	result1 := strings.Join(values, "...")
	fmt.Println(result1)

	// exemplo 3
	value := "Cat Dog    Bird  Fish"
	result2 := strings.Fields(value)

	// Display all fields, first field and count.
	fmt.Println(result2)
	fmt.Println(result2[0])
	fmt.Println(len(result2))

	// exemplo 4
	// Return true if comma or colon char.
	f := func(c rune) bool {
		return c == ',' || c == ':'
	}
	// Separate into fields with func.
	fields := strings.FieldsFunc("cat,dog:bird", f)
	fmt.Println(fields)

	// exemplo 5
	value5 := "one.two.three"
	// Split after the period.
	result5 := strings.SplitAfter(value5, ".")
	// Display our results.
	for i := range result5 {
		fmt.Println(result5[i])
		/*
			one.
			two.
			three
		*/
	}

	// exemplo 6 (regex split)
	value6 := "catXdogYbird"

	// First compile the delimiter expression.
	re := regexp.MustCompile("[XY]")

	// Split based on pattern.
	// ... Second argument means "no limit."
	result6 := re.Split(value6, -1)

	for i := range result6 {
		fmt.Println(result6[i])
		/*
			cat
			dog
			bird
		*/
	}

	// exemplo 7
	value7 := "bird"
	// Split after an empty string to get all letters.
	result7 := strings.SplitAfter(value7, "")
	for i := range result7 {
		// Get letter and display it.
		letter := result7[i]
		fmt.Println(letter)
		/*
			b
			i
			r
			d
		*/
	}

	// exemplo 8
	value8 := "12|34|56|78"
	// Split into three parts.
	// ... The last separator is not split.
	result8 := strings.SplitN(value8, "|", 3)
	for v := range result8 {
		fmt.Println(result8[v])
		/*
			12
			34
			56|78
		*/
	}
}
