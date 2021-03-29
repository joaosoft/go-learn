/*
Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}
Using append() function create a copy of the slice.
Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
*/

package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := []string{}

	yourFriends = append(yourFriends, friends...)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

}
