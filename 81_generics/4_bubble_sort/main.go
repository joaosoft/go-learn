package main

import "fmt"

func BubbleSort(input []int) []int {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

type Number interface {
	int16 | int32 | int64 | float32 | float64
}

func BubbleSortGeneric[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {
	fmt.Println("Go Generics Tutorial")
	listInt := []int{4, 3, 1, 5}
	listInt32 := []int32{4, 3, 1, 5}
	listFloat64 := []float64{4.3, 5.2, 10.5, 1.2, 3.2}

	sorted := BubbleSort(listInt)
	fmt.Println(sorted)

	sortedGenericInt32 := BubbleSortGeneric(listInt32)
	fmt.Println(sortedGenericInt32)

	sortedGenericFloat64 := BubbleSortGeneric(listFloat64)
	fmt.Println(sortedGenericFloat64)
}
