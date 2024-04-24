/////////////////////////////////
// Appending to a Slice
// Go Playground: https://play.golang.org/p/WYrW0Y_FeEF
/////////////////////////////////

package main

import "fmt"

func main() {
	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")
	fmt.Println(letters)
	fmt.Println(letters[0])
	//fmt.Println(letters[4]) // ends in error
	fmt.Println(letters[3:6])                                             // why? it still references the backing array!
	fmt.Printf("Length: %d, Capacity: %d \n", len(letters), cap(letters)) // Length: 3, Capacity: 6 (always duplicates)

	letters2 := append(letters, "END")
	fmt.Println("letters2 := append(letters, \"END\"); letters2:", letters2)

	letters2[0] = "START"
	fmt.Println("letters2[0] = \"START\"; letters2:", letters2)
	fmt.Println("letters2[0] = \"START\"; letters:", letters)

	n1 := []int{10, 20, 30, 40, 50}
	fmt.Printf("Length: %d, Capacity: %d \n", len(n1), cap(n1)) // Length: 5, Capacity: 5 (always duplicates)

	n2 := append(n1[0:4], 100)                                  // 'n1[0:4]' this way it will use the same memory! but not this way 'n1[0:5]' or 'n1'
	fmt.Printf("Length: %d, Capacity: %d \n", len(n2), cap(n2)) // Length: 6, Capacity: 10 (always duplicates)

	n1[0] = 66
	fmt.Println(n1) // [66 20 30 40 50]
	fmt.Println(n2) // [10 20 30 40 50 100]
}
