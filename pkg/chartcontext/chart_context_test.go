package chartcontext

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChartContextTestSuite struct {
	suite.Suite
	mockUnknownVariableName string
	mockKnownVariableName   string
	mockValue               int
	mockOverwriteValue      string

	mockContext map[string]any
}

func (suite *ChartContextTestSuite) SetupTest() {
	suite.mockUnknownVariableName = "unknownTest"
	suite.mockKnownVariableName = "knownTest"
	suite.mockValue = 1337
	suite.mockOverwriteValue = "leet"

	suite.mockContext = make(map[string]any)
}

func (suite *ChartContextTestSuite) TestNewChartContext_whenCreateChartContext_thenReturnChartContext() {
	chartContext := NewChartContext(suite.mockContext)

	suite.NotNil(chartContext)
	suite.IsType(&chartContextImpl{}, chartContext)
}

func (suite *ChartContextTestSuite) TestGetContext_whenGetContext_thenReturnContext() {
	context := NewChartContext(suite.mockContext).GetContext()

	suite.Equal(suite.mockContext, context)
}

func (suite *ChartContextTestSuite) TestGetVariable_whenGetUnknownVariable_thenReturnNil() {
	variable := NewChartContext(suite.mockContext).GetVariable(suite.mockUnknownVariableName)

	suite.Nil(variable)
}

func (suite *ChartContextTestSuite) TestGetVariable_whenGetKnownVariable_thenReturnVariable() {
	suite.mockContext[suite.mockKnownVariableName] = suite.mockValue

	variable := NewChartContext(suite.mockContext).GetVariable(suite.mockKnownVariableName)

	suite.Equal(suite.mockValue, variable)
}

func (suite *ChartContextTestSuite) TestSetVariable_whenSetVariable_thenAddValueToContext() {
	context := NewChartContext(suite.mockContext)
	context.SetVariable(suite.mockKnownVariableName, suite.mockValue)

	variable := context.GetVariable(suite.mockKnownVariableName)

	suite.Equal(suite.mockValue, variable)
}

func (suite *ChartContextTestSuite) TestSetVariable_whenSetKnownVariable_thenOverwriteValue() {
	suite.mockContext[suite.mockKnownVariableName] = suite.mockValue
	context := NewChartContext(suite.mockContext)
	context.SetVariable(suite.mockKnownVariableName, suite.mockOverwriteValue)

	variable := context.GetVariable(suite.mockKnownVariableName)

	suite.Equal(suite.mockOverwriteValue, variable)
}

func TestChartContextTestSuite(t *testing.T) {
	suite.Run(t, new(ChartContextTestSuite))
}
