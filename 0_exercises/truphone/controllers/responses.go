package controllers

type HelloWorldResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
type CreateResponse struct {
	Id string `json:"id"`
}