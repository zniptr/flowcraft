package actions

import (
	"fmt"

	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
	"github.com/zniptr/flowcraft/pkg/executableregistry"
)

type ProcessAction struct{}

var (
	getExecutableRegistryInstanceFunc = executableregistry.GetInstance
)

func NewProcessAction() Action {
	return &ProcessAction{}
}

func (action *ProcessAction) Execute(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) error {
	executable := getExecutableRegistryInstanceFunc().Get(object.Executable)

	return executable().Execute(chartContext)
}

func (action *ProcessAction) GetNext(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error) {
	connection := chart.GetSingleConnectionBySourceId(object.Id)

	if connection == nil {
		return nil, fmt.Errorf("no source connection for process action %s", object.Label)
	}

	target := chart.GetObjectById(connection.Cell.Target)
	if target == nil {
		return nil, fmt.Errorf("no target for connection action %s", connection.Label)
	}

	return target, nil
}
