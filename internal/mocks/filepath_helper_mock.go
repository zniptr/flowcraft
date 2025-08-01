package mocks

import "github.com/stretchr/testify/mock"

type FilepathHelperMock struct {
	mock.Mock
}

func NewFilepathHelperMock() *FilepathHelperMock {
	return &FilepathHelperMock{}
}

func (mock *FilepathHelperMock) Join(elem ...string) string {
	args := mock.Called(elem)
	return args.String(0)
}

func (mock *FilepathHelperMock) Ext(path string) string {
	args := mock.Called(path)
	return args.String(0)
}
