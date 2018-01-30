package dummy

import (
	domain "golang-learn/28_tests/2_mock/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type dataTest struct {
	interactor *domain.Interactor
	repository *RepositoryMock
}

func setupMocks() *dataTest {
	repository := new(RepositoryMock)
	interactor := domain.Interactor(repository)

	return &dataTest{
		interactor: &interactor,
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
				Expect(err.Error()).NotTo(Equal(BeEmpty()))
			})
		})
	})
})
