package actions

import (
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type PredefinedProcessAction struct{}

func NewPredefinedProcessAction() Action {
	return &PredefinedProcessAction{}
}

func (action *PredefinedProcessAction) Execute(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) error {
	return nil
}

func (action *PredefinedProcessAction) GetNext(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error) {
	return nil, nil
}
