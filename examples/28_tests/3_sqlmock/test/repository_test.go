package dummy_test

import (
	"encoding/json"
	"fmt"
		xpto "learn_go/examples/28_tests/3_sqlmock/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var _ = Describe("Repository", func() {
	id := "1"
	value := "teste"

	Describe("when calling DoSomething()", func() {
		query := "INSERT INTO dummy"

		Context("when success occurs", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := xpto.NewRepository(db)

			mock.ExpectExec(query).WithArgs(id, value).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repository.DoSomething(id, value)
			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
