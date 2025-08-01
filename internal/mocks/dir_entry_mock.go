package mocks

import "github.com/stretchr/testify/mock"

type DirEntryMock struct {
	mock.Mock
}

func NewDirEntryMock() *DirEntryMock {
	return &DirEntryMock{}
}

func (mock *DirEntryMock) IsDir() bool {
	args := mock.Called()

	return args.Bool(0)
}

func (mock *DirEntryMock) Name() string {
	args := mock.Called()

	return args.String(0)
}
