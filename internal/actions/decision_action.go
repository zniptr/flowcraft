package actions

import (
	"fmt"

	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/gojaexecutor"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type DecisionAction struct {
	vm gojaexecutor.GojaExecutor
}

var (
	newGojaExecutorFunc = gojaexecutor.NewGojaExecutor
)

func NewDecisionAction() Action {
	return &DecisionAction{
		vm: newGojaExecutorFunc(),
	}
}

func (action *DecisionAction) Execute(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) error {
	return nil
}

func (action *DecisionAction) GetNext(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error) {
	evaluatedConnection, err := action.getEvaluatedConnection(chartContext, chart, object)
	if err != nil {
		return nil, err
	}

	if evaluatedConnection != nil {
		return action.resolveTarget(chart, evaluatedConnection)
	}

	defaultConnection := chart.GetOutgoingDefaultConnection(object.Id)
	if defaultConnection != nil {
		return action.resolveTarget(chart, defaultConnection)
	}

	return nil, fmt.Errorf("no default connection for decision action %s", object.Label)
}

func (action *DecisionAction) getEvaluatedConnection(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error) {
	connections := chart.GetOutgoingNonDefaultConnections(object.Id)
	if len(connections) == 0 {
		return nil, nil
	}

	err := action.vm.SetVariables(chartContext.GetContext())
	if err != nil {
		return nil, err
	}

	for _, connection := range connections {
		ok, err := action.evaluateCondition(connection)
		if err != nil {
			return nil, err
		}

		if ok {
			return connection, nil
		}
	}

	return nil, nil
}

func (action *DecisionAction) evaluateCondition(connection *file.Object) (bool, error) {
	result, err := action.vm.Run(connection.Condition)
	if err != nil {
		return false, err
	}

	boolResult, ok := result.(bool)
	if !ok {
		return false, fmt.Errorf("condition result is not a boolean for connection %s", connection.Label)
	}

	return boolResult, nil
}

func (action *DecisionAction) resolveTarget(chart chart.Chart, connection *file.Object) (*file.Object, error) {
	target := chart.GetObject(connection.Cell.Target)
	if target == nil {
		return nil, fmt.Errorf("no target object found for connection %s", connection.Label)
	}

	return target, nil
}
