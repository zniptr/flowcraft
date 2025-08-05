package file

import "encoding/xml"

type Root struct {
	XMLName xml.Name `xml:"root"`
	Objects []Object `xml:"object"`
}
