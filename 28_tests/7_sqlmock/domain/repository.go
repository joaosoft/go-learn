package domain

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repository *Repository) DoSomethingSelect(name string) ([]*Person, error) {
	var persons []*Person

	rows, err := repository.db.Query("SELECT name, age FROM something WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		person := &Person{}

		if err := rows.Scan(&person.Name, &person.Age); err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (repository *Repository) DoSomethingInsert(name string, age int) (int64, int64, error) {
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

func (repository *Repository) DoSomethingUpdate(name string, age int) (bool, error) {
	tx, err := repository.db.Begin()
	if err != nil {
		return false, err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	result, err := tx.Exec("UPDATE something SET age = ? WHERE name = ?", age, name)
	if err != nil {
		return false, err
	}

	rows, _ := result.RowsAffected()

	return rows > 0, nil
}
