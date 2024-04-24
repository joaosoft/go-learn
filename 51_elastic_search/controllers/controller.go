package controllers

import (
	"github.com/joaosoft/golang-learn/51_elastic_search/domain"
	"github.com/joaosoft/golang-learn/51_elastic_search/interactors"
)

type Controller struct {
	Interactor interactors.Interactor
}

func (controller *Controller) CreateIndex(index string) error {
	return controller.Interactor.CreateIndex(index)
}

func (controller *Controller) Insert(data []domain.Something) error {
	return controller.Interactor.Insert(data)
}
