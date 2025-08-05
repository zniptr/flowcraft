package chartcontext

type ChartContext interface {
	GetVariable(name string) any
	SetVariable(name string, value any)
}

type chartContextImpl struct {
	context map[string]any
}

func NewChartContext(context map[string]any) ChartContext {
	return &chartContextImpl{
		context: context,
	}
}

func (context *chartContextImpl) GetVariable(name string) any {
	return context.context[name]
}

func (context *chartContextImpl) SetVariable(name string, value any) {
	context.context[name] = value
}
