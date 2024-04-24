package echo

import (
	"github.com/joaosoft/golang-learn/28_tests/8_http/http/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"fmt"

	echo "github.com/labstack/echo/v4"
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

	tests := []struct {
		Name            string
		Method          string
		UrlValues       map[string][]string
		RequestBody     string
		ResponseCode    int
		ResponseBody    string
		HandlerFunction func(w http.ResponseWriter, r *http.Request)
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

		for key, value := range test.UrlValues {
			req.Form[key] = value
		}

		test.HandlerFunction(rec, req)
		Expect(rec.Body.String()).To(Equal(test.ResponseBody), fmt.Sprintf("Should return the correct body {obtained: %s, expected: %s}", rec.Body.String(), test.ResponseBody))
		Expect(rec.Code).To(Equal(test.ResponseCode), fmt.Sprintf("Should return the correct code {obtained: %d, expected: %d}", rec.Code, test.ResponseCode))
	}
}
