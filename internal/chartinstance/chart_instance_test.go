package chartinstance

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/actions"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type ChartInstanceTestSuite struct {
	suite.Suite
	mockError error

	mockChart            *mocks.ChartMock
	mockChartContext     *mocks.ChartContextMock
	mockStartAction      *mocks.ActionMock
	mockEndAction        *mocks.ActionMock
	mockProcessAction    *mocks.ActionMock
	mockPredefinedAction *mocks.ActionMock
	mockDecisionAction   *mocks.ActionMock
	mockStartObject      *file.Object
	mockProcessObject    *file.Object
	mockPredefinedObject *file.Object
	mockDecisionObject   *file.Object
	mockEndObject        *file.Object
}

func (suite *ChartInstanceTestSuite) SetupTest() {
	suite.mockError = errors.New("A general error has occurred")

	suite.mockChart = mocks.NewChartMock()
	suite.mockChartContext = mocks.NewChartContextMock()
	suite.mockStartAction = mocks.NewActionMock()
	suite.mockEndAction = mocks.NewActionMock()
	suite.mockProcessAction = mocks.NewActionMock()
	suite.mockPredefinedAction = mocks.NewActionMock()
	suite.mockDecisionAction = mocks.NewActionMock()
	suite.mockStartObject = &file.Object{Type: "start"}
	suite.mockProcessObject = &file.Object{Type: "process"}
	suite.mockPredefinedObject = &file.Object{Type: "predefined"}
	suite.mockDecisionObject = &file.Object{Type: "decision"}
	suite.mockEndObject = &file.Object{Type: "end"}

	newStartActionFunc = suite.newStartActionMock
	newEndActionFunc = suite.newEndActionMock
	newProcessActionFunc = suite.newProcessActionMock
	newPredefinedActionFunc = suite.newPredefinedActionMock
	newDecisionActionFunc = suite.newDecisionActionMock
}

func (suite *ChartInstanceTestSuite) newStartActionMock() actions.Action {
	return suite.mockStartAction
}

func (suite *ChartInstanceTestSuite) newEndActionMock() actions.Action {
	return suite.mockEndAction
}

func (suite *ChartInstanceTestSuite) newProcessActionMock() actions.Action {
	return suite.mockProcessAction
}

func (suite *ChartInstanceTestSuite) newPredefinedActionMock() actions.Action {
	return suite.mockPredefinedAction
}

func (suite *ChartInstanceTestSuite) newDecisionActionMock() actions.Action {
	return suite.mockDecisionAction
}

func (suite *ChartInstanceTestSuite) TestNewChartManager_whenCreateChartManager_thenReturnChartManager() {
	chartInstance := NewChartInstance(suite.mockChartContext, suite.mockChart)

	suite.NotNil(chartInstance)
	suite.IsType(&ChartInstanceImpl{}, chartInstance)
}

func (suite *ChartInstanceTestSuite) TestRun_whenNoStart_thenReturnNil() {
	suite.mockChart.On("GetStart").Return(nil)

	err := NewChartInstance(suite.mockChartContext, suite.mockChart).Run()

	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartInstanceTestSuite) TestRun_whenErrorOnExecute_thenReturnError() {
	suite.mockChart.On("GetStart").Return(suite.mockStartObject)
	suite.mockStartAction.On("Execute", suite.mockChartContext, suite.mockChart, suite.mockStartObject).Return(suite.mockError)

	err := NewChartInstance(suite.mockChartContext, suite.mockChart).Run()

	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartInstanceTestSuite) TestRun_whenErrorOnGetNext_thenReturnError() {
	suite.mockChart.On("GetStart").Return(suite.mockStartObject)
	suite.mockStartAction.On("Execute", suite.mockChartContext, suite.mockChart, suite.mockStartObject).Return(nil)
	suite.mockStartAction.On("GetNext", suite.mockChartContext, suite.mockChart, suite.mockStartObject).Return(nil, suite.mockError)

	err := NewChartInstance(suite.mockChartContext, suite.mockChart).Run()

	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartInstanceTestSuite) TestRun_whenChartEnds_thenReturnNil() {
	suite.mockChart.On("GetStart").Return(suite.mockStartObject)
	suite.mockStartAction.On("Execute", suite.mockChartContext, suite.mockChart, suite.mockStartObject).Return(nil)
	suite.mockStartAction.On("GetNext", suite.mockChartContext, suite.mockChart, suite.mockStartObject).Return(suite.mockProcessObject, nil)
	suite.mockProcessAction.On("Execute", suite.mockChartContext, suite.mockChart, suite.mockProcessObject).Return(nil)
	suite.mockProcessAction.On("GetNext", suite.mockChartContext, suite.mockChart, suite.mockProcessObject).Return(suite.mockPredefinedObject, nil)
	suite.mockPredefinedAction.On("Execute", suite.mockChartContext, suite.mockChart, suite.mockPredefinedObject).Return(nil)
	suite.mockPredefinedAction.On("GetNext", suite.mockChartContext, suite.mockChart, suite.mockPredefinedObject).Return(suite.mockDecisionObject, nil)
	suite.mockDecisionAction.On("Execute", suite.mockChartContext, suite.mockChart, suite.mockDecisionObject).Return(nil)
	suite.mockDecisionAction.On("GetNext", suite.mockChartContext, suite.mockChart, suite.mockDecisionObject).Return(suite.mockEndObject, nil)
	suite.mockEndAction.On("Execute", suite.mockChartContext, suite.mockChart, suite.mockEndObject).Return(nil)
	suite.mockEndAction.On("GetNext", suite.mockChartContext, suite.mockChart, suite.mockEndObject).Return(nil, nil)

	err := NewChartInstance(suite.mockChartContext, suite.mockChart).Run()

	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestChartInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(ChartInstanceTestSuite))
}
