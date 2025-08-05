package xmlparser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type XmlParserTestSuite struct {
	suite.Suite
	mockData    []byte
	mockError   error
	mockFile    file.File
	mockDiagram file.Diagram

	mockXmlHelper *mocks.XmlHelperMock
}

func (suite *XmlParserTestSuite) SetupTest() {
	suite.mockData = []byte{}
	suite.mockError = errors.New("A general error has occurred")
	suite.mockDiagram = file.Diagram{}
	suite.mockFile = file.File{Diagrams: []file.Diagram{suite.mockDiagram}}

	suite.mockXmlHelper = mocks.NewXmlHelperMock()
}

func (suite *XmlParserTestSuite) TestNewXmlParser_whenCreateXmlParser_thenReturnXmlParser() {
	xmlHelper := NewXmlParser(suite.mockXmlHelper)

	suite.NotNil(xmlHelper)
	suite.IsType(&XmlParserImpl{}, xmlHelper)
}

func (suite *XmlParserTestSuite) TestParseDiagrams_whenErrorOnUnmarshal_thenReturnError() {
	suite.mockXmlHelper.On("Unmarshal", suite.mockData, mock.AnythingOfType("*file.File")).Return(suite.mockError)

	diagrams, err := NewXmlParser(suite.mockXmlHelper).ParseDiagrams(suite.mockData)

	suite.Nil(diagrams)
	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *XmlParserTestSuite) TestParseDiagrams_whenParseValidChartFile_thenReturnParsedDiagrams() {
	suite.mockXmlHelper.On("Unmarshal", suite.mockData, mock.AnythingOfType("*file.File")).
		Run(func(args mock.Arguments) {
			entry := args.Get(1).(*file.File)
			*entry = suite.mockFile
		}).
		Return(nil)

	diagrams, err := NewXmlParser(suite.mockXmlHelper).ParseDiagrams(suite.mockData)

	suite.Equal(suite.mockFile.Diagrams[0], diagrams[0])
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestChartManagerTestSuite(t *testing.T) {
	suite.Run(t, new(XmlParserTestSuite))
}
