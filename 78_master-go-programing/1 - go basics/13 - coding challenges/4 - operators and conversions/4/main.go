/*
Create a Go program that computes how long does it take for the Sunlight to reach the Earth if we know that the distance from the Sun to Earth is 149.6 million km and the speed of light  is 299,792,458 m / s

The formula used is: time = distance / speed
*/
package main

import "fmt"

func main() {
	const (
		distance = 149_600_000 * 1000 // distance from the Sun to Earth in m
		// (it's allowed to use underscores in numbers for a better readability)

		speed = 299_792_458 // speed of light in m / s
	)

	const time = distance / speed // time in seconds

	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time)

}
