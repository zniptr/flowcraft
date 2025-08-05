package chartinstance

import (
	"github.com/zniptr/flowcraft/internal/actions"
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type ChartInstance interface {
	Run() error
}

type ChartInstanceImpl struct {
	chartContext chartcontext.ChartContext
	chart        chart.Chart
	actions      map[string]func() actions.Action
}

var (
	newStartActionFunc      = actions.NewStartAction
	newEndActionFunc        = actions.NewEndAction
	newProcessActionFunc    = actions.NewProcessAction
	newPredefinedActionFunc = actions.NewPredefinedProcessAction
	newDecisionActionFunc   = actions.NewDecisionAction
)

func NewChartInstance(chartContext chartcontext.ChartContext, chart chart.Chart) ChartInstance {
	return &ChartInstanceImpl{
		chartContext: chartContext,
		chart:        chart,
		actions: map[string]func() actions.Action{
			"start":      func() actions.Action { return newStartActionFunc() },
			"end":        func() actions.Action { return newEndActionFunc() },
			"process":    func() actions.Action { return newProcessActionFunc() },
			"predefined": func() actions.Action { return newPredefinedActionFunc() },
			"decision":   func() actions.Action { return newDecisionActionFunc() },
		},
	}
}

func (instance *ChartInstanceImpl) Run() error {
	var err error
	next := instance.chart.GetStart()

	for next != nil {
		next, err = instance.executeAction(next)
		if err != nil {
			return err
		}
	}

	return nil
}

func (instance *ChartInstanceImpl) executeAction(object *file.Object) (*file.Object, error) {
	action := instance.actions[object.Type]()

	err := action.Execute(instance.chartContext, instance.chart, object)
	if err != nil {
		return nil, err
	}

	return action.GetNext(instance.chartContext, instance.chart, object)
}
