// @APIVersion 1.0.0
// @APITitle Users Service
// @APIDescription Users Service
// @Contact joaosoft@gmail.com
// @TermsOfServiceUrl
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
package swagger

import (
	"net/http"

	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"html/template"
	"os"
	"search-and-navigation/docs/swagger"
	"strings"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)

	e.GET("/", indexHandler)
	dir, _ := os.Getwd()
	e.Static("/swagger-ui", dir+"/examples/62_swagger/swagger/swagger-ui")
	e.GET("/:descriptor", descriptorHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
// @Title Get User By ID
// @Description Get all user by ID
// @Accept json
// @Produce json
// @Param ID "USER ID"
// @Success 200 {object} "Success"
// @Failure 404 {object} "Not Found"
// @Failure 500 {object} "Unexpected error"
// @Resource /users
// @Router "/users/{userid}" [get]
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.POST("/save", save)
// @Title Save User
// @Description Save user
// @Accept json
// @Produce json
// @Param name "NAME"
// @Param email "EMAIL"
// @Success 200 {object} "Success"
// @Failure 500 {object} "Unexpected
// @Resource /users
// @Router "/users" [post]
func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

// README
// $ curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
// => name:Joe Smith, email:joe@labstack.com

func indexHandler(ctx echo.Context) error {
	isJsonRequest := false

	if acceptHeaders, ok := ctx.Request().Header["Accept"]; ok {
		for _, acceptHeader := range acceptHeaders {
			if strings.Contains(acceptHeader, "json") {
				isJsonRequest = true
				break
			}
		}
	}

	buf := bytes.NewBuffer(make([]byte, 0, 10))
	t, e := template.New("swagger_index").Parse(swagger.ResourceListingJson)
	if e != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	t.Execute(buf, "http://localhost:1323")

	if isJsonRequest {
		var response interface{}
		json.Unmarshal(buf.Bytes(), &response)

		return ctx.JSON(http.StatusOK, response)
	}

	return ctx.Redirect(http.StatusFound, "/swagger-ui/")
}

func descriptorHandler(ctx echo.Context) error {
	buf := bytes.NewBuffer(make([]byte, 0, 10))
	apiKey := ctx.Param("descriptor")

	if json, ok := swagger.ApiDescriptionsJson[apiKey]; ok {
		t, e := template.New("swagger_descriptor").Parse(json)
		if e != nil {
			return ctx.JSON(http.StatusInternalServerError, nil)
		}

		t.Execute(buf, "http://localhost:1323/")

		return ctx.HTML(http.StatusOK, buf.String())
	}

	return ctx.JSON(http.StatusNotFound, nil)
}
