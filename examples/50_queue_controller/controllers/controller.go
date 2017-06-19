package controllers

import (
	"github.com/labstack/gommon/log"
	"golang-learn/examples/50_queue_controller/repositories"
)

type Controller struct {
	repository repositories.Repository
	data interface{}
}

func NewController(repository repositories.Repository) *Controller {
	controller := Controller{
		repository: repository,
	}

	return &controller
}

func (controller *Controller) Do(data interface{}) error {

	return controller.repository.DoSomething(data)
}

func (controller *Controller) Undo() error {
	log.Infof("Undo()")

	return nil
}
