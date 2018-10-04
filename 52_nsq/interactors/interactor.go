package interactors

import "fmt"

type Interactor struct {
}

// NewInteractor ...
func NewInteractor() *Interactor {
	interactor := new(Interactor)

	return interactor
}

// DoSomethingInsert ...
func (interactor *Interactor) DoSomething() error {
	fmt.Println("OLA, ENTREI NO INTERACTOR!")
	return nil
}
