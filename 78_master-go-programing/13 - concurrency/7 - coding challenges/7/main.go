/*
Create a function literal (a.k.a. anonymous function)
that sends the string value if receives as argument to main func using a channel.
*/

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func(n string) {
		ch <- n
	}("Gophers!")

	value := <-ch
	fmt.Println("Value received from channel:", value)
}
