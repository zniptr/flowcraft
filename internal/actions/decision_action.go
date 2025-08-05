package actions

import (
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type DecisionAction struct{}

func NewDecisionAction() Action {
	return &DecisionAction{}
}

func (action *DecisionAction) Execute(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) error {
	return nil
}

func (action *DecisionAction) GetNext(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error) {
	return nil, nil
}
