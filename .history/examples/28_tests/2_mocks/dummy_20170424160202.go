package dummy

import (
	"fmt"
)

// TesteStruct ...
type TesteStruct struct {
	repository IMCIRepository
}

// NewTesteStruct ...
func NewTesteStruct(
	repository IMCIRepository,
) *TesteStruct {
	interactor := new(TesteStruct)
	interactor.repository = repository
	return interactor
}

// HandleMessage ...
func (interactor *TesteStruct) HandleMessage(msg *vo.ReportMessage) error {
	if state, ok := vo.States[msg.State]; !ok {
		return fmt.Errorf("unrecognized state sent in message: [state:%s]", msg.State)
	} else {
		return interactor.repository.UpdateRecordState(msg.ImportID, state)
	}
}
