package mocks

import (
	"github.com/stretchr/testify/mock"
)

type GojaValueHelperMock struct {
	mock.Mock
}

func NewGojaValueHelperMock() *GojaValueHelperMock {
	return &GojaValueHelperMock{}
}

func (mock *GojaValueHelperMock) Export() interface{} {
	args := mock.Called()
	return args.Get(0)
}
