package actions

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type EndActionTestSuite struct {
	suite.Suite

	mockChartContext *mocks.ChartContextMock
	mockChart        *mocks.ChartMock
	mockObject       *file.Object
}

func (suite *EndActionTestSuite) SetupTest() {
	suite.mockChartContext = mocks.NewChartContextMock()
	suite.mockChart = mocks.NewChartMock()
	suite.mockObject = &file.Object{}
}

func (suite *EndActionTestSuite) TestNewEndAction_whenNewEndAction_thenReturnEndAction() {
	endAction := NewEndAction()

	suite.NotNil(endAction)
	suite.IsType(&EndAction{}, endAction)
}

func (suite *EndActionTestSuite) TestEndActionExecute_whenExecute_thenReturnNil() {
	err := NewEndAction().Execute(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(err)
}

func (suite *EndActionTestSuite) TestEndActionExecute_whenGetNext_thenReturnNil() {
	next, err := NewEndAction().GetNext(suite.mockChartContext, suite.mockChart, suite.mockObject)

	suite.Nil(next)
	suite.Nil(err)
}

func TestEndActionTestSuite(t *testing.T) {
	suite.Run(t, new(EndActionTestSuite))
}
