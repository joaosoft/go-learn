package controllers

import (
	"github.com/joaosoft/golang-learn/48_queue/3_queue_controller/repositories"

	"github.com/labstack/gommon/log"
)

type Controller struct {
	repository repositories.Repository
	data       interface{}
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
