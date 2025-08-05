package chartmanager

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/chartinstance"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/filereader"
	"github.com/zniptr/flowcraft/internal/helpers"
	"github.com/zniptr/flowcraft/internal/mocks"
	"github.com/zniptr/flowcraft/internal/xmlparser"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type ChartManagerTestSuite struct {
	suite.Suite
	mockPath           string
	mockError          error
	mockFileContent    []byte
	mockChartNameWrong string
	mockChartName      string
	mockContext        map[string]any
	mockDiagram        file.Diagram

	mockDirEntry      *mocks.DirEntryHelperMock
	mockFileReader    *mocks.ChartFileReaderMock
	mockXmlParser     *mocks.ChartXmlParserMock
	mockChartContext  *mocks.ChartContextMock
	mockChartInstance *mocks.ChartInstanceMock
}

func (suite *ChartManagerTestSuite) SetupTest() {
	suite.mockPath = "charts"
	suite.mockError = errors.New("A general error has occurred")
	suite.mockFileContent = []byte{}
	suite.mockChartNameWrong = "nonexistent"
	suite.mockChartName = "existent"
	suite.mockContext = map[string]any{}
	suite.mockDiagram = file.Diagram{Name: suite.mockChartName}

	suite.mockDirEntry = mocks.NewDirEntryHelperMock()
	suite.mockFileReader = mocks.NewChartFileReaderMock()
	suite.mockXmlParser = mocks.NewChartXmlParserMock()
	suite.mockChartContext = mocks.NewChartContextMock()
	suite.mockChartInstance = mocks.NewChartInstanceMock()

	newChartFileReaderFunc = suite.newChartFileReaderMock
	newChartXmlParserFunc = suite.newXmlParserMock
	newChartContextFunc = suite.newChartContextMock
	newChartInstanceFunc = suite.newChartInstanceMock
}

func (suite *ChartManagerTestSuite) newChartFileReaderMock(_ helpers.OsHelper, _ helpers.FilepathHelper) filereader.FileReader {
	return suite.mockFileReader
}

func (suite *ChartManagerTestSuite) newXmlParserMock(_ helpers.XmlHelper) xmlparser.XmlParser {
	return suite.mockXmlParser
}

func (suite *ChartManagerTestSuite) newChartContextMock(_ map[string]any) chartcontext.ChartContext {
	return suite.mockChartContext
}

func (suite *ChartManagerTestSuite) newChartInstanceMock(_ chartcontext.ChartContext, _ chart.Chart) chartinstance.ChartInstance {
	return suite.mockChartInstance
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
	suite.mockXmlParser.On("ParseDiagrams", suite.mockFileContent).Return([]file.Diagram{suite.mockDiagram}, nil)

	err := NewChartManager().LoadCharts(suite.mockPath)

	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartManagerTestSuite) TestStartChartInstance_whenChartNotFound_thenReturnError() {
	err := NewChartManager().StartChartInstance(suite.mockChartNameWrong, suite.mockContext)

	suite.EqualError(err, "chart not found")
}

func (suite *ChartManagerTestSuite) TestStartChartInstance_whenErrorOnRun_thenReturnError() {
	suite.mockChartInstance.On("Run").Return(suite.mockError)

	manager := NewChartManager()
	manager.storeCharts([]file.Diagram{suite.mockDiagram})

	err := manager.StartChartInstance(suite.mockChartName, suite.mockContext)

	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *ChartManagerTestSuite) TestStartChartInstance_whenRunSuccessfull_thenReturnNil() {
	suite.mockChartInstance.On("Run").Return(nil)

	manager := NewChartManager()
	manager.storeCharts([]file.Diagram{suite.mockDiagram})

	err := manager.StartChartInstance(suite.mockChartName, suite.mockContext)

	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestChartManagerTestSuite(t *testing.T) {
	suite.Run(t, new(ChartManagerTestSuite))
}
