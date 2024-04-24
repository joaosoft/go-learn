package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {

	// calling http.Get() which gives a response and an error value.
	resp, err := http.Get(url)

	// error handling
	if err != nil {
		fmt.Printf("%s is DOWN!\n", url)
	} else {
		// a good practice is to close the Response Body if there is one
		// If you forget to close it there will be a resource leak.
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d  \n", url, resp.StatusCode)

		// fetching the page if HTTP Status Code is 200 (success)
		if resp.StatusCode == 200 {

			// The resp.Body implements the io.Reader interface
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			// Creating the file's name
			file := strings.Split(url, "//")[1]
			file += ".txt" // and I concatenate .txt to file value

			fmt.Printf("Writing response Body to %s\n", file)

			// Writing the response Body to File
			// If the file doesn't exist WriteFile() creates it and if it already exists
			// the function will truncate it before writing to file.
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	urls := []string{"https://www.golang1.org", "https://www.google.com", "https://www.medium.com"}

	// Iterating over the URLs and call the function for each URL
	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}

// Run the program: go run main.go

//** EXPECTED OUTPUT: **//
// https://www.golang1.org is DOWN!
// ####################
// https://www.google.com -> Status Code: 200
// Writing response Body to www.google.com.txt
// ####################
// https://www.medium.com -> Status Code: 200
// Writing response Body to www.medium.com.txt
// ####################
