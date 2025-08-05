package mocks

import "github.com/stretchr/testify/mock"

type ChartInstanceMock struct {
	mock.Mock
}

func NewChartInstanceMock() *ChartInstanceMock {
	return &ChartInstanceMock{}
}

func (mock *ChartInstanceMock) Run() error {
	args := mock.Called()

	return args.Error(0)
}
