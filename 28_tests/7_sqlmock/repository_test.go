package example

import (
	"errors"
	"fmt"
	"go-learn/28_tests/7_sqlmock/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var _ = Describe("Repository", func() {

	Describe("when calling DoSomethingSelect to SELECT()", func() {
		query := "SELECT name, age FROM something WHERE name = ?"
		person1 := domain.Person{
			Name: "joao",
			Age:  30,
		}
		person2 := domain.Person{
			Name: "joao",
			Age:  40,
		}

		rows := sqlmock.NewRows([]string{"name", "age"}).
			AddRow(person1.Name, person1.Age).
			AddRow(person2.Name, person2.Age)

		Context("when success occurs", func() {
			db, mock, _ := sqlmock.New()

			defer db.Close()
			repository := domain.NewRepository(db)

			mock.ExpectQuery(query).WithArgs("joao").WillReturnRows(rows)

			persons, err := repository.DoSomethingSelect("joao")
			It("should not return an error", func() {
				Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				Expect(err).ToNot(HaveOccurred())

				Expect(len(persons)).To(Equal(2))

				Expect(*persons[0]).To(Equal(person1))
				Expect(*persons[1]).To(Equal(person2))
			})
		})

		Context("when error occurs", func() {
			db, mock, _ := sqlmock.New()

			defer db.Close()
			repository := domain.NewRepository(db)

			mock.ExpectQuery(query).WithArgs("joao").WillReturnError(errors.New("oops"))

			persons, err := repository.DoSomethingSelect("joao")
			It("should return an error", func() {
				Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("oops"))
				Expect(persons).To(BeNil())
			})
		})
	})

	Describe("when calling DoSomethingInsert to INSERT()", func() {
		query := "INSERT INTO something \\(name, age\\) VALUES "
		person := domain.Person{
			Name: "joao",
			Age:  30,
		}

		Context("when success occurs", func() {
			db, mock, _ := sqlmock.New()

			defer db.Close()
			repository := domain.NewRepository(db)

			mock.ExpectBegin()
			mock.ExpectExec(query).WithArgs("joao", 30).WillReturnResult(sqlmock.NewResult(2, 1))
			mock.ExpectCommit()

			id, rows, err := repository.DoSomethingInsert(person.Name, person.Age)
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

			id, rows, err := repository.DoSomethingInsert(person.Name, person.Age)
			It("should return an error", func() {
				Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("oops"))
				Expect(id).To(Equal(int64(0)))
				Expect(rows).To(Equal(int64(0)))
			})
		})
	})

	Describe("when calling DoSomethingUpdate to UPDATE()", func() {
		query := "UPDATE something SET age = ?"
		person := domain.Person{
			Name: "joao",
			Age:  30,
		}

		Context("when success occurs", func() {
			db, mock, _ := sqlmock.New()

			defer db.Close()
			repository := domain.NewRepository(db)

			mock.ExpectBegin()
			mock.ExpectExec(query).WithArgs(30, "joao").WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			updated, err := repository.DoSomethingUpdate(person.Name, person.Age)
			It("should not return an error", func() {
				Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				Expect(err).ToNot(HaveOccurred())
				Expect(updated).To(BeTrue())
			})
		})

		Context("when error occurs", func() {
			db, mock, _ := sqlmock.New()

			defer db.Close()
			repository := domain.NewRepository(db)

			mock.ExpectBegin()
			mock.ExpectExec(query).WithArgs(30, "joao").WillReturnError(fmt.Errorf("oops"))
			mock.ExpectRollback()

			updated, err := repository.DoSomethingUpdate(person.Name, person.Age)
			It("should return an error", func() {
				Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("oops"))
				Expect(updated).To(BeFalse())
			})
		})
	})
})
