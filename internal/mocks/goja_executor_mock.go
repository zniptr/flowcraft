package mocks

import "github.com/stretchr/testify/mock"

type GojaExecutorMock struct {
	mock.Mock
}

func NewGojaExecutorMock() *GojaExecutorMock {
	return &GojaExecutorMock{}
}

func (mock *GojaExecutorMock) SetVariables(vars map[string]any) error {
	args := mock.Called(vars)
	return args.Error(0)
}

func (mock *GojaExecutorMock) Run(code string) (interface{}, error) {
	args := mock.Called(code)
	return args.Get(0), args.Error(1)
}
