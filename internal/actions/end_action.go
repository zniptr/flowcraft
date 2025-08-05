package actions

import (
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type EndAction struct{}

func NewEndAction() Action {
	return &EndAction{}
}

func (action *EndAction) Execute(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) error {
	return nil
}

func (action *EndAction) GetNext(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error) {
	return nil, nil
}
