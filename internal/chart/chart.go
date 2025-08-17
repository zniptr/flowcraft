package chart

import (
	"github.com/zniptr/flowcraft/internal/file"
)

type Chart interface {
	GetStart() *file.Object
	GetObject(id string) *file.Object
	GetOutgoingConnection(id string) *file.Object
	GetOutgoingNonDefaultConnections(id string) []*file.Object
	GetOutgoingDefaultConnection(id string) *file.Object
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

func (chart *ChartImpl) GetObject(id string) *file.Object {
	for _, object := range chart.diagram.Graph.Root.Objects {
		if object.Id == id {
			return &object
		}
	}

	return nil
}

func (chart *ChartImpl) GetOutgoingConnection(id string) *file.Object {
	for _, object := range chart.diagram.Graph.Root.Objects {
		if object.Type == "connection" && object.Cell.Source == id {
			return &object
		}
	}

	return nil
}

func (chart *ChartImpl) GetOutgoingNonDefaultConnections(id string) []*file.Object {
	objects := []*file.Object{}

	for _, object := range chart.diagram.Graph.Root.Objects {
		if object.Type == "connection" && object.Cell.Source == id && object.Default != "1" {
			objects = append(objects, &object)
		}
	}

	return objects
}

func (chart *ChartImpl) GetOutgoingDefaultConnection(id string) *file.Object {
	for _, object := range chart.diagram.Graph.Root.Objects {
		if object.Type == "connection" && object.Cell.Source == id && object.Default == "1" {
			return &object
		}
	}

	return nil
}
