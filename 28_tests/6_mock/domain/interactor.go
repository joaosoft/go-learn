package domain

type Interactor struct {
	repository IRepository
}

func NewInteractor(repository IRepository) *Interactor {
	return &Interactor{
		repository: repository,
	}
}

func (interactor *Interactor) DoSomething(name string, age int) error {
	return interactor.repository.Store(name, age)
}
