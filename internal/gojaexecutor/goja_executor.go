package gojaexecutor

import "github.com/zniptr/flowcraft/internal/helpers"

type GojaExecutor interface {
	SetVariables(vars map[string]any) error
	Run(code string) (any, error)
}

type GojaExecutorImpl struct {
	vm helpers.GojaHelper
}

var (
	newGojaFunc = helpers.NewGojaHelper
)

func NewGojaExecutor() GojaExecutor {
	return &GojaExecutorImpl{
		vm: newGojaFunc(),
	}
}

func (executor *GojaExecutorImpl) SetVariables(variables map[string]any) error {
	for key, value := range variables {
		err := executor.vm.Set(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (executor *GojaExecutorImpl) Run(code string) (any, error) {
	value, err := executor.vm.RunString(code)
	if err != nil {
		return nil, err
	}
	return value.Export(), nil
}
