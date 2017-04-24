package interactors

import (
	"fmt"
	"productoffers/services/mirakl-mci/vo"
)

// MCIListenerInteractor ...
type MCIListenerInteractor struct {
	repository IMCIRepository
}

// NewMCIListenerInteractor ...
func NewMCIListenerInteractor(
	repository IMCIRepository,
) *MCIListenerInteractor {
	interactor := new(MCIListenerInteractor)
	interactor.repository = repository
	return interactor
}

// HandleMessage ...
func (interactor *MCIListenerInteractor) HandleMessage(msg *vo.ReportMessage) error {
	if state, ok := vo.States[msg.State]; !ok {
		return fmt.Errorf("unrecognized state sent in message: [state:%s]", msg.State)
	} else {
		return interactor.repository.UpdateRecordState(msg.ImportID, state)
	}
}
