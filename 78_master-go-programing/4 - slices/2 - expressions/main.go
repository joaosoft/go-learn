/////////////////////////////////
// Slices Expressions
// Go Playground: https://play.golang.org/p/Wfb7Lpq2WxK
/////////////////////////////////

package main

import "fmt"

func main() {

	// arrays, slices and strings are sliceable
	// slicing doesn't modify the array or the slice, it returns a new one

	// declaring an [5]int array
	a := [5]int{1, 2, 3, 4, 5}

	// a slice expression is formed by specifying a start or a low bound and a stop or high bound like  a[start:stop].
	// this selects a range of elements which includes the element at index start, but excludes the element at index stop.

	// slicing an array returns a slice, not an array
	b := a[1:3]                                 // 1 is called start (included), 3 is called stop (excluded)
	fmt.Printf("Type: %T , Value: %#v\n", b, b) // => Type: []int , Value: []int{2, 3}

	// declaring a slice
	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:3]   //start included, stop excluded
	fmt.Println(s2) //[2 3]

	//for convenience, any of the indexes may be omitted.
	// a missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
	fmt.Println(s1[2:])       // => [3 4 5 6], same as s1[2:len(s1)]
	fmt.Println(s1[:3])       // => [1 2 3], same as s1[0:3]
	fmt.Println(s1[:])        // => [1 2 3 4 5 6], same with s1[0:len(s1)]
	fmt.Println(s1[:len(s1)]) // => => [1 2 3 4 5 6], returns the entire slice
	// fmt.Println(s1[:45])   //panic: runtime error: slice bounds out of range

	s1 = append(s1[:4], 100) // adds 100 after index 4 (excluded)
	fmt.Println(s1)          // -> [1 2 3 4 100]

	s1 = append(s1[:4], 200) // overwrites the last element
	fmt.Println(s1)          // -> [1 2 3 4 200]
}
