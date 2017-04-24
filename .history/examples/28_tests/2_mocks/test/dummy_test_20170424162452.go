package dummy_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
)

type dataTest struct {
	controller *TesteStruct
	repository *RepositoryMock
}

func setupMocks() *dataTest {
	repository := new(RepositoryMock)

	controller := testeStruct.NewTesteStruct(repository)

	return &dataTest{
		controller: testStruct,
		repository: repository}
}

var _ = Describe("testStruct", func() {

	Describe("when processing DoSomething method", func() {
		id := "123"
		value := "456"

		Context("when method returns succcess", func() {
			test := setupMocks()
			test.repository.On("SomeMethod", id, value).Return(nil)

			err := test.interactor.DoSomething(id, value)

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when method returns error", func() {
			test := setupMocks()
			test.repository.On("SomeMethod", id, value).Return(nil)

			err := test.interactor.DoSomething(id, value)

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(errorMsg))
			})
		})
	})
})
