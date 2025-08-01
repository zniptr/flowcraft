package filereader

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/helpers"
	"github.com/zniptr/flowcraft/internal/mocks"
)

type FileReaderTestSuite struct {
	suite.Suite
	mockPath               string
	mockError              error
	mockFileName           string
	mockFileExtension      string
	mockFileNameWrong      string
	mockFileExtensionWrong string
	mockFullPath           string
	mockData               []byte

	mockDirEntryHelper *mocks.DirEntryHelperMock
	mockOsHelper       *mocks.OsHelperMock
	mockFilepathHelper *mocks.FilepathHelperMock
}

func (suite *FileReaderTestSuite) SetupTest() {
	suite.mockPath = "charts"
	suite.mockError = errors.New("A general error has occurred")
	suite.mockFileName = "test.drawio"
	suite.mockFileExtension = ".drawio"
	suite.mockFileNameWrong = "test.txt"
	suite.mockFileExtensionWrong = ".txt"
	suite.mockFullPath = fmt.Sprintf("%s/%s", suite.mockPath, suite.mockFileName)
	suite.mockData = []byte{}

	suite.mockDirEntryHelper = mocks.NewDirEntryHelperMock()
	suite.mockOsHelper = mocks.NewOsHelperMock()
	suite.mockFilepathHelper = mocks.NewFilepathHelperMock()
}

func (suite *FileReaderTestSuite) TestNewFileReader_whenCreateFileReader_thenReturnFileReader() {
	fileReader := NewFileReader(suite.mockOsHelper, suite.mockFilepathHelper)

	suite.NotNil(fileReader)
	suite.IsType(&FileReaderImpl{}, fileReader)
}

func (suite *FileReaderTestSuite) TestReadDirectory_whenErrorOnReadDir_thenReturnError() {
	suite.mockOsHelper.On("ReadDir", suite.mockPath).Return(nil, suite.mockError)

	files, err := NewFileReader(suite.mockOsHelper, suite.mockFilepathHelper).ReadDirectory(suite.mockPath)

	suite.Nil(files)
	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *FileReaderTestSuite) TestReadDirectory_whenReadDir_thenReturnEntries() {
	suite.mockOsHelper.On("ReadDir", suite.mockPath).Return([]helpers.DirEntryHelper{}, nil)

	files, err := NewFileReader(suite.mockOsHelper, suite.mockFilepathHelper).ReadDirectory(suite.mockPath)

	suite.NotNil(files)
	suite.Len(files, 0)
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *FileReaderTestSuite) TestReadFile_whenErrorOnReadFile_thenReturnError() {
	suite.mockDirEntryHelper.On("Name").Return(suite.mockFileName)
	suite.mockFilepathHelper.On("Join", []string{suite.mockPath, suite.mockFileName}).Return(suite.mockFullPath)
	suite.mockOsHelper.On("ReadFile", suite.mockFullPath).Return(nil, suite.mockError)

	data, err := NewFileReader(suite.mockOsHelper, suite.mockFilepathHelper).ReadFile(suite.mockPath, suite.mockDirEntryHelper)

	suite.Nil(data)
	suite.EqualError(err, suite.mockError.Error())
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *FileReaderTestSuite) TestReadFile_whenReadFile_thenReturnData() {
	suite.mockDirEntryHelper.On("Name").Return(suite.mockFileName)
	suite.mockFilepathHelper.On("Join", []string{suite.mockPath, suite.mockFileName}).Return(suite.mockFullPath)
	suite.mockOsHelper.On("ReadFile", suite.mockFullPath).Return(suite.mockData, nil)

	data, err := NewFileReader(suite.mockOsHelper, suite.mockFilepathHelper).ReadFile(suite.mockPath, suite.mockDirEntryHelper)

	suite.Equal(suite.mockData, data)
	suite.Nil(err)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *FileReaderTestSuite) TestIsValidChartFile_whenIsDirectory_thenReturnFalse() {
	suite.mockDirEntryHelper.On("IsDir").Return(true)

	isValidChartFile := NewFileReader(suite.mockOsHelper, suite.mockFilepathHelper).IsValidChartFile(suite.mockDirEntryHelper)

	suite.False(isValidChartFile)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *FileReaderTestSuite) TestIsValidChartFile_whenIsWrongExtension_thenReturnFalse() {
	suite.mockDirEntryHelper.On("IsDir").Return(false)
	suite.mockDirEntryHelper.On("Name").Return(suite.mockFileNameWrong)
	suite.mockFilepathHelper.On("Ext", suite.mockFileNameWrong).Return(suite.mockFileExtensionWrong)

	isValidChartFile := NewFileReader(suite.mockOsHelper, suite.mockFilepathHelper).IsValidChartFile(suite.mockDirEntryHelper)

	suite.False(isValidChartFile)
	mock.AssertExpectationsForObjects(suite.T())
}

func (suite *FileReaderTestSuite) TestIsValidChartFile_whenValid_thenReturnTrue() {
	suite.mockDirEntryHelper.On("IsDir").Return(false)
	suite.mockDirEntryHelper.On("Name").Return(suite.mockFileName)
	suite.mockFilepathHelper.On("Ext", suite.mockFileName).Return(suite.mockFileExtension)

	isValidChartFile := NewFileReader(suite.mockOsHelper, suite.mockFilepathHelper).IsValidChartFile(suite.mockDirEntryHelper)

	suite.True(isValidChartFile)
	mock.AssertExpectationsForObjects(suite.T())
}

func TestFileReaderTestSuite(t *testing.T) {
	suite.Run(t, new(FileReaderTestSuite))
}
