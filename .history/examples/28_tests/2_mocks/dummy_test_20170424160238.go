package dummy_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
)

type dataListenerTest struct {
	interactor *interactors.ListenerInteractor
	repository *mocks.RepositoryMock
}

func setupListenerMocks() *dataListenerTest {
	repository := new(mocks.MciRepositoryMock)

	interactor := interactors.NewMCIListenerInteractor(repository)

	return &dataListenerTest{
		interactor: interactor,
		repository: repository}
}

var _ = Describe("Listener Interactor", func() {

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
			test := setupListenerMocks()
			test.repository.On("UpdateRecordState", importID, stateInt).Return(nil)

			err := test.interactor.HandleMessage(&msg)

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when HandleMessage returns error", func() {
			test := setupListenerMocks()
			test.repository.On("UpdateRecordState", importID, stateInt).Return(fmt.Errorf(errorMsg))

			err := test.interactor.HandleMessage(&msg)

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(errorMsg))
			})
		})
	})
})
