package controllers

import (
	"github.com/labstack/gommon/log"
)

type Controller struct {
}

func NewController() *Controller {
	controller := Controller{}

	return &controller
}

func (controller *Controller) Do(data []byte) error {
	log.Infof("GetWork()")
	log.Infof(string(data))

	return nil
}
