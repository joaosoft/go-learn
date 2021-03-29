/*
Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

Using copy() function create a copy of the slice.
Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
*/

package main

import "fmt"

func main() {
	// Using copy()
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

	// Using append() appends elements to the end of a slice
	friends = []string{"Marry", "John", "Paul", "Diana"}
	yourFriends = []string{}

	yourFriends = append(yourFriends, friends...)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

}
