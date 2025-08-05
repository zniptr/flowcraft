package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type ExecutableMock struct {
	mock.Mock
}

func NewExecutableMock() *ExecutableMock {
	return &ExecutableMock{}
}

func (mock *ExecutableMock) Execute(chartContext chartcontext.ChartContext) error {
	args := mock.Called(chartContext)

	return args.Error(0)
}

type ExecutableFactoryMock func() ExecutableMock
