package controllers

import (
	"json"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	bytes, _ := json.Marshal(
		HelloWorldResponse{
			Message: "hello world",
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bytes)
}
