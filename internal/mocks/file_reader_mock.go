package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zniptr/flowcraft/internal/helpers"
)

type ChartFileReaderMock struct {
	mock.Mock
}

func NewChartFileReaderMock() *ChartFileReaderMock {
	return &ChartFileReaderMock{}
}

func (mock *ChartFileReaderMock) ReadDirectory(path string) ([]helpers.DirEntryHelper, error) {
	args := mock.Called(path)

	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]helpers.DirEntryHelper), args.Error(1)
}

func (mock *ChartFileReaderMock) ReadFile(path string, file helpers.DirEntryHelper) ([]byte, error) {
	args := mock.Called(path, file)

	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]byte), args.Error(1)
}

func (mock *ChartFileReaderMock) IsValidChartFile(file helpers.DirEntryHelper) bool {
	args := mock.Called(file)

	return args.Bool(0)
}
