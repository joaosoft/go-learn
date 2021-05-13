package models

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

var(
	devices = make(map[string]*Device, 0)
)

type DeviceList []Device
type Device struct {
	Id string `json:"id"`
	Type DeviceType `json:"type"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"date"`
}

type DeviceType string

const(
	DeviceSmartphone DeviceType = "smartphone"
	DeviceTablet DeviceType = "tablet"
	DeviceIot DeviceType = "iot"
)

func (d *Device) Create() (string, error) {
	uid := uuid.NewV4()
	devices[uid.String()] = d
	return uid.String(),  nil
}

func (d *Device) Update() error {
	if _, ok := devices[d.Id]; ok {
		devices[d.Id].Name = d.Name
		devices[d.Id].Type = d.Type
		return nil
	}

	return fmt.Errorf("not found!")
}

func (d *Device) Delete() error {
	if _, ok := devices[d.Id]; ok {
		delete(devices, d.Id)
		return nil
	}

	return fmt.Errorf("not found!")
}