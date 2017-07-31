package pm

import (
	"fmt"
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
