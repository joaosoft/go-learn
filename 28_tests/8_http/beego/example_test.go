package echo

import (
	"go-learn/28_tests/8_http/beego/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/astaxie/beego/context"

	"fmt"

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

	controller := domain.Controller{}
	tests := []struct {
		Name            string
		Method          string
		UrlValues       map[string][]string
		RequestBody     string
		ResponseCode    int
		ResponseBody    string
		HandlerFunction func()
	}{
		{
			Name:            "Valid json request",
			UrlValues:       map[string][]string{"name": []string{"joao"}},
			Method:          "POST",
			RequestBody:     `{"cenas":"worked"}`,
			ResponseCode:    http.StatusOK,
			ResponseBody:    `{"name":"joao","age":30}`,
			HandlerFunction: controller.SayHello,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.Method, "/hello/joao", strings.NewReader(test.RequestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		context := context.Context{
			Input: &context.BeegoInput{
				Context: &context.Context{
					Request:        req,
					ResponseWriter: &context.Response{ResponseWriter: rec},
				},
			},
			Output: &context.BeegoOutput{
				Context: &context.Context{
					Request:        req,
					ResponseWriter: &context.Response{ResponseWriter: rec},
				},
			},
			Request: req,
			ResponseWriter: &context.Response{
				ResponseWriter: rec,
			},
		}

		controller.Init(&context, "", "", "")

		for key, value := range test.UrlValues {
			controller.Ctx.Input.SetParam(key, value[0])
		}

		test.HandlerFunction()

		Expect(rec.Body.String()).To(Equal(test.ResponseBody), fmt.Sprintf("Should return the correct body {obtained: %s, expected: %s}", rec.Body.String(), test.ResponseBody))
		Expect(rec.Code).To(Equal(test.ResponseCode), fmt.Sprintf("Should return the correct code {obtained: %d, expected: %d}", rec.Code, test.ResponseCode))
	}
}
