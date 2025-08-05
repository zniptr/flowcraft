package file

import "encoding/xml"

type File struct {
	XMLName  xml.Name  `xml:"mxfile"`
	Diagrams []Diagram `xml:"diagram"`
}
