package chart

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zniptr/flowcraft/internal/file"
)

type ChartTestSuite struct {
	suite.Suite
	mockId string

	mockDiagram              file.Diagram
	mockObject               file.Object
	mockConnection           file.Object
	mockNonDefaultConnection file.Object
	mockDefaultConnection    file.Object
}

func (suite *ChartTestSuite) SetupTest() {
	suite.mockId = "testId"

	suite.mockDiagram = file.Diagram{}
	suite.mockObject = file.Object{Id: suite.mockId, Type: "start"}
	suite.mockConnection = file.Object{Type: "connection", Cell: file.Cell{Source: suite.mockId}}
	suite.mockNonDefaultConnection = file.Object{Type: "connection", Default: "", Cell: file.Cell{Source: suite.mockId}}
	suite.mockDefaultConnection = file.Object{Type: "connection", Default: "1", Cell: file.Cell{Source: suite.mockId}}
}

func (suite *ChartTestSuite) TestNewChart_whenCreateChart_thenReturnChart() {
	chart := NewChart(suite.mockDiagram)

	suite.NotNil(chart)
	suite.IsType(&ChartImpl{}, chart)
}

func (suite *ChartTestSuite) TestGetStart_whenNotFound_thenReturnNil() {
	start := NewChart(suite.mockDiagram).GetStart()

	suite.Nil(start)
}

func (suite *ChartTestSuite) TestGetStart_whenFound_thenReturnStart() {
	suite.mockDiagram.Graph.Root.Objects = []file.Object{suite.mockObject}

	start := NewChart(suite.mockDiagram).GetStart()

	suite.Equal(&suite.mockObject, start)
}

func (suite *ChartTestSuite) TestGetObject_whenNotFound_thenReturnNil() {
	object := NewChart(suite.mockDiagram).GetObject(suite.mockId)

	suite.Nil(object)
}

func (suite *ChartTestSuite) TestGetObject_whenFound_returnReturnObject() {
	suite.mockDiagram.Graph.Root.Objects = []file.Object{suite.mockObject}

	object := NewChart(suite.mockDiagram).GetObject(suite.mockId)

	suite.Equal(&suite.mockObject, object)
}

func (suite *ChartTestSuite) TestGetOutgoingConnection_whenNotFound_thenReturnNil() {
	connection := NewChart(suite.mockDiagram).GetOutgoingConnection(suite.mockId)

	suite.Nil(connection)
}

func (suite *ChartTestSuite) TestGetOutgoingConnection_whenFound_thenReturnConnection() {
	suite.mockDiagram.Graph.Root.Objects = []file.Object{suite.mockConnection}

	connection := NewChart(suite.mockDiagram).GetOutgoingConnection(suite.mockId)

	suite.Equal(&suite.mockConnection, connection)
}

func (suite *ChartTestSuite) TestGetOutgoingNonDefaultConnections_whenNotFound_thenReturnEmptyList() {
	connections := NewChart(suite.mockDiagram).GetOutgoingNonDefaultConnections(suite.mockId)

	suite.Len(connections, 0)
}

func (suite *ChartTestSuite) TestGetOutgoingNonDefaultConnections_whenFound_thenReturnList() {
	suite.mockDiagram.Graph.Root.Objects = []file.Object{suite.mockNonDefaultConnection}

	connections := NewChart(suite.mockDiagram).GetOutgoingNonDefaultConnections(suite.mockId)

	suite.Len(connections, 1)
	suite.Equal(&suite.mockNonDefaultConnection, connections[0])
}

func (suite *ChartTestSuite) TestGetOutgoingDefaultConnection_whenNotFound_thenReturnNil() {
	connection := NewChart(suite.mockDiagram).GetOutgoingDefaultConnection(suite.mockId)

	suite.Nil(connection)
}

func (suite *ChartTestSuite) TestGetOutgoingDefaultConnection_whenFound_thenReturnConnection() {
	suite.mockDiagram.Graph.Root.Objects = []file.Object{suite.mockDefaultConnection}

	connection := NewChart(suite.mockDiagram).GetOutgoingDefaultConnection(suite.mockId)

	suite.Equal(&suite.mockDefaultConnection, connection)
}

func TestChartTestSuite(t *testing.T) {
	suite.Run(t, new(ChartTestSuite))
}
