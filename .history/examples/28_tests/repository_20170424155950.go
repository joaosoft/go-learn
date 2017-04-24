package repositories

import (
	"encoding/json"
	"fmt"
	"productoffers/common/session"
	"productoffers/services/mirakl-mci/vo"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Repository defines the repository for Mirakl's MCI integration.
type Repository struct {
	conn session.IDBContext
	name string
}

// NewRepository creates a new Repository.
func NewRepository(dbcontext session.IDBContext) *Repository {
	repository := new(Repository)
	repository.conn = dbcontext
	repository.name = "mirakl-mci"

	return repository
}

// SetDBContext sets the db context allowing to replace the repository's connection with another connection or transaction.
func (repository *Repository) SetDBContext(ctx session.IDBContext) {
	repository.conn = ctx
}

// StoreImport stores an MCI import.
func (repository *Repository) StoreImport(importID string, account string) error {
	if _, err := repository.conn.Exec(`
		INSERT INTO imports(import_id, account)
			VALUES($1, $2)
	`,
		importID,
		account,
	); err != nil {
		return err
	}

	return nil
}
