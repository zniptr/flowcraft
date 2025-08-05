package chart

import "github.com/zniptr/flowcraft/internal/file"

type Chart interface {
	GetStart() *file.Object
	GetObjectById(id string) *file.Object
	GetSingleConnectionBySourceId(id string) *file.Object
}

type ChartImpl struct {
	diagram file.Diagram
}

func NewChart(diagram file.Diagram) Chart {
	return &ChartImpl{
		diagram: diagram,
	}
}

func (chart *ChartImpl) GetStart() *file.Object {
	for _, object := range chart.diagram.Graph.Root.Objects {
		if object.Type == "start" {
			return &object
		}
	}

	return nil
}

func (chart *ChartImpl) GetObjectById(id string) *file.Object {
	for _, object := range chart.diagram.Graph.Root.Objects {
		if object.Id == id {
			return &object
		}
	}

	return nil
}

func (chart *ChartImpl) GetSingleConnectionBySourceId(id string) *file.Object {
	for _, object := range chart.diagram.Graph.Root.Objects {
		if object.Type == "connection" && object.Cell.Source == id {
			return &object
		}
	}

	return nil
}
