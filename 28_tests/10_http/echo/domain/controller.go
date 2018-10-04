package domain

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/hello/:name", SayHello)
	e.Logger.Fatal(e.Start(":1323"))
}

func SayHello(context echo.Context) error {

	name := context.Param("name")

	body, error := ioutil.ReadAll(context.Request().Body)
	if error != nil {
		return error
	}

	fmt.Printf("received body: %s", body)

	return context.JSON(http.StatusOK, Person{
		Name: name,
		Age:  30,
	})
}
