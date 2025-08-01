package chartmanager

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/helpers"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type ChartManagerTestSuite struct {
	suite.Suite
	mockPath        string
	mockError       error
	mockDirEntry    *mocks.DirEntryHelperMock
	mockFileContent []byte
	mockDiagram     chart.Diagram

	mockFileReader *mocks.ChartFileReaderMock
	mockXmlParser  *mocks.ChartXmlParserMock
}

func (suite *ChartManagerTestSuite) SetupTest() {
	suite.mockPath = "charts"
	suite.mockError = errors.New("A general error has occurred")
	suite.mockDirEntry = mocks.NewDirEntryHelperMock()
	suite.mockFileContent = []byte{}
	suite.mockDiagram = chart.Diagram{Id: "123"}

	suite.mockFileReader = mocks.NewChartFileReaderMock()
	suite.mockXmlParser = mocks.NewChartXmlParserMock()

	newChartFileReaderFunc = newChartFileReaderMock(suite)
	newChartXmlParserFunc = newChartXmlParserMock(suite)
}

func newChartFileReaderMock(suite *ChartManagerTestSuite) *mocks.ChartFileReaderMock {
	return suite.mockFileReader
}

func newChartXmlParserMock(suite *ChartManagerTestSuite) *mocks.ChartXmlParserMock {
	return suite.mockXmlParser
}

func (suite *ChartManagerTestSuite) TestNewChartManager_whenCreateChartManager_thenReturnChartManager() {
	chartManager := NewChartManager()

	suite.NotNil(chartManager)
	suite.IsType(&chartManagerImpl{}, chartManager)
}

func (suite *ChartManagerTestSuite) TestLoadCharts_whenErrorOnReadDirectory_thenReturnError() {
	suite.mockFileReader.On("ReadDirectory", suite.mockPath).Return(nil, suite.mockError)

	err := NewChartManager().LoadCharts(suite.mockPath)

	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartManagerTestSuite) TestLoadCharts_whenIsNotValidChartFile_thenDontParse() {
	suite.mockFileReader.On("ReadDirectory", suite.mockPath).Return([]helpers.DirEntryHelper{suite.mockDirEntry}, nil)
	suite.mockFileReader.On("IsValidChartFile", suite.mockDirEntry).Return(false)

	err := NewChartManager().LoadCharts(suite.mockPath)

	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartManagerTestSuite) TestLoadCharts_whenErrorOnReadFile_thenReturnError() {
	suite.mockFileReader.On("ReadDirectory", suite.mockPath).Return([]helpers.DirEntryHelper{suite.mockDirEntry}, nil)
	suite.mockFileReader.On("IsValidChartFile", suite.mockDirEntry).Return(true)
	suite.mockFileReader.On("ReadFile", suite.mockPath, suite.mockDirEntry).Return(nil, suite.mockError)

	err := NewChartManager().LoadCharts(suite.mockPath)

	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartManagerTestSuite) TestLoadCharts_whenErrorOnParseDiagrams_thenReturnError() {
	suite.mockFileReader.On("ReadDirectory", suite.mockPath).Return([]helpers.DirEntryHelper{suite.mockDirEntry}, nil)
	suite.mockFileReader.On("IsValidChartFile", suite.mockDirEntry).Return(true)
	suite.mockFileReader.On("ReadFile", suite.mockPath, suite.mockDirEntry).Return(suite.mockFileContent, nil)
	suite.mockXmlParser.On("ParseDiagrams", suite.mockFileContent).Return(nil, suite.mockError)

	err := NewChartManager().LoadCharts(suite.mockPath)

	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartManagerTestSuite) TestLoadCharts_whenFileSuccessfullyParsed_thenStoreDiagrams() {
	suite.mockFileReader.On("ReadDirectory", suite.mockPath).Return([]helpers.DirEntryHelper{suite.mockDirEntry}, nil)
	suite.mockFileReader.On("IsValidChartFile", suite.mockDirEntry).Return(true)
	suite.mockFileReader.On("ReadFile", suite.mockPath, suite.mockDirEntry).Return(suite.mockFileContent, nil)
	suite.mockXmlParser.On("ParseDiagrams", suite.mockFileContent).Return([]chart.Diagram{suite.mockDiagram}, nil)

	err := NewChartManager().LoadCharts(suite.mockPath)

	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestChartManagerTestSuite(t *testing.T) {
	suite.Run(t, new(ChartManagerTestSuite))
}
