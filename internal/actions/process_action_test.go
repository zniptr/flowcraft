package actions

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/mocks"
	"github.com/zniptr/flowcraft/pkg/executable"
	"github.com/zniptr/flowcraft/pkg/executableregistry"
)

type ProcessActionTestSuite struct {
	suite.Suite
	mockError           error
	mockExecutableName  string
	mockId              string
	mockLabel           string
	mockTargetId        string
	mockConnectionLabel string

	mockExecutableRegistry *mocks.ExecutableRegistryMock
	mockExecutable         *mocks.ExecutableMock
	mockChartContext       *mocks.ChartContextMock
	mockChart              *mocks.ChartMock
	mockObject             *file.Object
	mockConnectionObject   *file.Object
	mockNextObject         *file.Object
	mockExecutableFactory  executable.ExecutableFactory
}

func (suite *ProcessActionTestSuite) SetupTest() {
	suite.mockError = errors.New("A general error has occurred")
	suite.mockExecutableName = "testExecutable"
	suite.mockId = "testId"
	suite.mockLabel = "testLabel"
	suite.mockTargetId = "testTargetId"
	suite.mockConnectionLabel = "testConnectionLabel"

	suite.mockExecutableRegistry = mocks.NewExecutableRegistryMock()
	suite.mockExecutable = mocks.NewExecutableMock()
	suite.mockChartContext = mocks.NewChartContextMock()
	suite.mockChart = mocks.NewChartMock()
	suite.mockObject = &file.Object{Id: suite.mockId, Label: suite.mockLabel, Executable: suite.mockExecutableName}
	suite.mockConnectionObject = &file.Object{Label: suite.mockConnectionLabel, Cell: file.Cell{Target: suite.mockTargetId}}
	suite.mockNextObject = &file.Object{}
	suite.mockExecutableFactory = func() executable.Executable { return suite.mockExecutable }

	getExecutableRegistryInstanceFunc = suite.getExecutableRegistryMock
}

func (suite *ProcessActionTestSuite) getExecutableRegistryMock() executableregistry.ExecutableRegistry {
	return suite.mockExecutableRegistry
}

func (suite *ProcessActionTestSuite) TestNewProcessAction_whenNewProcessAction_thenReturnProcessAction() {
	processAction := NewProcessAction()

	suite.NotNil(processAction)
	suite.IsType(&ProcessAction{}, processAction)
}

func (suite *ProcessActionTestSuite) TestProcessActionExecute_whenErrorOnExecute_thenReturnError() {
	suite.mockExecutableRegistry.On("Get", suite.mockExecutableName).Return(suite.mockExecutableFactory)
	suite.mockExecutable.On("Execute", suite.mockChartContext).Return(suite.mockError)

	err := NewProcessAction().Execute(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ProcessActionTestSuite) TestProcessActionExecute_whenExecute_thenReturnNil() {
	suite.mockExecutableRegistry.On("Get", suite.mockExecutableName).Return(suite.mockExecutableFactory)
	suite.mockExecutable.On("Execute", suite.mockChartContext).Return(nil)

	err := NewProcessAction().Execute(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ProcessActionTestSuite) TestProcessActionGetNext_whenNoConnection_thenReturnError() {
	suite.mockChart.On("GetSingleConnectionBySourceId", suite.mockId).Return(nil)

	next, err := NewProcessAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("no source connection for process action %s", suite.mockObject.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ProcessActionTestSuite) TestProcessActionGetNext_whenConnectionHasNoTarget_thenReturnError() {
	suite.mockChart.On("GetSingleConnectionBySourceId", suite.mockId).Return(suite.mockConnectionObject)
	suite.mockChart.On("GetObjectById", suite.mockTargetId).Return(nil)

	next, err := NewProcessAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("no target for connection action %s", suite.mockConnectionObject.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ProcessActionTestSuite) TestProcessActionGetNext_whenNextFound_thenReturnNext() {
	suite.mockChart.On("GetSingleConnectionBySourceId", suite.mockId).Return(suite.mockConnectionObject)
	suite.mockChart.On("GetObjectById", suite.mockTargetId).Return(suite.mockNextObject)

	next, err := NewProcessAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Equal(suite.mockNextObject, next)
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestProcessActionTestSuite(t *testing.T) {
	suite.Run(t, new(ProcessActionTestSuite))
}
