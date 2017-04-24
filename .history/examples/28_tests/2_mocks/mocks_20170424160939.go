package mocks

import (
	"bytes"
	"mime/multipart"
	session "productoffers/common/session"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
)

// RepositoryMock ...
type RepositoryMock struct {
	mock.Mock
}

// SetDBContext ...
func (mock *RepositoryMock) SetDBContext(dbc session.IDBContext) {}

// StoreImport ...
func (mock *RepositoryMock) StoreImport(id string, value string) error {
	args := mock.Called(id, value)
	return args.Error(0)
}
