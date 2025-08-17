package helpers

import (
	"github.com/dop251/goja"
)

type GojaValueHelper interface {
	Export() interface{}
}

type GojaValueHelperImpl struct {
	value goja.Value
}

func NewGojaValueHelper(value goja.Value) GojaValueHelper {
	return &GojaValueHelperImpl{
		value: value,
	}
}

func (helper *GojaValueHelperImpl) Export() interface{} {
	return helper.value.Export()
}

type GojaHelper interface {
	Set(name string, value interface{}) error
	RunString(str string) (GojaValueHelper, error)
}

type GojaHelperImpl struct {
	goja *goja.Runtime
}

func NewGojaHelper() GojaHelper {
	return &GojaHelperImpl{
		goja: goja.New(),
	}
}

func (helper *GojaHelperImpl) Set(name string, value interface{}) error {
	return helper.goja.Set(name, value)
}

func (helper *GojaHelperImpl) RunString(str string) (GojaValueHelper, error) {
	value, err := helper.goja.RunString(str)
	if err != nil {
		return nil, err
	}

	return NewGojaValueHelper(value), nil
}
