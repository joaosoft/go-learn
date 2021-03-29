/*
Change the function from the previous exercise to do a case-insensitive search.

Example:

animals := []string{"Lion", "tiger", "bear"}
result := searchItem(animals, "beaR")
fmt.Println(result) // => true
result = searchItem(animals, "lion")
fmt.Println(result)     // => true
*/

package main

import (
	"fmt"
	"strings"
)

func searchItem(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if strings.EqualFold(v, myStr) {
			return true
		}

	}
	return false
}

func main() {

	animals := []string{"lion", "tiger", "bear"}

	result := searchItem(animals, "BEAR")
	fmt.Println(result) // => true

	result = searchItem(animals, "lION")
	fmt.Println(result) // => true

	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

}
