package mocks

import "github.com/stretchr/testify/mock"

type ChartContextMock struct {
	mock.Mock
}

func NewChartContextMock() *ChartContextMock {
	return &ChartContextMock{}
}

func (mock *ChartContextMock) GetContext() map[string]any {
	args := mock.Called()

	return args.Get(0).(map[string]any)
}

func (mock *ChartContextMock) GetVariable(name string) any {
	args := mock.Called(name)

	return args.Get(0)
}

func (mock *ChartContextMock) SetVariable(name string, value any) {
	_ = mock.Called(name, value)
}
