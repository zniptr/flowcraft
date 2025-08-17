package actions

import (
	"fmt"

	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type StartAction struct{}

func NewStartAction() Action {
	return &StartAction{}
}

func (action *StartAction) Execute(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) error {
	return nil
}

func (action *StartAction) GetNext(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error) {
	connection := chart.GetOutgoingConnection(object.Id)

	if connection == nil {
		return nil, fmt.Errorf("no source connection for start action %s", object.Label)
	}

	return action.resolveTarget(chart, connection)
}

func (action *StartAction) resolveTarget(chart chart.Chart, connection *file.Object) (*file.Object, error) {
	target := chart.GetObject(connection.Cell.Target)
	if target == nil {
		return nil, fmt.Errorf("no target object found for connection %s", connection.Label)
	}
	return target, nil
}
