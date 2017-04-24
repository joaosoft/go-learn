package dummy

import (
	"fmt"
)

// TesteStruct ...
type TesteStruct struct {
	repository Repository
}

// NewTesteStruct ...
func NewTesteStruct(
	repository Repository,
) *TesteStruct {
	interactor := new(TesteStruct)
	interactor.repository = repository
	return interactor
}

// HandleMessage ...
func (interactor *TesteStruct) DoSomething(importID string, account string) error {
	return interactor.repository.DoSomething(msg.ImportID, state)
}
