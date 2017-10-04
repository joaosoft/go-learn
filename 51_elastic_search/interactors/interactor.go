package interactors

import (
	"golang-learn/examples/51_elastic_search/domain"
	"golang-learn/examples/51_elastic_search/repositories"
	"os"
	"golang-learn/examples/51_elastic_search/common/config"
	"github.com/labstack/gommon/log"
)

var _mapping map[string]interface{}

func init() {
	if err := config.LoadConfigFromPath("mapping", &_mapping); err != nil {
		log.Error("Error loading config: ", err)
		os.Exit(0)
	}
}

type Interactor struct {
	Repository repositories.IRepository
}

// NewInteractor ...
func NewInteractor(repository repositories.IRepository) *Interactor {
	interactor := new(Interactor)
	interactor.Repository = repository

	return interactor
}

// DoSomething ...
func (interactor *Interactor) CreateIndex(index string) error {
	return interactor.Repository.CreateIndex(index, _mapping)
}

// DoSomething ...
func (interactor *Interactor) Insert(data []domain.Something) error {
	return interactor.Repository.Insert(data)
}
