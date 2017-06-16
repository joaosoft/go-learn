package controllers

import (
	"github.com/labstack/gommon/log"
)

type Controller struct {
	data []byte
}

func NewController(data []byte) *Controller {
	controller := Controller{
		data: data,
	}

	return &controller
}

func (controller *Controller) Do() error {
	log.Infof("Do()")
	log.Infof(string(controller.data))
	return nil
}

func (controller *Controller) Undo() error {
	log.Infof("Undo()")
	return nil
}
