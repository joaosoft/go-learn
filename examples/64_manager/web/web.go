package web

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// IWebController ... web controller interface
type IWebController interface {
	AddRoute(method string, route string, handler func(context echo.Context) error)
	Start() error
	Stop() error
}

// WebController ... web controller structure
type WebController struct {
	httpServer *echo.Echo
	Config     *Config
	started    bool
}

// NewWebController ... create a new WebController
func NewWebController(config *Config) (IWebController, error) {

	echo := echo.New()

	return &WebController{
		httpServer: echo,
		Config:     config,
	}, nil
}

// AddRoute ... adds a new route handler
func (instance *WebController) AddRoute(method string, route string, handler func(context echo.Context) error) {
	log.Infof("web, adding '%s' method to route '%s'", method, route)
	switch method {
	case "POST":
		instance.httpServer.POST(route, handler)
	case "PUT":
		instance.httpServer.PUT(route, handler)
	case "GET":
		instance.httpServer.GET(route, handler)
	case "HEAD":
		instance.httpServer.HEAD(route, handler)
	case "PATCH":
		instance.httpServer.PATCH(route, handler)
	case "OPTIONS":
		instance.httpServer.OPTIONS(route, handler)
	}
}

// Start ... starts the server
func (instance *WebController) Start() error {
	instance.started = true

	log.Infof("web, starting webserver [address:%s]", instance.Config.Address)
	if err := instance.httpServer.Start(instance.Config.Address); err != nil {
		log.Errorf("web, error starting webserver [address:%s], %s", instance.Config.Address, err.Error())
		instance.started = false
		return err
	}

	return nil
}

// Stop ... stops the server
func (instance *WebController) Stop() error {
	if !instance.started {
		instance.started = false
		return nil
	}

	return instance.httpServer.Server.Shutdown(context.Background())
}
