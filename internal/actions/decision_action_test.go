package actions

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/gojaexecutor"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type DecisionActionTestSuite struct {
	suite.Suite
	mockId        string
	mockTargetId  string
	mockLabel     string
	mockCondition string
	mockError     error
	mockContext   map[string]any

	mockGojaExecutor                 *mocks.GojaExecutorMock
	mockChartContext                 *mocks.ChartContextMock
	mockChart                        *mocks.ChartMock
	mockObject                       *file.Object
	mockNonDefaultOutgoingConnection *file.Object
	mockDefaultOutgoingConnection    *file.Object
	mockNextObject                   *file.Object
}

func (suite *DecisionActionTestSuite) SetupTest() {
	suite.mockId = "testId"
	suite.mockTargetId = "testTargetId"
	suite.mockLabel = "testLabel"
	suite.mockCondition = "testCondition"
	suite.mockError = errors.New("A general error has occurred")
	suite.mockContext = make(map[string]any)

	suite.mockGojaExecutor = mocks.NewGojaExecutorMock()
	suite.mockChartContext = mocks.NewChartContextMock()
	suite.mockChart = mocks.NewChartMock()
	suite.mockObject = &file.Object{Id: suite.mockId, Label: suite.mockLabel}
	suite.mockNonDefaultOutgoingConnection = &file.Object{Condition: suite.mockCondition, Label: suite.mockLabel, Cell: file.Cell{Target: suite.mockTargetId}}
	suite.mockDefaultOutgoingConnection = &file.Object{Default: "1", Label: suite.mockLabel, Cell: file.Cell{Target: suite.mockTargetId}}
	suite.mockNextObject = &file.Object{}

	newGojaExecutorFunc = suite.newGojaExecutorMock
}

func (suite *DecisionActionTestSuite) newGojaExecutorMock() gojaexecutor.GojaExecutor {
	return suite.mockGojaExecutor
}

func (suite *DecisionActionTestSuite) TestNewDecisionAction_whenNewDecisionAction_thenReturnDecisionAction() {
	decisionAction := NewDecisionAction()

	suite.NotNil(decisionAction)
	suite.IsType(&DecisionAction{}, decisionAction)
}

func (suite *DecisionActionTestSuite) TestDecisionActionExecute_whenExecute_thenReturnNil() {
	err := NewDecisionAction().Execute(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(err)
}

func (suite *DecisionActionTestSuite) TestGetNext_whenNoOutgoingConnections_thenReturnError() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{})
	suite.mockChart.On("GetOutgoingDefaultConnection", suite.mockId).Return(nil)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("no default connection for decision action %s", suite.mockObject.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *DecisionActionTestSuite) TestGetNext_whenErrorOnSetVariable_thenReturnError() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{suite.mockNonDefaultOutgoingConnection})
	suite.mockChartContext.On("GetContext").Return(suite.mockContext)
	suite.mockGojaExecutor.On("SetVariables", suite.mockContext).Return(suite.mockError)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *DecisionActionTestSuite) TestGetNext_whenErrorOnRun_thenReturnError() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{suite.mockNonDefaultOutgoingConnection})
	suite.mockChartContext.On("GetContext").Return(suite.mockContext)
	suite.mockGojaExecutor.On("SetVariables", suite.mockContext).Return(nil)
	suite.mockGojaExecutor.On("Run", suite.mockCondition).Return(nil, suite.mockError)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *DecisionActionTestSuite) TestGetNext_whenErrorOnCast_thenReturnError() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{suite.mockNonDefaultOutgoingConnection})
	suite.mockChartContext.On("GetContext").Return(suite.mockContext)
	suite.mockGojaExecutor.On("SetVariables", suite.mockContext).Return(nil)
	suite.mockGojaExecutor.On("Run", suite.mockCondition).Return("", nil)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("condition result is not a boolean for connection %s", suite.mockNonDefaultOutgoingConnection.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *DecisionActionTestSuite) TestGetNext_whenNoEvaluatedOutgoingAndDefaultConnection_thenReturnError() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{suite.mockNonDefaultOutgoingConnection})
	suite.mockChartContext.On("GetContext").Return(suite.mockContext)
	suite.mockGojaExecutor.On("SetVariables", suite.mockContext).Return(nil)
	suite.mockGojaExecutor.On("Run", suite.mockCondition).Return(false, nil)
	suite.mockChart.On("GetOutgoingDefaultConnection", suite.mockId).Return(nil)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("no default connection for decision action %s", suite.mockObject.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *DecisionActionTestSuite) TestGetNext_whenErrorOnGetObjectForEvaluatedConnection_thenReturnError() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{suite.mockNonDefaultOutgoingConnection})
	suite.mockChartContext.On("GetContext").Return(suite.mockContext)
	suite.mockGojaExecutor.On("SetVariables", suite.mockContext).Return(nil)
	suite.mockGojaExecutor.On("Run", suite.mockCondition).Return(true, nil)
	suite.mockChart.On("GetObject", suite.mockTargetId).Return(nil)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("no target object found for connection %s", suite.mockNonDefaultOutgoingConnection.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *DecisionActionTestSuite) TestGetNext_whenErrorOnGetObjectForDefaultConnection_thenReturnError() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{})
	suite.mockChart.On("GetOutgoingDefaultConnection", suite.mockId).Return(suite.mockDefaultOutgoingConnection)
	suite.mockChartContext.On("GetContext").Return(suite.mockContext)
	suite.mockGojaExecutor.On("SetVariables", suite.mockContext).Return(nil)
	suite.mockGojaExecutor.On("Run", suite.mockCondition).Return(true, nil)
	suite.mockChart.On("GetObject", suite.mockTargetId).Return(nil)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("no target object found for connection %s", suite.mockDefaultOutgoingConnection.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *DecisionActionTestSuite) TestGetNext_whenGetObjectForEvaluatedConnection_thenReturnNext() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{suite.mockNonDefaultOutgoingConnection})
	suite.mockChartContext.On("GetContext").Return(suite.mockContext)
	suite.mockGojaExecutor.On("SetVariables", suite.mockContext).Return(nil)
	suite.mockGojaExecutor.On("Run", suite.mockCondition).Return(true, nil)
	suite.mockChart.On("GetObject", suite.mockTargetId).Return(suite.mockNextObject)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Equal(suite.mockNextObject, next)
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *DecisionActionTestSuite) TestGetNext_whenGetObjectForDefaultConnection_thenReturnNext() {
	suite.mockChart.On("GetOutgoingNonDefaultConnections", suite.mockId).Return([]*file.Object{})
	suite.mockChart.On("GetOutgoingDefaultConnection", suite.mockId).Return(suite.mockDefaultOutgoingConnection)
	suite.mockChartContext.On("GetContext").Return(suite.mockContext)
	suite.mockGojaExecutor.On("SetVariables", suite.mockContext).Return(nil)
	suite.mockGojaExecutor.On("Run", suite.mockCondition).Return(true, nil)
	suite.mockChart.On("GetObject", suite.mockTargetId).Return(suite.mockNextObject)

	next, err := NewDecisionAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Equal(suite.mockNextObject, next)
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestDecisionActionTestSuite(t *testing.T) {
	suite.Run(t, new(DecisionActionTestSuite))
}
