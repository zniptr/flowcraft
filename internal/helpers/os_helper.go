package helpers

import "os"

type DirEntryHelper interface {
	IsDir() bool
	Name() string
}

type DirEntryHelperImpl struct {
	dirEntry os.DirEntry
}

func NewDirEntryHelper(dirEntry os.DirEntry) *DirEntryHelperImpl {
	return &DirEntryHelperImpl{
		dirEntry: dirEntry,
	}
}

func (helper DirEntryHelperImpl) IsDir() bool {
	return helper.dirEntry.IsDir()
}

func (helper DirEntryHelperImpl) Name() string {
	return helper.dirEntry.Name()
}

type OsHelper interface {
	ReadDir(name string) ([]DirEntryHelper, error)
	ReadFile(name string) ([]byte, error)
}

type OsHelperImpl struct{}

func NewOsHelper() OsHelper {
	return &OsHelperImpl{}
}

func (helper *OsHelperImpl) ReadDir(name string) ([]DirEntryHelper, error) {
	dirEntries, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	result := make([]DirEntryHelper, len(dirEntries))
	for i, entry := range dirEntries {
		result[i] = NewDirEntryHelper(entry)
	}

	return result, nil
}

func (helper *OsHelperImpl) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
