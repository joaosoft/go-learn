package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/26_web/3_webserver/controllers"
	"golang-learn/examples/26_web/3_webserver/middleware"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.Use(middleware.CORS(middleware.CORSOptions{}))

	controllers.InitXPTOResources(r)

	log.Info("Service starting on port 8080")

	r.Run(":8080") // listen and serve
}
