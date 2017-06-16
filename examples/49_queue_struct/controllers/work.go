package controllers

import (
	"github.com/labstack/gommon/log"
)

type Work struct {
	data []byte
}

func NewWork(data []byte) *Work {
	work := Work{
		data: data,
	}

	return &work
}

func (work *Work) GetWork() ([]byte, error) {
	log.Infof("GetWork()")
	log.Infof(string(work.data))

	return work.data, nil
}
