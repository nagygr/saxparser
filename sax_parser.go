package saxparser

import (
	"encoding/xml"
	"io"
)

type SaxParser struct {
	decoder *xml.Decoder
	handler SaxHandler
}

func NewParser(reader io.Reader, handler SaxHandler) *SaxParser {
	return &SaxParser{
		decoder: xml.NewDecoder(reader),
		handler: handler,
	}
}
