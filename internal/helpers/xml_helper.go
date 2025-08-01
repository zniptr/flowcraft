package helpers

import "encoding/xml"

type XmlHelper interface {
	Unmarshal(data []byte, v any) error
}

type XmlHelperImpl struct{}

func NewXmlHelper() XmlHelper {
	return &XmlHelperImpl{}
}

func (helper *XmlHelperImpl) Unmarshal(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}
