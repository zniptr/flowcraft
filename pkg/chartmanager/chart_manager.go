package chartmanager

import (
	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/filereader"
	"github.com/zniptr/flowcraft/internal/helpers"
	"github.com/zniptr/flowcraft/internal/xmlparser"
)

type chartManager interface {
	LoadCharts(path string) error
}

type chartManagerImpl struct {
	charts     map[string]chart.Diagram
	fileReader filereader.FileReader
	xmlParser  xmlparser.XmlParser
}

var (
	newChartFileReaderFunc = filereader.NewFileReader(helpers.NewOsHelper(), helpers.NewFilepathHelper())
	newChartXmlParserFunc  = xmlparser.NewXmlParser(helpers.NewXmlHelper())
)

func NewChartManager() chartManager {
	return &chartManagerImpl{
		charts:     make(map[string]chart.Diagram),
		fileReader: newChartFileReaderFunc,
		xmlParser:  newChartXmlParserFunc,
	}
}

func (manager *chartManagerImpl) LoadCharts(path string) error {
	files, err := manager.fileReader.ReadDirectory(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := manager.parseChart(file, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (manager *chartManagerImpl) parseChart(file helpers.DirEntryHelper, path string) error {
	if !manager.fileReader.IsValidChartFile(file) {
		return nil
	}

	data, err := manager.fileReader.ReadFile(path, file)
	if err != nil {
		return err
	}

	diagrams, err := manager.xmlParser.ParseDiagrams(data)
	if err != nil {
		return err
	}

	manager.storeDiagrams(diagrams)

	return nil
}

func (manager *chartManagerImpl) storeDiagrams(diagrams []chart.Diagram) {
	for _, diagram := range diagrams {
		manager.charts[diagram.Id] = diagram
	}
}
