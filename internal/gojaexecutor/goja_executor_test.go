package gojaexecutor

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/helpers"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type GojaExecutorTestSuite struct {
	suite.Suite
	mockError     error
	mockKey       string
	mockValue     string
	mockVariables map[string]any
	mockCode      string

	mockGoja      *mocks.GojaHelperMock
	mockGojaValue *mocks.GojaValueHelperMock
}

func (suite *GojaExecutorTestSuite) SetupTest() {
	suite.mockError = errors.New("A general error has occurred")
	suite.mockKey = "testKey"
	suite.mockValue = "testValue"
	suite.mockVariables = map[string]any{suite.mockKey: suite.mockValue}
	suite.mockCode = "testCode"

	suite.mockGoja = mocks.NewGojaHelperMock()
	suite.mockGojaValue = mocks.NewGojaValueHelperMock()

	newGojaFunc = suite.newGojaMock
}

func (suite *GojaExecutorTestSuite) newGojaMock() helpers.GojaHelper {
	return suite.mockGoja
}

func (suite *GojaExecutorTestSuite) TestGojaExecutor_whenCreateGojaExecutor_thenReturnGojaExecutor() {
	gojaExecutor := NewGojaExecutor()

	suite.NotNil(gojaExecutor)
	suite.IsType(&GojaExecutorImpl{}, gojaExecutor)
}

func (suite *GojaExecutorTestSuite) TestSetVariables_whenErrorOnSet_thenReturnError() {
	suite.mockGoja.On("Set", suite.mockKey, suite.mockValue).Return(suite.mockError)

	err := NewGojaExecutor().SetVariables(suite.mockVariables)

	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *GojaExecutorTestSuite) TestSetVariables_whenSet_thenSetVariables() {
	suite.mockGoja.On("Set", suite.mockKey, suite.mockValue).Return(nil)

	err := NewGojaExecutor().SetVariables(suite.mockVariables)

	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *GojaExecutorTestSuite) TestRun_whenErrorOnRunString_thenReturnError() {
	suite.mockGoja.On("RunString", suite.mockCode).Return(nil, suite.mockError)

	value, err := NewGojaExecutor().Run(suite.mockCode)

	suite.Nil(value)
	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *GojaExecutorTestSuite) TestRun_whenRunString_thenReturnValue() {
	suite.mockGojaValue.On("Export").Return(suite.mockValue)
	suite.mockGoja.On("RunString", suite.mockCode).Return(suite.mockGojaValue, nil)

	value, err := NewGojaExecutor().Run(suite.mockCode)

	suite.Equal(suite.mockValue, value)
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestGojaExecutorTestSuite(t *testing.T) {
	suite.Run(t, new(GojaExecutorTestSuite))
}
