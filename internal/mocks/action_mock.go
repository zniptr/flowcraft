package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type ActionMock struct {
	mock.Mock
}

func NewActionMock() *ActionMock {
	return &ActionMock{}
}

func (mock *ActionMock) Execute(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) error {
	args := mock.Called(chartContext, chart, object)

	return args.Error(0)
}

func (mock *ActionMock) GetNext(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error) {
	args := mock.Called(chartContext, chart, object)

	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*file.Object), args.Error(1)
}
