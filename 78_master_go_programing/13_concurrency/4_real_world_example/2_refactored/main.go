package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s is DOWN!\n", url)
	} else {

		fmt.Printf("Status Code: %d  ", resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing response Body to %s\n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	// 3. Inside each goroutine, call wg.Done() to indicate to the WaitGroup that the goroutine has finished executing.
	wg.Done()
}

func main() {

	urls := []string{"https://www.golang.org", "https://www.google1.com", "https://www.medium.com"}

	// 1.
	// Create a new instance of  sync.WaitGroup
	var wg sync.WaitGroup

	// 2.
	// call the waitgroup.Add(n) method just before attempting to spawn the goroutine.
	wg.Add(len(urls)) //  n is len(urls) - the number of goroutines to wait for

	for _, url := range urls {
		// 4.
		// Launch the goroutines
		go checkAndSaveBody(url, &wg)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	// 5.
	// Finally, call waitgroup.Wait() to block the execution of main() until the goroutines
	// in the waitgroup have successfully completed.
	wg.Wait()
}

// Run the program: go run main.go

//** EXPECTED OUTPUT: **//
// No. of Goroutines: 4
// https://www.google1.com is DOWN!
// Status Code: 200  Writing response Body to www.golang.org.txt
// Status Code: 200  Writing response Body to www.medium.com.txt
