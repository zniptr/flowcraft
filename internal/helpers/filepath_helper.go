package helpers

import "path/filepath"

type FilepathHelper interface {
	Join(elem ...string) string
	Ext(path string) string
}

type FilepathHelperImpl struct{}

func NewFilepathHelper() FilepathHelper {
	return &FilepathHelperImpl{}
}

func (helper *FilepathHelperImpl) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func (helper *FilepathHelperImpl) Ext(path string) string {
	return filepath.Ext(path)
}
