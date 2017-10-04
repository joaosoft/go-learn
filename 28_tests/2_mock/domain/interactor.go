package domain

// Interactor ...
type Interactor struct {
	repository IRepository
}

// Interactor ...
func NewInteractor(repository Repository) *Interactor {
	interactor := new(Interactor)
	interactor.repository = repository
	return interactor
}

// DoSomething ...
func (interactor *Interactor) DoSomething(id string, value string) error {
	return interactor.repository.DoSomething(id, value)
}
