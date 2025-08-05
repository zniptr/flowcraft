package mocks

import "github.com/stretchr/testify/mock"

type ChartContextMock struct {
	mock.Mock
}

func NewChartContextMock() *ChartContextMock {
	return &ChartContextMock{}
}
