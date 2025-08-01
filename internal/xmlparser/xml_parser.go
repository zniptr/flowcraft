package xmlparser

import (
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/helpers"
)

type XmlParser interface {
	ParseDiagrams(data []byte) ([]chart.Diagram, error)
}

type XmlParserImpl struct {
	xmlHelper helpers.XmlHelper
}

func NewXmlParser(xmlHelper helpers.XmlHelper) XmlParser {
	return &XmlParserImpl{
		xmlHelper: xmlHelper,
	}
}

func (parser *XmlParserImpl) ParseDiagrams(data []byte) ([]chart.Diagram, error) {
	var chart chart.ChartFile

	err := parser.xmlHelper.Unmarshal(data, &chart)
	if err != nil {
		return nil, err
	}

	return chart.Diagrams, nil
}
