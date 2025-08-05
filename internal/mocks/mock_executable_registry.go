package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zniptr/flowcraft/pkg/executable"
)

type ExecutableRegistryMock struct {
	mock.Mock
}

func NewExecutableRegistryMock() *ExecutableRegistryMock {
	return &ExecutableRegistryMock{}
}

func (mock *ExecutableRegistryMock) Get(name string) executable.ExecutableFactory {
	args := mock.Called(name)

	result := args.Get(0)
	if result == nil {
		return nil
	}

	return args.Get(0).(executable.ExecutableFactory)
}

func (mock *ExecutableRegistryMock) Register(name string, factory executable.ExecutableFactory) {
	_ = mock.Called(name)
}

func (mock *ExecutableRegistryMock) Unregister(name string) {
	_ = mock.Called(name)
}
