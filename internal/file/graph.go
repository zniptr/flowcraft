package file

import "encoding/xml"

type Graph struct {
	XMLName xml.Name `xml:"mxGraphModel"`
	Root    Root     `xml:"root"`
}
