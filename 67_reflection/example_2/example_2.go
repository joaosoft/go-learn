package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/satori/go.uuid"
)

// Retorno ...
type Retorno struct{}

// MyTest ...
type MyTest struct{}

// NewMyTest ...
func NewMyTest() *MyTest {
	instance := &MyTest{}
	return instance
}

// Get ...
func (instance *MyTest) Get(test string, rtnType reflect.Type) (map[uuid.UUID]interface{}, error) {
	var data []byte
	response := make(map[uuid.UUID]interface{})

	newInstance := reflect.New(rtnType).Interface()
	if err := json.Unmarshal(data, &newInstance); err != nil {
		panic("failed to unmarshal to instance")
	}
	id, _ := uuid.NewV4()
	response[id] = newInstance

	return response, nil
}

func main() {
	instance := NewMyTest()
	instance.Get("teste", reflect.TypeOf(Retorno{}))
	fmt.Println("FEITO")
}
