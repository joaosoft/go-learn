package echo

import (
	"go-learn/28_tests/8_http/echo/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"fmt"

	"github.com/labstack/echo"
	. "github.com/onsi/gomega"
)

type Person struct {
	Name string
	Age  int
}

var (
	person = Person{
		Name: "joao",
		Age:  30,
	}
)

func TestSomething(t *testing.T) {
	RegisterTestingT(t)

	e := echo.New()

	tests := []struct {
		Name            string
		Method          string
		UrlValues       map[string][]string
		RequestBody     string
		ResponseCode    int
		ResponseBody    string
		HandlerFunction func(echo.Context) error
	}{
		{
			Name:            "Valid json request",
			UrlValues:       map[string][]string{"name": []string{"joao"}},
			Method:          "POST",
			RequestBody:     `{"cenas":"worked"}`,
			ResponseCode:    http.StatusOK,
			ResponseBody:    `{"name":"joao","age":30}`,
			HandlerFunction: domain.SayHello,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.Method, "/", strings.NewReader(test.RequestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Form = test.UrlValues

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		for key, value := range test.UrlValues {
			c.SetParamNames(key)
			c.SetParamValues(value[0])
		}

		err := test.HandlerFunction(c)
		Expect(err).To(BeNil(), "Should not return error")
		Expect(rec.Body.String()).To(Equal(test.ResponseBody), fmt.Sprintf("Should return the correct body {obtained: %s, expected: %s}", rec.Body.String(), test.ResponseBody))
		Expect(rec.Code).To(Equal(test.ResponseCode), fmt.Sprintf("Should return the correct code {obtained: %d, expected: %d}", rec.Code, test.ResponseCode))
	}
}
