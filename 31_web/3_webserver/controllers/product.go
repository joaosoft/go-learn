package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type XPTOController struct {
}

func InitXPTOResources(e *gin.Engine) {
	u := XPTOController{}
	// Setup Routes
	e.GET("/xpto", u.HandleXPTO)
}

func (r *XPTOController) HandleXPTO(c *gin.Context) {

	out := "hello, i'im in!"
	c.JSON(http.StatusOK, out)

	return
}
