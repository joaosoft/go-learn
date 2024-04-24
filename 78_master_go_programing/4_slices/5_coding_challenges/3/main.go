/*
1. Declare a slice called nums with three float64 numbers.

2. Append the value 10.1 to the slice

3. In one statement append to the slice the values: 4.1,  5.5 and 6.6

4. Print out the slice

5. Declare a slice called n with two float64 values

6. Append n to nums

7. Print out the nums slice
*/

package main

import "fmt"

func main() {
	// 1. Declare a slice called nums with 3 float64 numbers.
	nums := []float64{1.1, 2.2, 3.3}

	// 2. Append the value 10.1 to the slice
	nums = append(nums, 10.1)

	// 3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
	nums = append(nums, 4.1, 5.5, 6.6)

	// 4. Print out the slice
	fmt.Println(nums)

	// 5. Declare a slice called n with 2 float64 values
	n := []float64{10.10, 20.20}

	// 6. Append n to nums
	nums = append(nums, n...)

	// 7. Print out the slice
	fmt.Println(nums)

}
