package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Create a struct to match the format of JSON
type Person struct {
	Name string `json:"Name"`
	City string `json:"City"`
}

// create a single instance of Person
var person Person

// create an array of Person types
var people []Person

func main() {
	load1()
	load2()
	load3()
}

func load1() {
	fmt.Println(":::::::::::::::: LOAD 1")

	// JSON string to parse, see below for example read in from file
	json_str := "{ \"Name\": \"Marcus\", \"City\": \"San Jose\"}"

	// JSON Unmarshal command takes []byte, so string needs to be cast
	// The second parameter is a pointer to the struct that matches format
	if err := json.Unmarshal([]byte(json_str), &person); err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	// output result
	fmt.Printf("Name: %v, City: %v\n", person.Name, person.City)
}

func load2() {
	fmt.Println(":::::::::::::::: LOAD 2")

	// read json in from a file
	dir, err := os.Getwd()
	fmt.Println("PATH:", dir, "/29_json/people.json")

	file, err := ioutil.ReadFile("./29_json/people.json")
	if err != nil {
		fmt.Println("Error reading file")
	}

	// the names.json file has an array of person objects, so read into people
	if err := json.Unmarshal(file, &people); err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	// output result
	fmt.Println(people)

	// encoding a Go object into JSON is simply using the Marshal command
	json, err := json.Marshal(people)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println(string(json))
}

func load3() {
	fmt.Println(":::::::::::::::: LOAD 3")

	// read json in from a file
	dir, err := os.Getwd()
	fmt.Println("PATH:", dir, "/29_json/person.json")

	file, err := ioutil.ReadFile("./29_json/person.json")
	if err != nil {
		fmt.Println("Error reading file")
	}

	// the names.json file has an array of person objects, so read into people
	if err := json.Unmarshal(file, &person); err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	// output result
	fmt.Println(person)

	// encoding a Go object into JSON is simply using the Marshal command
	json, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println(string(json))
}
