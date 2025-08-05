package file

import "encoding/xml"

type Cell struct {
	XMLName xml.Name `xml:"mxCell"`
	Source  string   `xml:"source,attr"`
	Target  string   `xml:"target,attr"`
}
