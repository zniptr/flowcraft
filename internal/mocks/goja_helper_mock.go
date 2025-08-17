package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zniptr/flowcraft/internal/helpers"
)

type GojaHelperMock struct {
	mock.Mock
}

func NewGojaHelperMock() *GojaHelperMock {
	return &GojaHelperMock{}
}

func (mock *GojaHelperMock) Set(name string, value interface{}) error {
	args := mock.Called(name, value)
	return args.Error(0)
}

func (mock *GojaHelperMock) RunString(str string) (helpers.GojaValueHelper, error) {
	args := mock.Called(str)

	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(helpers.GojaValueHelper), args.Error(1)
}
