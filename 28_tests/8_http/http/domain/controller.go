package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":1323", nil)
}

func SayHello(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		bytes, _ := json.Marshal(error)
		w.Write(bytes)
		return
	}

	fmt.Printf("received body: %s", body)
	bytes, _ := json.Marshal(Person{
		Name: name,
		Age:  30,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
