package file

import "encoding/xml"

type Object struct {
	XMLName    xml.Name `xml:"object"`
	Id         string   `xml:"id,attr"`
	Label      string   `xml:"label,attr"`
	Type       string   `xml:"type,attr"`
	Condition  string   `xml:"condition,attr"`
	Default    string   `xml:"default,attr"`
	Executable string   `xml:"executable,attr"`
	Link       string   `xml:"link,attr"`
	Cell       Cell     `xml:"mxCell"`
}
