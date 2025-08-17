package actions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type StartActionTestSuite struct {
	suite.Suite
	mockId              string
	mockLabel           string
	mockTargetId        string
	mockConnectionLabel string

	mockChartContext     *mocks.ChartContextMock
	mockChart            *mocks.ChartMock
	mockObject           *file.Object
	mockConnectionObject *file.Object
	mockNextObject       *file.Object
}

func (suite *StartActionTestSuite) SetupTest() {
	suite.mockId = "testId"
	suite.mockLabel = "testLabel"
	suite.mockTargetId = "testTargetId"
	suite.mockConnectionLabel = "testConnectionLabel"

	suite.mockChartContext = mocks.NewChartContextMock()
	suite.mockChart = mocks.NewChartMock()
	suite.mockObject = &file.Object{Id: suite.mockId, Label: suite.mockLabel}
	suite.mockConnectionObject = &file.Object{Label: suite.mockConnectionLabel, Cell: file.Cell{Target: suite.mockTargetId}}
	suite.mockNextObject = &file.Object{}
}

func (suite *StartActionTestSuite) TestNewStartAction_whenNewStartAction_thenReturnStartAction() {
	startAction := NewStartAction()

	suite.NotNil(startAction)
	suite.IsType(&StartAction{}, startAction)
}

func (suite *StartActionTestSuite) TestStartActionExecute_whenExecute_thenReturnNil() {
	err := NewStartAction().Execute(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(err)
}

func (suite *StartActionTestSuite) TestStartActionGetNext_whenNoConnection_thenReturnError() {
	suite.mockChart.On("GetOutgoingConnection", suite.mockId).Return(nil)

	next, err := NewStartAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("no source connection for start action %s", suite.mockObject.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *StartActionTestSuite) TestStartActionGetNext_whenConnectionHasNoTarget_thenReturnError() {
	suite.mockChart.On("GetOutgoingConnection", suite.mockId).Return(suite.mockConnectionObject)
	suite.mockChart.On("GetObject", suite.mockTargetId).Return(nil)

	next, err := NewStartAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.EqualError(err, fmt.Sprintf("no target object found for connection %s", suite.mockConnectionObject.Label))
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *StartActionTestSuite) TestStartActionGetNext_whenNextFound_thenReturnNext() {
	suite.mockChart.On("GetOutgoingConnection", suite.mockId).Return(suite.mockConnectionObject)
	suite.mockChart.On("GetObject", suite.mockTargetId).Return(suite.mockNextObject)

	next, err := NewStartAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Equal(suite.mockNextObject, next)
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestStartActionTestSuite(t *testing.T) {
	suite.Run(t, new(StartActionTestSuite))
}
