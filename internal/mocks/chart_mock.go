package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zniptr/flowcraft/internal/file"
)

type ChartMock struct {
	mock.Mock
}

func NewChartMock() *ChartMock {
	return &ChartMock{}
}

func (mock *ChartMock) GetName() string {
	args := mock.Called()

	return args.String(0)
}

func (mock *ChartMock) GetStart() *file.Object {
	args := mock.Called()

	result := args.Get(0)
	if result == nil {
		return nil
	}

	return args.Get(0).(*file.Object)
}

func (mock *ChartMock) GetObject(id string) *file.Object {
	args := mock.Called(id)

	result := args.Get(0)
	if result == nil {
		return nil
	}

	return args.Get(0).(*file.Object)
}

func (mock *ChartMock) GetOutgoingConnection(id string) *file.Object {
	args := mock.Called(id)

	result := args.Get(0)
	if result == nil {
		return nil
	}

	return args.Get(0).(*file.Object)
}

func (mock *ChartMock) GetOutgoingNonDefaultConnections(id string) []*file.Object {
	args := mock.Called(id)

	return args.Get(0).([]*file.Object)
}

func (mock *ChartMock) GetOutgoingDefaultConnection(id string) *file.Object {
	args := mock.Called(id)

	result := args.Get(0)
	if result == nil {
		return nil
	}

	return args.Get(0).(*file.Object)
}
