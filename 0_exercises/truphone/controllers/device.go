package controllers

import (
	"github.com/gorilla/mux"
	"go-learn/0_exercises/truphone/models"
	"json"
	"net/http"
	"time"
)

var(
	ErrorUnexpected = []byte(`{"message": "unexpected error!"}`)
	SuccessMessage = []byte(`{"message": true}`)
)

func Create(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request CreateDeviceRequest
	defer req.Body.Close()

	var body []byte
	if _, err := req.Body.Read(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var bytes []byte
		bytes, err = json.Marshal(ErrorResponse{
			Message: err.Error(),
		})

		if err != nil {
			_, _ = w.Write(ErrorUnexpected)
			return
		}
		_, _ = w.Write(bytes)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var bytes []byte
		bytes, err = json.Marshal(ErrorResponse{
			Message: err.Error(),
		})

		if err != nil {
			_, _ = w.Write(ErrorUnexpected)
			return
		}
		_, _ = w.Write(bytes)
		return
	}

	// TODO: validate the request

	device := &models.Device{
		Name: request.Name,
		Type: models.DeviceType(request.Type),
		CreatedAt: time.Now(),
	}

	id, err := device.Create()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var bytes []byte
		bytes, err = json.Marshal(ErrorResponse{
			Message: err.Error(),
		})

		if err != nil {
			_, _ = w.Write(ErrorUnexpected)
			return
		}
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := CreateResponse {
		Id: id,
	}

	var bytes []byte
	bytes, err = json.Marshal(response)

	if err != nil {
		_, _ = w.Write(ErrorUnexpected)
		return
	}

	_, _ = w.Write(bytes)
}

func Update(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request UpdateDeviceRequest
	defer req.Body.Close()

	var body []byte
	if _, err := req.Body.Read(body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var bytes []byte
		bytes, err = json.Marshal(ErrorResponse{
			Message: err.Error(),
		})

		if err != nil {
			_, _ = w.Write(ErrorUnexpected)
			return
		}
		_, _ = w.Write(bytes)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var bytes []byte
		bytes, err = json.Marshal(ErrorResponse{
			Message: err.Error(),
		})

		if err != nil {
			_, _ = w.Write(ErrorUnexpected)
			return
		}
		_, _ = w.Write(bytes)
		return
	}

	vars := mux.Vars(req)
	request.Id = vars["id"]

	// TODO: validate the request

	device := &models.Device{
		Id:  request.Id,
		Name: request.Name,
		Type: models.DeviceType(request.Type),
		CreatedAt: time.Now(),
	}

	if err := device.Update(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var bytes []byte
		bytes, err = json.Marshal(ErrorResponse{
			Message: err.Error(),
		})

		if err != nil {
			_, _ = w.Write(ErrorUnexpected)
			return
		}
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(SuccessMessage)
}

func Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request DeleteDeviceRequest

	vars := mux.Vars(req)
	request.Id = vars["id"]

	// TODO: validate the request

	device := &models.Device{
		Id:  request.Id,
	}

	if err := device.Delete(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var bytes []byte
		bytes, err = json.Marshal(ErrorResponse{
			Message: err.Error(),
		})

		if err != nil {
			_, _ = w.Write(ErrorUnexpected)
			return
		}
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(SuccessMessage)
}