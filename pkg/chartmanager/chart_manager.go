package chartmanager

import (
	"errors"

	"github.com/zniptr/flowcraft/internal/chart"
	"github.com/zniptr/flowcraft/internal/chartinstance"
	"github.com/zniptr/flowcraft/internal/file"
	"github.com/zniptr/flowcraft/internal/filereader"
	"github.com/zniptr/flowcraft/internal/helpers"
	"github.com/zniptr/flowcraft/internal/xmlparser"
	"github.com/zniptr/flowcraft/pkg/chartcontext"
)

type ChartManager interface {
	LoadCharts(path string) error
	StartChartInstance(name string, context map[string]any) error
	storeCharts(diagrams []file.Diagram)
}

type chartManagerImpl struct {
	charts     map[string]chart.Chart
	fileReader filereader.FileReader
	xmlParser  xmlparser.XmlParser
}

var (
	newChartFileReaderFunc = filereader.NewFileReader
	newChartXmlParserFunc  = xmlparser.NewXmlParser
	newChartContextFunc    = chartcontext.NewChartContext
	newChartInstanceFunc   = chartinstance.NewChartInstance
)

func NewChartManager() ChartManager {
	return &chartManagerImpl{
		charts:     make(map[string]chart.Chart),
		fileReader: newChartFileReaderFunc(helpers.NewOsHelper(), helpers.NewFilepathHelper()),
		xmlParser:  newChartXmlParserFunc(helpers.NewXmlHelper()),
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

	manager.storeCharts(diagrams)

	return nil
}

func (manager *chartManagerImpl) storeCharts(diagrams []file.Diagram) {
	for _, diagram := range diagrams {
		manager.charts[diagram.Name] = chart.NewChart(diagram)
	}
}

func (manager *chartManagerImpl) StartChartInstance(name string, context map[string]any) error {
	chart := manager.charts[name]

	if chart == nil {
		return errors.New("chart not found")
	}

	chartContext := newChartContextFunc(context)
	chartInstance := newChartInstanceFunc(chartContext, chart)

	err := chartInstance.Run()
	if err != nil {
		return err
	}

	return nil
}
