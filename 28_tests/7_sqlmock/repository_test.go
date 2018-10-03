package example

import (
	"fmt"
	"go-learn/28_tests/7_sqlmock/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Person struct {
	Name string
	Age  int
}

var _ = Describe("Repository", func() {
	person := Person{
		Name: "joao",
		Age:  30,
	}

	Describe("when calling DoSomething()", func() {
		query := "INSERT INTO something"

		Context("when success occurs", func() {
			db, mock, _ := sqlmock.New()

			defer db.Close()
			repository := domain.NewRepository(db)

			mock.ExpectBegin()
			mock.ExpectExec(query).WithArgs("joao", 30).WillReturnResult(sqlmock.NewResult(2, 1))
			mock.ExpectCommit()

			id, rows, err := repository.DoSomething(person.Name, person.Age)
			It("should not return an error", func() {
				Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(int64(2)))
				Expect(rows).To(Equal(int64(1)))
			})
		})

		Context("when error occurs", func() {
			db, mock, _ := sqlmock.New()

			defer db.Close()
			repository := domain.NewRepository(db)

			mock.ExpectBegin()
			mock.ExpectExec(query).WithArgs("joao", 30).WillReturnError(fmt.Errorf("oops"))
			mock.ExpectRollback()

			id, rows, err := repository.DoSomething(person.Name, person.Age)
			It("should return an error", func() {
				Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("oops"))
				Expect(id).To(Equal(int64(0)))
				Expect(rows).To(Equal(int64(0)))
			})
		})
	})
})
