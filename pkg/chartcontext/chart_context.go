package chartcontext

type ChartContext interface{}

type chartContextImpl struct {
	context map[string]any
}

func NewChartContext(context map[string]any) ChartContext {
	return &chartContextImpl{
		context: context,
	}
}
