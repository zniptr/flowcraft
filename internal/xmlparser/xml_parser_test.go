package xmlparser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type XmlParserTestSuite struct {
	suite.Suite
	mockData  []byte
	mockError error
	mockChart chart.ChartFile

	mockXmlHelper *mocks.XmlHelperMock
}

func (suite *XmlParserTestSuite) SetupTest() {
	suite.mockData = []byte{}
	suite.mockError = errors.New("A general error has occurred")
	suite.mockChart = chart.ChartFile{Diagrams: []chart.Diagram{{Id: "123"}}}

	suite.mockXmlHelper = mocks.NewXmlHelperMock()
}

func (suite *XmlParserTestSuite) TestNewXmlParser_whenCreateXmlParser_thenReturnXmlParser() {
	xmlHelper := NewXmlParser(suite.mockXmlHelper)

	suite.NotNil(xmlHelper)
	suite.IsType(&XmlParserImpl{}, xmlHelper)
}

func (suite *XmlParserTestSuite) TestParseDiagrams_whenErrorOnUnmarshal_thenReturnError() {
	suite.mockXmlHelper.On("Unmarshal", suite.mockData, mock.AnythingOfType("*chart.ChartFile")).Return(suite.mockError)

	diagrams, err := NewXmlParser(suite.mockXmlHelper).ParseDiagrams(suite.mockData)

	suite.Nil(diagrams)
	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *XmlParserTestSuite) TestParseDiagrams_whenParseValidChartFile_thenReturnParsedDiagrams() {
	suite.mockXmlHelper.On("Unmarshal", suite.mockData, mock.AnythingOfType("*chart.ChartFile")).
		Run(func(args mock.Arguments) {
			entry := args.Get(1).(*chart.ChartFile)
			*entry = suite.mockChart
		}).
		Return(nil)

	diagrams, err := NewXmlParser(suite.mockXmlHelper).ParseDiagrams(suite.mockData)

	suite.Equal(suite.mockChart.Diagrams[0], diagrams[0])
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestChartManagerTestSuite(t *testing.T) {
	suite.Run(t, new(XmlParserTestSuite))
}
