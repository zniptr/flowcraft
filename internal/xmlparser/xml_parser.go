package xmlparser

import (
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/helpers"
)

type XmlParser interface {
	ParseDiagrams(data []byte) ([]file.Diagram, error)
}

type XmlParserImpl struct {
	xmlHelper helpers.XmlHelper
}

func NewXmlParser(xmlHelper helpers.XmlHelper) XmlParser {
	return &XmlParserImpl{
		xmlHelper: xmlHelper,
	}
}

func (parser *XmlParserImpl) ParseDiagrams(data []byte) ([]file.Diagram, error) {
	var chart file.File

	err := parser.xmlHelper.Unmarshal(data, &chart)
	if err != nil {
		return nil, err
	}

	return chart.Diagrams, nil
}
