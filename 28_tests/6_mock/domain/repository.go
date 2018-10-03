package domain

type Repository struct{}

func NewRepository() IRepository {
	return &Repository{}
}

func (repository *Repository) Store(name string, age int) error {
	// do something...

	return nil
}
