package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {

		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v  \n", err)

		// Sending the string over the channel.
		// This is a blocking call, so this goroutine will
		// wait for the main goroutine to receive it on the other part of the channel.
		c <- s
	} else {

		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			file := strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("Writing response Body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {

				s += "Error writing to file!\n"

				// sending s over the channel
				c <- s
			}
		}
		s += fmt.Sprintf("%s is UP\n", url)

		// sending s over the channel
		c <- s
	}
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com"}

	// Declaring a new channel
	c := make(chan string)

	for _, url := range urls {
		// Launching the goroutines
		go checkAndSaveBody(url, c)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => 4

	// Receiving the messages from the channel
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

// Run the program: go run main.go

//** EXPECTED OUTPUT: **//
// No. of Goroutines: 4
// Status Code: 200
// Writing response Body to www.google.com.txt
// https://www.google.com is UP

// Status Code: 200
// Writing response Body to www.golang.org.txt
// https://www.golang.org is UP

// Status Code: 200
// Writing response Body to www.medium.com.txt
// https://www.medium.com is UP
