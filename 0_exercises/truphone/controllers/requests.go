package controllers

type CreateDeviceRequest struct {
	Type DeviceType `json:"type"`
	Name string `json:"name"`
}

type UpdateDeviceRequest struct {
	Id string `json:"id"`
	Type DeviceType `json:"type"`
	Name string `json:"name"`
}

type DeleteDeviceRequest struct {
	Id string `json:"id"`
}


type DeviceType string

const(
	DeviceSmartphone DeviceType = "smartphone"
	DeviceTablet DeviceType = "tablet"
	DeviceIot DeviceType = "iot"
)