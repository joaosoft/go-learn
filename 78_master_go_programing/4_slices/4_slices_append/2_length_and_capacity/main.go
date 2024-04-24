/////////////////////////////////
// Appending to a Slice
// Go Playground: https://play.golang.org/p/WYrW0Y_FeEF
/////////////////////////////////

package main

import "fmt"

func main() {
	//** Slice's Length and Capacity **//

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 1, Capacity: 1

	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 5, Capacity: 8

	nums = append(nums[0:4], 100, 200, 300, 400, 500)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 9, Capacity: 16 (always duplicates)
}
