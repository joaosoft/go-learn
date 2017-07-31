package pm

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/web"
)

// -------------- WEB SERVERS --------------
// NewWEBServer ... creates a new web server
func (instance *manager) NewWEBServer(config *web.Config) (web.IWebController, error) {
	log.Infof(fmt.Sprintf("web, creating web server"))
	return web.NewWebController(config)
}

// -------------- METHODS --------------
func (instance *manager) AddWebServer(key string, webServer web.IWebController) {
	log.Infof(fmt.Sprintf("web, adding a new web server '%s'", key))
	instance.httpController[key] = webServer
}

// AddWebServerRoute ... add a new route to the web server
func (instance *manager) AddWebServerRoute(key string, method string, route string, handler func(context echo.Context) error) {

	log.Infof(fmt.Sprintf("web, %s web server adding a new route '%s' = '%s' added", key, method, route))
	instance.httpController[key].AddRoute(method, route, handler)
}

// Start ... starts the web server
func (instance *manager) StartWebServer(key string) error {

	log.Infof(fmt.Sprintf("web, starting web server '%s' starting", key))
	return instance.httpController[key].Start()
}

// Stop ... stops the web server
func (instance *manager) StopWebServer(key string) error {

	log.Infof(fmt.Sprintf("web, starting web server '%s' stopping", key))
	return instance.httpController[key].Stop()
}
