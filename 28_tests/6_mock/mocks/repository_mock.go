package mocks

import (
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (mock *RepositoryMock) Store(name string, age int) error {
	args := mock.Called(name, age)
	return args.Error(0)
}
