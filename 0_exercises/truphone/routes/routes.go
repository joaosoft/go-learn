package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaosoft/golang-learn/0_exercises/truphone/controllers"
	"net/http"
)

var (
	Router = mux.NewRouter()
)

func init() {
	Router.HandleFunc("/hello-world", controllers.HelloWorld).Methods(http.MethodGet)
	Router.HandleFunc("/devices", controllers.Create).Methods(http.MethodPost)
	Router.HandleFunc("/devices/{id}", controllers.Update).Methods(http.MethodPut)
	Router.HandleFunc("/devices/{id}", controllers.Delete).Methods(http.MethodDelete)
}
