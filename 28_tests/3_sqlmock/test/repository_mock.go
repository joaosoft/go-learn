package dummytest

import (
	"github.com/stretchr/testify/mock"
)

// RepositoryMock ...
type RepositoryMock struct {
	mock.Mock
}

// StoreImport ...
func (mock *RepositoryMock) StoreImport(id string, value string) error {
	args := mock.Called(id, value)
	return args.Error(0)
}
