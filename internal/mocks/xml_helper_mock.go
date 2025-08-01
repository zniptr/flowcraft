package mocks

import "github.com/stretchr/testify/mock"

type XmlHelperMock struct {
	mock.Mock
}

func NewXmlHelperMock() *XmlHelperMock {
	return &XmlHelperMock{}
}

func (mock *XmlHelperMock) Unmarshal(data []byte, v any) error {
	args := mock.Called(data, v)
	return args.Error(0)
}
