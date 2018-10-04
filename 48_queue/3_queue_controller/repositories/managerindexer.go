package repositories

import (
	"github.com/labstack/gommon/log"
)

type Repository struct {
}

func NewRepository() *Repository {
	repository := Repository{}

	return &repository
}

func (repository *Repository) DoSomething(data interface{}) error {
	log.Infof("Repository: DoSomethingInsert")

	return nil
}
