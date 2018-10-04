package domain

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

func main() {

	controller := Controller{}
	beego.Router("/hello/:joao", &controller, "post:SayHello")
	beego.Run(":1323")
}

type Controller struct {
	beego.Controller
}

func (controller *Controller) SayHello() {
	//defer controller.ServeJSON()

	name := controller.GetString("name")

	fmt.Printf("received body: %s", string(controller.Ctx.Input.RequestBody))

	controller.Ctx.Output.SetStatus(http.StatusOK)
	controller.Ctx.Output.JSON(Person{
		Name: name,
		Age:  30,
	}, false, false)
}
