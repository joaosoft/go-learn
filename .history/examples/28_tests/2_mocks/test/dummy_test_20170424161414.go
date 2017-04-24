package dummy_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
)

type dataTest struct {
	testeStruct *domain.dummy.TesteStruct
	repository *mocks.RepositoryMock
}

func setupMocks() *dataTest {
	repository := new(mocks.RepositoryMock)

	testeStruct := testeStruct.NewTesteStruct(repository)

	return &dataTest{
		interactor: interactor,
		repository: repository}
}

var _ = Describe(" Interactor", func() {

	Describe("when processing HandleMessage method", func() {
		productID := uuid.NewV4()
		importID := uuid.NewV4()
		sourceID := uuid.NewV4()
		sourceProductCode := "123"
		stateStr := "pending"
		stateInt := vo.States[stateStr]
		errorMsg := "oops"

		msg := vo.ReportMessage{
			ProductID:         productID,
			ImportID:          importID,
			SourceID:          sourceID,
			SourceProductCode: sourceProductCode,
			State:             stateStr}

		Context("when HandleMessage returns succcess", func() {
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
