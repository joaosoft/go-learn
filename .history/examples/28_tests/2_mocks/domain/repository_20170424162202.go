package dummy

import (
	"productoffers/common/session"
)

type Repository struct {
	data string
	name string
}

// NewRepository creates a new Repository.
func NewRepository(data string) *Repository {
	repository := new(Repository)
	repository.data = data
	repository.name = "dummy"

	return repository
}

// SetDBContext ...
func (repository *Repository) SetDBContext(ctx session.IDBContext) {
	repository.conn = ctx
}

// DoSomething ...
func (repository *Repository) DoSomething(id string, value string) error {
	if _, err := repository.conn.Exec(`
		INSERT INTO dummy(id, value)
			VALUES($1, $2)
	`,
		id,
		value,
	); err != nil {
		return err
	}

	return nil
}
