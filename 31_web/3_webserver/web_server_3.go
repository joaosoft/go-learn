package main

import (
	"github.com/joaosoft/golang-learn/31_web/3_webserver/controllers"
	"github.com/joaosoft/golang-learn/31_web/3_webserver/middleware"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.Use(middleware.CORS(middleware.CORSOptions{}))

	controllers.InitXPTOResources(r)

	log.Info("Service starting on port 8080")

	r.Run(":8080") // listen and serve
}
