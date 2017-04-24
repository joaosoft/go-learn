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
	if state, ok := vo.States[msg.State]; !ok {
		return fmt.Errorf("unrecognized state sent in message: [state:%s]", msg.State)
	} else {
		return interactor.repository.UpdateRecordState(msg.ImportID, state)
	}
}
