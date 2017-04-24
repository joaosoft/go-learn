package dummy

// TesteStruct ...
type TesteStruct struct {
	repository Repository
}

// NewTesteStruct ...
func NewTesteStruct(
	repository Repository,
) *TesteStruct {
	controller := new(TesteStruct)
	interactor.repository = repository
	return interactor
}

// DoSomething ...
func (interactor *TesteStruct) DoSomething(id string, value string) error {
	return interactor.repository.DoSomething(id, value)
}
