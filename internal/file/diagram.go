package file

import "encoding/xml"

type Diagram struct {
	XMLName xml.Name `xml:"diagram"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Graph   Graph    `xml:"mxGraphModel"`
}
