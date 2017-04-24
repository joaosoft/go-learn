package dummy_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
)

type dataTest struct {
	testeStruct *TesteStruct
	repository *mocks.RepositoryMock
}

func setupMocks() *dataTest {
	repository := new(mocks.RepositoryMock)

	testeStruct := testeStruct.NewTesteStruct(repository)

	return &dataTest{
		testStruct: testStruct,
		repository: repository}
}

var _ = Describe("testStruct", func() {

	Describe("when processing DoSomething method", func() {
		id := "123"
		value := "456"

		Context("when method returns succcess", func() {
			test := setupMocks()
			test.repository.On("UpdateRecordState", importID, stateInt).Return(nil)

			err := test.interactor.HandleMessage(&msg)

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when HandleMessage returns error", func() {
			test := setupMocks()
			test.repository.On("UpdateRecordState", importID, stateInt).Return(fmt.Errorf(errorMsg))

			err := test.interactor.HandleMessage(&msg)

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(errorMsg))
			})
		})
	})
})
