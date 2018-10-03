package domain

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repository *Repository) DoSomething(name string, age int) (int64, int64, error) {
	tx, err := repository.db.Begin()
	if err != nil {
		return 0, 0, err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	result, err := tx.Exec("INSERT INTO something (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		return 0, 0, err
	}

	id, _ := result.LastInsertId()
	rows, _ := result.RowsAffected()

	return id, rows, nil
}
