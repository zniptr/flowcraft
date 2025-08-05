package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zniptr/flowcraft/internal/file"
)

type ChartXmlParserMock struct {
	mock.Mock
}

func NewChartXmlParserMock() *ChartXmlParserMock {
	return &ChartXmlParserMock{}
}

func (mock *ChartXmlParserMock) ParseDiagrams(data []byte) ([]file.Diagram, error) {
	args := mock.Called(data)

	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]file.Diagram), args.Error(1)
}
