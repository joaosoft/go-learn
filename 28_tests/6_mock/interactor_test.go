package example

import (
	"errors"
	domain "github.com/joaosoft/golang-learn/28_tests/6_mock/domain"
	"github.com/joaosoft/golang-learn/28_tests/6_mock/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Person struct {
	Name string
	Age  int
}

type setup struct {
	interactor *domain.Interactor
	repository *mocks.RepositoryMock
	person     Person
}

func setupMocks() *setup {
	repository := &mocks.RepositoryMock{}
	interactor := domain.NewInteractor(repository)
	person := Person{
		Name: "joao",
		Age:  30,
	}

	return &setup{
		interactor: interactor,
		repository: repository,
		person:     person,
	}
}

var _ = Describe("Tests with mock", func() {

	Describe("when processing DoSomethingInsert method", func() {

		Context("when method returns success", func() {
			test := setupMocks()
			test.repository.On("Store", "joao", 30).Return(nil)

			err := test.interactor.DoSomething(test.person.Name, test.person.Age)

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when method returns error", func() {
			test := setupMocks()
			test.repository.On("Store", "joao", 30).Return(errors.New("oops"))

			err := test.interactor.DoSomething(test.person.Name, test.person.Age)

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).NotTo(Equal(BeEmpty()))
			})
		})
	})
})
