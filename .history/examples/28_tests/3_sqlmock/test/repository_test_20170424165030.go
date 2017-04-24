package repositories_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var _ = Describe("MCI Repository", func() {

	Describe("when calling DoSomething()", func() {
		query := "INSERT INTO mirakl_mci_imports"

		account := "123"
		importID := uuid.NewV4()
		miraklImportID := "1"
		shopID := "2"
		rawData := "3"
		reported := true

		importData := vo.MCIImport{
			ImportID:       importID,
			MiraklImportID: miraklImportID,
			ShopID:         shopID,
			RawData:        rawData,
			Reported:       reported}

		Context("when success occurs", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(importID, miraklImportID, shopID, rawData).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repository.StoreImport(account, &importData)
			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when an error occurs", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(importID, miraklImportID, shopID, rawData).WillReturnError(fmt.Errorf("oops"))

			err := repository.StoreImport(account, &importData)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("when calling SetImportProcessed()", func() {
		query := "UPDATE mirakl_mci_imports SET processed = 't' WHERE import_id = ?"

		importID := uuid.NewV4()

		Context("when success occurs", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(importID).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repository.SetImportProcessed(importID)
			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when error occurs with no results", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(importID).WillReturnResult(sqlmock.NewResult(0, 0))

			err := repository.SetImportProcessed(importID)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when an error occurs", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(importID).WillReturnError(fmt.Errorf("oops"))

			err := repository.SetImportProcessed(importID)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("when calling StoreRecords()", func() {
		query := "INSERT INTO mirakl_mci_import_records"

		recordID := uuid.NewV4()
		importID := uuid.NewV4()
		data := []string{"1", "2", "3"}
		dataMap := map[string]string{
			"A": "111",
			"B": "222",
			"C": "333"}
		notified := true
		stat := vo.States["pending"]

		importRecordData := vo.MCIImportRecord{
			RecordID: recordID,
			ImportID: importID,
			Data:     data,
			DataMap:  dataMap,
			Notified: notified,
			State:    uint8(stat)}
		records := []*vo.MCIImportRecord{
			&importRecordData,
			&importRecordData}
		raw, _ := json.Marshal(importRecordData.DataMap)

		// Context("when success occurs", func() {
		// 	db, mock, _ := sqlmock.New()
		// 	defer db.Close()
		// 	repository := repositories.NewMCIRepository(db)

		// 	mock.ExpectPrepare(query).ExpectExec().WithArgs(recordID, importID, raw)
		// 	mock.ExpectExec(query).WithArgs(recordID, importID, raw).WillReturnResult(sqlmock.NewResult(1, 1))

		// 	err := repository.StoreRecords(records)
		// 	It("should not return an error", func() {
		// 		Expect(err).ToNot(HaveOccurred())
		// 	})
		// })

		Context("when error occurs with no results", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectPrepare(query).ExpectExec().WithArgs(recordID, importID, raw)
			mock.ExpectExec(query).WithArgs(recordID, importID, raw).WillReturnResult(sqlmock.NewResult(0, 0))

			err := repository.StoreRecords(records)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("when calling StoreRecord()", func() {
		query := "INSERT INTO mirakl_mci_import_records"

		recordID := uuid.NewV4()
		importID := uuid.NewV4()
		data := []string{"1", "2", "3"}
		dataMap := map[string]string{
			"A": "111",
			"B": "222",
			"C": "333"}
		notified := true
		stat := vo.States["pending"]

		importRecordData := vo.MCIImportRecord{
			RecordID: recordID,
			ImportID: importID,
			Data:     data,
			DataMap:  dataMap,
			Notified: notified,
			State:    uint8(stat)}

		raw, _ := json.Marshal(importRecordData.DataMap)

		Context("when success occurs", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(recordID, importID, raw).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repository.StoreRecord(&importRecordData)
			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when error occurs with no results", func() {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(recordID, importID, raw).WillReturnError(fmt.Errorf("oops"))

			err := repository.StoreRecord(&importRecordData)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("when calling UpdateRecordState()", func() {
		query := "UPDATE mirakl_mci_import_records SET state = ? WHERE record_id = ?"

		recordID := uuid.NewV4()

		Context("when success occurs", func() {
			stat := vo.States["invalid"]
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(recordID, stat).WillReturnResult(sqlmock.NewResult(1, 1))

			err := repository.UpdateRecordState(recordID, stat)
			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when error occurs with no results", func() {
			stat := vo.States["invalid"]
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(recordID, stat).WillReturnResult(sqlmock.NewResult(0, 0))

			err := repository.UpdateRecordState(recordID, stat)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when an error occurs", func() {
			stat := vo.States["invalid"]
			db, mock, _ := sqlmock.New()
			defer db.Close()
			repository := repositories.NewMCIRepository(db)

			mock.ExpectExec(query).WithArgs(recordID).WillReturnError(fmt.Errorf("oops"))

			err := repository.UpdateRecordState(recordID, stat)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
