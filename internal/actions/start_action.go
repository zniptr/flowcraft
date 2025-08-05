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
	connection := chart.GetSingleConnectionBySourceId(object.Id)

	if connection == nil {
		return nil, fmt.Errorf("no source connection for start action %s", object.Label)
	}

	target := chart.GetObjectById(connection.Cell.Target)
	if target == nil {
		return nil, fmt.Errorf("no target for connection action %s", connection.Label)
	}

	return target, nil
}
