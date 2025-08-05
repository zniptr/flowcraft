package actions

import (
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type Action interface {
	Execute(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) error
	GetNext(chartContext chartcontext.ChartContext, chart chart.Chart, object *file.Object) (*file.Object, error)
}
