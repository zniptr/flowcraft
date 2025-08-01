package chart

import "encoding/xml"

type ChartFile struct {
	XMLName  xml.Name  `xml:"mxfile"`
	Diagrams []Diagram `xml:"diagram"`
}

type Diagram struct {
	XMLName xml.Name `xml:"diagram"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Graph   Graph    `xml:"mxGraphModel"`
}

type Graph struct {
	XMLName xml.Name `xml:"mxGraphModel"`
	Root    Root     `xml:"root"`
}

type Root struct {
	XMLName xml.Name `xml:"root"`
	Objects []Object `xml:"object"`
}

type Object struct {
	XMLName    xml.Name `xml:"object"`
	Id         string   `xml:"id,attr"`
	Label      string   `xml:"label,attr"`
	Type       string   `xml:"type,attr"`
	Condition  string   `xml:"condition,attr"`
	Default    string   `xml:"Default,attr"`
	Executable string   `xml:"executable,attr"`
	Link       string   `xml:"link,attr"`
}
