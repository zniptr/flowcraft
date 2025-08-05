package executable

import "github.com/zniptr/flowcraft/pkg/chartcontext"

type Executable interface {
	Execute(chartContext chartcontext.ChartContext) error
}

type ExecutableFactory func() Executable
