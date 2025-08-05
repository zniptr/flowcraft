package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MutexHelperMock struct {
	mock.Mock
}

func NewMutexHelperMock() *MutexHelperMock {
	return &MutexHelperMock{}
}

func (mock *MutexHelperMock) Lock() {
	_ = mock.Called()
}

func (mock *MutexHelperMock) Unlock() {
	_ = mock.Called()
}
