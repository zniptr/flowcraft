package executableregistry

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/helpers"
	"github.com/zniptr/flowcraft/internal/mocks"
	"github.com/zniptr/flowcraft/pkg/executable"
)

type ExecutableRegistryTestSuite struct {
	suite.Suite
	mockExecutableName string

	mockMutex             *mocks.MutexHelperMock
	mockExecutableFactory executable.ExecutableFactory
}

func (suite *ExecutableRegistryTestSuite) SetupTest() {
	suite.mockExecutableName = "testExecutable"

	suite.mockMutex = mocks.NewMutexHelperMock()
	suite.mockExecutableFactory = func() executable.Executable { return nil }

	getMutexFunc = suite.getMutexMock
}

func (suite *ExecutableRegistryTestSuite) AfterTest(suiteName, methodName string) {
	instance = nil
}

func (suite *ExecutableRegistryTestSuite) getMutexMock() helpers.MutexHelper {
	return suite.mockMutex
}

func (suite *ExecutableRegistryTestSuite) TestGetInstance_whenCallGetInstance_thenReturnSameInstance() {
	suite.mockMutex.On("Lock")
	suite.mockMutex.On("Unlock")

	instance1 := GetInstance()
	instance2 := GetInstance()

	suite.Equal(instance1, instance2)
	suite.mockMutex.AssertNumberOfCalls(suite.T(), "Lock", 1)
	suite.mockMutex.AssertNumberOfCalls(suite.T(), "Unlock", 1)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ExecutableRegistryTestSuite) TestGet_whenGetExecutable_thenReturnExecutable() {
	suite.mockMutex.On("Lock")
	suite.mockMutex.On("Unlock")

	registry := GetInstance()
	registry.Register(suite.mockExecutableName, suite.mockExecutableFactory)

	factory := registry.Get(suite.mockExecutableName)

	suite.Equal(suite.mockExecutableFactory(), factory())
	suite.mockMutex.AssertNumberOfCalls(suite.T(), "Lock", 3)
	suite.mockMutex.AssertNumberOfCalls(suite.T(), "Unlock", 3)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ExecutableRegistryTestSuite) TestRegister_whenRegisterExecutable_thenAddExecutableToRegistry() {
	suite.mockMutex.On("Lock")
	suite.mockMutex.On("Unlock")

	registry := GetInstance()
	registry.Register(suite.mockExecutableName, suite.mockExecutableFactory)

	suite.mockMutex.AssertNumberOfCalls(suite.T(), "Lock", 2)
	suite.mockMutex.AssertNumberOfCalls(suite.T(), "Unlock", 2)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ExecutableRegistryTestSuite) TestUnregister_whenUnregisterExecutable_thenRemoveExecutableFromRegistry() {
	suite.mockMutex.On("Lock")
	suite.mockMutex.On("Unlock")

	registry := GetInstance()
	registry.Register(suite.mockExecutableName, suite.mockExecutableFactory)
	registry.Unregister(suite.mockExecutableName)

	factory := registry.Get(suite.mockExecutableName)

	suite.Nil(factory)
	suite.mockMutex.AssertNumberOfCalls(suite.T(), "Lock", 4)
	suite.mockMutex.AssertNumberOfCalls(suite.T(), "Unlock", 4)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestExecutableRegistryTestSuite(t *testing.T) {
	suite.Run(t, new(ExecutableRegistryTestSuite))
}
