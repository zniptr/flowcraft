package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zniptr/flowcraft/internal/helpers"
)

type DirEntryHelperMock struct {
	mock.Mock
}

func NewDirEntryHelperMock() *DirEntryHelperMock {
	return &DirEntryHelperMock{}
}

func (mock *DirEntryHelperMock) IsDir() bool {
	args := mock.Called()
	return args.Bool(0)
}

func (mock *DirEntryHelperMock) Name() string {
	args := mock.Called()
	return args.String(0)
}

type OsHelperMock struct {
	mock.Mock
}

func NewOsHelperMock() *OsHelperMock {
	return &OsHelperMock{}
}

func (mock *OsHelperMock) ReadDir(name string) ([]helpers.DirEntryHelper, error) {
	args := mock.Called(name)

	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]helpers.DirEntryHelper), args.Error(1)
}

func (mock *OsHelperMock) ReadFile(name string) ([]byte, error) {
	args := mock.Called(name)

	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]byte), args.Error(1)
}
