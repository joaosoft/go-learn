/*
Create a function called searchItem() that takes 2 parameters: a) a string slice and b) a string.

The function should search for the string (the second parameter)
in the slice (the first parameter) and returns true if it finds the string in the slice and false otherwise.
Do function does an case-sensitive search.

Call the function and see how it works.

Example:

animals := []string{"lion", "tiger", "bear"}
result := searchItem(animals, "bear")
fmt.Println(result) // => true
result = searchItem(animals, "pig")
fmt.Println(result)     // => false
*/

package main

import "fmt"

func searchItem(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if v == myStr {
			return true
		}

	}
	return false
}

func main() {

	animals := []string{"lion", "tiger", "bear"}

	result := searchItem(animals, "bear")
	fmt.Println(result) // => true

	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

}
