package dummy

type Repository struct {
	data string
	name string
}

// NewRepository creates a new Repository.
func NewRepository(name string, data string) *Repository {
	repository := new(Repository)
	repository.data = data
	repository.name = name

	return repository
}

// DoSomething ...
func (repository *Repository) DoSomething(id string, value string) error {
	// do something

	return nil
}
