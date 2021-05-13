package main

import (
	"fmt"
	"go-learn/0_exercises/truphone/routes"
	"net/http"
)

func main() {
	fmt.Println("server started at http://localhost:8081/")
	if err := http.ListenAndServe(":8081", routes.Router); err != nil {
		panic(err)
	}
}
