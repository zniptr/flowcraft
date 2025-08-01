package filereader

import (
	"github.com/zniptr/flowcraft/internal/helpers"
)

type FileReader interface {
	ReadDirectory(path string) ([]helpers.DirEntryHelper, error)
	ReadFile(path string, file helpers.DirEntryHelper) ([]byte, error)
	IsValidChartFile(file helpers.DirEntryHelper) bool
}

type FileReaderImpl struct {
	osHelper       helpers.OsHelper
	filepathHelper helpers.FilepathHelper
}

func NewFileReader(osHelper helpers.OsHelper, filepathHelper helpers.FilepathHelper) FileReader {
	return &FileReaderImpl{
		osHelper:       osHelper,
		filepathHelper: filepathHelper,
	}
}

var extension = ".drawio"

func (reader *FileReaderImpl) ReadDirectory(path string) ([]helpers.DirEntryHelper, error) {
	return reader.osHelper.ReadDir(path)
}

func (reader *FileReaderImpl) ReadFile(path string, file helpers.DirEntryHelper) ([]byte, error) {
	fullPath := reader.filepathHelper.Join(path, file.Name())
	return reader.osHelper.ReadFile(fullPath)
}

func (reader *FileReaderImpl) IsValidChartFile(file helpers.DirEntryHelper) bool {
	return !file.IsDir() && reader.filepathHelper.Ext(file.Name()) == extension
}
