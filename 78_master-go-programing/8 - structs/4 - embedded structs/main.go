/////////////////////////////////
// Anonymous and Embedded Structs
// Go Playground: https://play.golang.org/p/NtH6I30gtxb
/////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	//** EMBEDDED STRUCTS **//
	// An embedded struct is just a struct that acts like a field inside another struct.

	// define a new struct type
	type Contact struct {
		email, address string
		phone          int
	}

	// define a struct type that contains another struct as a field
	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	// declaring a value of type Employee
	john := Employee{
		name:   "John Keller",
		salary: 3000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   042324234,
		},
	}

	fmt.Printf("%+v\n", john)
	// => {name:John Keller salary:3000 contactInfo:{email:jkeller@company.com address:Street 20, London phone:295619381404}}

	// accessing a field
	fmt.Printf("Employee's salary: %d\n", john.salary)

	// accessing a field from the embedded struct
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // => Employee's email:jkeller@company.com

	// updating a field
	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email)
	// =>  Employee's new email address:new_email@thecompany.com
}
